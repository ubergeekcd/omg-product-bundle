package collector 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CollectorJob struct {

	/*Collector - Descr: enable Datadog plugin Default: false
*/
	Collector *Collector `yaml:"collector,omitempty"`

	/*Nats - Descr: NATS TCP port Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

}
