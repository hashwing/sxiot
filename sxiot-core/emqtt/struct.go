package emqtt

type ClusterSubLists struct{
	Code int 			`json:"code"`
	Result ClusterSubs 	`json:"result"`
}

type ClusterSubs struct{
	Objectds []ClusterSub `json:"objects"`
}

type ClusterSub struct{
	ID string 		`json:"client_id"`
	Topic string 	`json:"topic"`
	QOS int 	 	`json:"qos"`
}