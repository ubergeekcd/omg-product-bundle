package nats 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*MonitorPort - Descr: Port for varz and connz monitoring. 0 means disabled. Default: 0
*/
	MonitorPort interface{} `yaml:"monitor_port,omitempty"`

	/*AuthorizationTimeout - Descr: After accepting a connection, wait up to this many seconds for credentials. Default: 15
*/
	AuthorizationTimeout interface{} `yaml:"authorization_timeout,omitempty"`

	/*Trace - Descr: Enable trace logging output. Default: false
*/
	Trace interface{} `yaml:"trace,omitempty"`

	/*Port - Descr: The port for the NATS server to listen on. Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Machines - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

	/*Debug - Descr: Enable debug logging output. Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

	/*User - Descr: Username for server authentication. Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Password - Descr: Password for server authentication. Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*ProfPort - Descr: Port for pprof. 0 means disabled. Default: 0
*/
	ProfPort interface{} `yaml:"prof_port,omitempty"`

}
