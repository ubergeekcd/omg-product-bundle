package broker_registrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BrokerRegistrarJob struct {

	/*Cf - Descr: Username of the admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*Broker - Descr: Port for the service broker Default: 80
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*Redis - Descr: Service name Default: p-redis
*/
	Redis *Redis `yaml:"redis,omitempty"`

}
