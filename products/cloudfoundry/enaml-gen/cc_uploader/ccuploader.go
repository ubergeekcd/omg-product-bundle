package cc_uploader 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CcUploader struct {

	/*Cc - Descr: External Cloud Controller port Default: 9022
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*ListenAddr - Descr: Address of interface on which to serve files Default: 0.0.0.0:9090
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

	/*DebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17018
*/
	DebugAddr interface{} `yaml:"debug_addr,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*ConsulAgentPort - Descr: local consul agent's port Default: 8500
*/
	ConsulAgentPort interface{} `yaml:"consul_agent_port,omitempty"`

}