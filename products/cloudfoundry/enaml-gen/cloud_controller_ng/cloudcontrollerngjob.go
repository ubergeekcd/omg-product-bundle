package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CloudControllerNgJob struct {

	/*SupportAddress - Descr: 'support' attribute in the /v2/info endpoint Default: 
*/
	SupportAddress interface{} `yaml:"support_address,omitempty"`

	/*NfsServer - Descr: NFS server for droplets and apps (not used in an AWS deploy, use s3 instead) Default: <nil>
*/
	NfsServer *NfsServer `yaml:"nfs_server,omitempty"`

	/*Router - Descr: Support for route services is disabled when no value is configured. Default: 
*/
	Router *Router `yaml:"router,omitempty"`

	/*Name - Descr: 'name' attribute in the /v2/info endpoint Default: 
*/
	Name interface{} `yaml:"name,omitempty"`

	/*Version - Descr: 'version' attribute in the /v2/info endpoint Default: 0
*/
	Version interface{} `yaml:"version,omitempty"`

	/*Ccdb - Descr: The type of database being used. mysql or postgres Default: postgres
*/
	Ccdb *Ccdb `yaml:"ccdb,omitempty"`

	/*Ssl - Descr: specifies that the job is allowed to skip ssl cert verification Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*SystemDomainOrganization - Descr: The User Org that owns the system_domain, required if system_domain is defined Default: 
*/
	SystemDomainOrganization interface{} `yaml:"system_domain_organization,omitempty"`

	/*Domain - Descr: domain where cloud_controller will listen (api.domain) often the same as the system domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*AppSsh - Descr: External port for SSH access to application instances Default: 2222
*/
	AppSsh *AppSsh `yaml:"app_ssh,omitempty"`

	/*Doppler - Descr: Port for doppler_logging_endpoint listed at /v2/info Default: 443
*/
	Doppler *Doppler `yaml:"doppler,omitempty"`

	/*Build - Descr: 'build' attribute in the /v2/info endpoint Default: 
*/
	Build interface{} `yaml:"build,omitempty"`

	/*Login - Descr: whether use login as the authorization endpoint or not Default: true
*/
	Login *Login `yaml:"login,omitempty"`

	/*RoutingApi - Descr: Whether to expose the routing_endpoint listed at /v2/info. Enable this after deploying the Routing API Default: false
*/
	RoutingApi *RoutingApi `yaml:"routing_api,omitempty"`

	/*LoggerEndpoint - Descr: Port for logger endpoint listed at /v2/info Default: 443
*/
	LoggerEndpoint *LoggerEndpoint `yaml:"logger_endpoint,omitempty"`

	/*AppDomains - Descr: Array of domains for user apps (example: 'user.app.space.foo', a user app called 'neat' will listen at 'http://neat.user.app.space.foo') Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*Description - Descr: 'description' attribute in the /v2/info endpoint Default: 
*/
	Description interface{} `yaml:"description,omitempty"`

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*DeaNext - Descr: Advertise interval for DEAs Default: 5
*/
	DeaNext *DeaNext `yaml:"dea_next,omitempty"`

	/*Hm9000 - Descr: URL of the hm9000 server Default: <nil>
*/
	Hm9000 *Hm9000 `yaml:"hm9000,omitempty"`

	/*MetronEndpoint - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*Cc - Descr: URL of the Diego stager service Default: http://stager.service.cf.internal:8888
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Uaa - Descr: Used for fetching routing information from the Routing API Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Nats - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*RequestTimeoutInSeconds - Descr: Timeout for requests in seconds. Default: 900
*/
	RequestTimeoutInSeconds interface{} `yaml:"request_timeout_in_seconds,omitempty"`

}
