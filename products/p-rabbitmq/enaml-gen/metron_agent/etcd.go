package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Etcd struct {

	/*Maxconcurrentrequests - Descr: Number of concurrent requests to ETCD Default: 10
*/
	Maxconcurrentrequests interface{} `yaml:"maxconcurrentrequests,omitempty"`

	/*Machines - Descr: IPs pointing to the ETCD cluster Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

}