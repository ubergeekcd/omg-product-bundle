package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MetronAgent struct {

	/*ListeningPort - Descr: Port the metron agent is listening on to receive dropsonde log messages Default: 3457
*/
	ListeningPort interface{} `yaml:"listening_port,omitempty"`

	/*Tcp - Descr: The maximum time that a message can stay in the batching buffer before being flushed Default: 100
*/
	Tcp *Tcp `yaml:"tcp,omitempty"`

	/*Protocols - Descr: A priority list of protocols for metron to connect to doppler over.  Metron will refuse to connect to doppler over any protocol not in this list. Default: [udp]
*/
	Protocols interface{} `yaml:"protocols,omitempty"`

	/*BufferSize - Descr: DEPRECATED Default: 10000
*/
	BufferSize interface{} `yaml:"buffer_size,omitempty"`

	/*PreferredProtocol - Descr: DEPRECATED - replaced with metron_agent.protocols Default: udp
*/
	PreferredProtocol interface{} `yaml:"preferred_protocol,omitempty"`

	/*ListeningAddress - Descr: Address the metron agent is listening on to receive dropsonde log messages provided for BOSH links and should not be overwritten Default: 127.0.0.1
*/
	ListeningAddress interface{} `yaml:"listening_address,omitempty"`

	/*EnableBuffer - Descr: DEPRECATED Default: false
*/
	EnableBuffer interface{} `yaml:"enable_buffer,omitempty"`

	/*Deployment - Descr: Name of deployment (added as tag on all outgoing metrics) Default: <nil>
*/
	Deployment interface{} `yaml:"deployment,omitempty"`

	/*Logrotate - Descr: The number of files that logrotate will keep around on the VM Default: 7
*/
	Logrotate *Logrotate `yaml:"logrotate,omitempty"`

	/*Etcd - Descr: PEM-encoded client key Default: 
*/
	Etcd *MetronAgentEtcd `yaml:"etcd,omitempty"`

	/*Tls - Descr: TLS client certificate Default: 
*/
	Tls *MetronAgentTls `yaml:"tls,omitempty"`

	/*Zone - Descr: Availability zone where this agent is running Default: <nil>
*/
	Zone interface{} `yaml:"zone,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose mode Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

}