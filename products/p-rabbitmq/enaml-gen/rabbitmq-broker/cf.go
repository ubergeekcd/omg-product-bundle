package rabbitmq_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cf struct {

	/*Domain - Descr: Domain shared by the UAA and CF API, e.g. 'bosh-lite.com' Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Nats - Descr: Port that NATS listens on Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

}