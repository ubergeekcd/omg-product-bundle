package broker_registrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*Name - Descr: Name of the service broker Default: <nil>
*/
	Name interface{} `yaml:"name,omitempty"`

	/*Port - Descr: Port for the service broker Default: 443
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Username - Descr: Basic Auth username for the service broker Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Protocol - Descr: Protocol (http/https) provided to the Cloud Controller when registering/deregistering the broker Default: https
*/
	Protocol interface{} `yaml:"protocol,omitempty"`

	/*Host - Descr: Host address of the service broker Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*Password - Descr: Basic Auth password for the service broker Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

}