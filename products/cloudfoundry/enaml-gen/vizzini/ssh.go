package vizzini 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ssh struct {

	/*ProxyAddress - Descr: Host and port for the SSH proxy Default: ssh-proxy.service.cf.internal:2222
*/
	ProxyAddress interface{} `yaml:"proxy_address,omitempty"`

	/*ProxySecret - Descr: Shared secret for the SSH proxy's Diego authenticator Default: <nil>
*/
	ProxySecret interface{} `yaml:"proxy_secret,omitempty"`

}
