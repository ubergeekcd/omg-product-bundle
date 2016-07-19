package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CcServiceBrokerClient struct {

	/*Secret - Descr: (DEPRECATED) - Used for generating SSO clients for service brokers. Default: <nil>
*/
	Secret interface{} `yaml:"secret,omitempty"`

	/*Scope - Descr: (DEPRECATED) - Used to grant scope for SSO clients for service brokers Default: openid,cloud_controller_service_permissions.read
*/
	Scope interface{} `yaml:"scope,omitempty"`

}
