package swarm_manager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Docker struct {

	/*Tls - Descr: Use TLS Default: false
*/
	Tls interface{} `yaml:"tls,omitempty"`

	/*TlsVerify - Descr: Use TLS and verify the remote Default: false
*/
	TlsVerify interface{} `yaml:"tls_verify,omitempty"`

	/*TlsKey - Descr: TLS key file Default: <nil>
*/
	TlsKey interface{} `yaml:"tls_key,omitempty"`

	/*TlsCacert - Descr: Trust only remotes providing a certificate signed by the CA given here Default: <nil>
*/
	TlsCacert interface{} `yaml:"tls_cacert,omitempty"`

	/*TlsCert - Descr: TLS certificate file Default: <nil>
*/
	TlsCert interface{} `yaml:"tls_cert,omitempty"`

}