package mysql 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SyslogAggregator struct {

	/*Transport - Descr: Transport to be used when forwarding logs (tcp|udp|relp). Default: tcp
*/
	Transport interface{} `yaml:"transport,omitempty"`

	/*All - Descr: Define whether forwarders should also send non-mysql syslog activity to the aggregator. Default: false
*/
	All interface{} `yaml:"all,omitempty"`

	/*Address - Descr: IP address for syslog aggregator Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Port - Descr: TCP port of syslog aggregator Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

}