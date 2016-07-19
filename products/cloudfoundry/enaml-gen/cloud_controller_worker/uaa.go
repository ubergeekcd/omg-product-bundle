package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaa struct {

	/*Jwt - Descr: ssl cert defined in the manifest by the UAA, required by the cc to communicate with UAA Default: 
*/
	Jwt *Jwt `yaml:"jwt,omitempty"`

	/*Clients - Descr: Used for generating SSO clients for service brokers. Default: <nil>
*/
	Clients *Clients `yaml:"clients,omitempty"`

	/*Cc - Descr: Symmetric secret used to decode uaa tokens. Used for testing. Default: <nil>
*/
	Cc *UaaCc `yaml:"cc,omitempty"`

	/*Url - Descr: URL of the UAA server Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

}
