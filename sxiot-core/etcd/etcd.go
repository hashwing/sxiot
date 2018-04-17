package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/hashwing/sxiot/sxiot-core/config"
)
// Client etcd client
type Client struct {
	client     *clientv3.Client
	reqTimeout time.Duration
	lID        clientv3.LeaseID
}

// NewEtcd new etcd client
func NewEtcd() (*Client, error) {
	cfg := clientv3.Config{
		Endpoints:   config.CommonConfig.Etcd.URL,
		DialTimeout: time.Duration(10) * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{
		client:     client,
		reqTimeout: time.Duration(10) * time.Second,
	}, nil
}

// Put etcd put
func (ec *Client) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	defer cancel()
	return ec.client.Put(ctx, key, val, opts...)
}

// Get etcd get
func (ec *Client) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	defer cancel()
	return ec.client.Get(ctx, key, opts...)
}

// Delete etcd delete
func (ec *Client) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	defer cancel()
	return ec.client.Delete(ctx, key, opts...)
}

// Grant etcd grant
func (ec *Client) Grant(ttl int64) (*clientv3.LeaseGrantResponse, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	defer cancel()
	return ec.client.Grant(ctx, ttl)
}

// Watch etcd watch
func (ec *Client) Watch(key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	return ec.client.Watch(context.Background(), key, opts...)
}

// KeepAliveOnce keepalive one time
func (ec *Client) KeepAliveOnce(id clientv3.LeaseID) (*clientv3.LeaseKeepAliveResponse, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	defer cancel()
	return ec.client.KeepAliveOnce(ctx, id)
}

// GetLock get lock
func (ec *Client) GetLock(key string, id clientv3.LeaseID) (bool, error) {
	ctx, cancel := NewEtcdTimeoutContext(ec)
	resp, err := ec.client.Txn(ctx).
		If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, "", clientv3.WithLease(id))).
		Commit()
	cancel()

	if err != nil {
		return false, err
	}

	return resp.Succeeded, nil
}

// DelLock delete lock
func (ec *Client) DelLock(key string) error {
	_, err := ec.Delete(key)
	return err
}

// etcdTimeoutContext etcd timeout context
type etcdTimeoutContext struct {
	context.Context

	etcdEndpoints []string
}

// Err err
func (c *etcdTimeoutContext) Err() error {
	err := c.Context.Err()
	if err == context.DeadlineExceeded {
		err = fmt.Errorf("%s: etcd(%v) lost",
			err, c.etcdEndpoints)
	}
	return err
}
// NewEtcdTimeoutContext new etcd timeout context
func NewEtcdTimeoutContext(c *Client) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), c.reqTimeout)
	etcdCtx := &etcdTimeoutContext{}
	etcdCtx.Context = ctx
	etcdCtx.etcdEndpoints = c.client.Endpoints()
	return etcdCtx, cancel
}
