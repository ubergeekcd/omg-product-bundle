package broker_deregistrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BrokerDeregistrarJob struct {

	/*Cf - Descr: Password of the admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*Broker - Descr: Basic Auth password for the service broker Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

}