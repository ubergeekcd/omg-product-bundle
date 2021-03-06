package proxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ProxyJob struct {

	/*NetworkName - Descr: The name of the network (needed for the syslog aggregator) Default: <nil>
*/
	NetworkName interface{} `yaml:"network_name,omitempty"`

	/*Standalone - Descr: Standalone Mode: Are you deploying MySQL without a CloudFoundry deployment? Default: false
*/
	Standalone interface{} `yaml:"standalone,omitempty"`

	/*SyslogAggregator - Descr: Transport to be used when forwarding logs (tcp|udp|relp). Default: tcp
*/
	SyslogAggregator *SyslogAggregator `yaml:"syslog_aggregator,omitempty"`

	/*Nats - Descr: Password to register a route via NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Proxy - Descr: Timeout (milliseconds) before assuming a backend is unhealthy Default: 5000
*/
	Proxy *Proxy `yaml:"proxy,omitempty"`

	/*ExternalHost - Descr: Domain of the route registered for the UI via NATS (with the router in cf-release) Default: <nil>
*/
	ExternalHost interface{} `yaml:"external_host,omitempty"`

	/*ClusterIps - Descr: List of nodes.  Must have the same number of ips as there are nodes in the cluster Default: <nil>
*/
	ClusterIps interface{} `yaml:"cluster_ips,omitempty"`

}