package destroy_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DestroyBrokerJob struct {

	/*Domain - Descr: The CF top-level domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Autoscale - Descr: Username of the CF admin user Default: <nil>
*/
	Autoscale *Autoscale `yaml:"autoscale,omitempty"`

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}