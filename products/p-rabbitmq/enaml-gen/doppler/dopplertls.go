package doppler 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DopplerTls struct {

	/*ServerKey - Descr: TLS server key Default: 
*/
	ServerKey interface{} `yaml:"server_key,omitempty"`

	/*Enable - Descr: Enable TLS listener on doppler so that it can receive dropsonde envelopes over TLS transport. If enabled, Cert and Key files must be specified. Default: false
*/
	Enable interface{} `yaml:"enable,omitempty"`

	/*Port - Descr: Port for incoming messages in the dropsonde format over tls listener Default: 3458
*/
	Port interface{} `yaml:"port,omitempty"`

	/*ServerCert - Descr: TLS server certificate Default: 
*/
	ServerCert interface{} `yaml:"server_cert,omitempty"`

}