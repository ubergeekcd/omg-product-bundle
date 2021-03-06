package rabbitmq_server 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ssl struct {

	/*SecurityOptions - Descr: SSL security options (currently only 'enable_tls1_0') Default: []
*/
	SecurityOptions interface{} `yaml:"security_options,omitempty"`

	/*Verify - Descr: Peer verification method used by RabbitMQ server Default: false
*/
	Verify interface{} `yaml:"verify,omitempty"`

	/*VerificationDepth - Descr: Peer verification depth used by RabbitMQ server Default: 5
*/
	VerificationDepth interface{} `yaml:"verification_depth,omitempty"`

	/*Cacert - Descr: RabbitMQ server CA certificate Default: <nil>
*/
	Cacert interface{} `yaml:"cacert,omitempty"`

	/*FailIfNoPeerCert - Descr: Should RabbitMQ server reject connection if there is no peer cert Default: false
*/
	FailIfNoPeerCert interface{} `yaml:"fail_if_no_peer_cert,omitempty"`

	/*Key - Descr: RabbitMQ server private key Default: <nil>
*/
	Key interface{} `yaml:"key,omitempty"`

	/*Cert - Descr: RabbitMQ server certificate Default: <nil>
*/
	Cert interface{} `yaml:"cert,omitempty"`

}