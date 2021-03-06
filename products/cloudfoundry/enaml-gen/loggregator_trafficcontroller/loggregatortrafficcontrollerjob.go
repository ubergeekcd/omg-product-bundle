package loggregator_trafficcontroller 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type LoggregatorTrafficcontrollerJob struct {

	/*Uaa - Descr: Doppler's client secret to connect to UAA Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*MetronEndpoint - Descr: The port used to emit dropsonde messages to the Metron agent Default: 3457
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*TrafficController - Descr: Enable logging of all requests made to the Traffic Controller in CEF format Default: false
*/
	TrafficController *TrafficController `yaml:"traffic_controller,omitempty"`

	/*Doppler - Descr: Doppler's client id to connect to UAA Default: doppler
*/
	Doppler *Doppler `yaml:"doppler,omitempty"`

	/*Login - Descr: Protocol to use to connect to UAA (used in case uaa.url is not set) Default: https
*/
	Login *Login `yaml:"login,omitempty"`

	/*Loggregator - Descr: Number of concurrent requests to ETCD Default: 10
*/
	Loggregator *Loggregator `yaml:"loggregator,omitempty"`

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*Cc - Descr: API URI of cloud controller Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Ssl - Descr: when connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}