package blackbox 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Destination struct {

	/*Address - Descr: address for syslog drain (including port) Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Transport - Descr: transport protocol for syslog drain (udp/tcp/tls) Default: tls
*/
	Transport interface{} `yaml:"transport,omitempty"`

}
