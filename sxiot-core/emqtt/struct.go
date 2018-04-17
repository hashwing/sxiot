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

type ClusterNodes struct{
	Code int 			`json:"code"`
	Result []NodeStatus 	`json:"result"`
}

type NodeStatus struct{
	Name string `json:"name"`
	Memory string `json:"memory_used"`
	Clients int `json:"clients"`
	Status string `json:"node_status"`
}


type Clients struct{
	Code int 			`json:"code"`
	Result ClientsResult 		`json:"result"`
}

type ClientsResult struct{
	Objects []ClientDesc	`json:"objects"`
}

type ClientDesc struct{
	ID	string	`json:"client_id"`
	IP	string	`json:"ipaddress"`
}