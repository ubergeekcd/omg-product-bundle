package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CloudControllerWorkerJob struct {

	/*Cc - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Version - Descr: 'version' attribute in the /v2/info endpoint Default: 0
*/
	Version interface{} `yaml:"version,omitempty"`

	/*RequestTimeoutInSeconds - Descr: Timeout for requests in seconds. Default: 900
*/
	RequestTimeoutInSeconds interface{} `yaml:"request_timeout_in_seconds,omitempty"`

	/*AppDomains - Descr: Array of domains for user apps (example: 'user.app.space.foo', a user app called 'neat' will listen at 'http://neat.user.app.space.foo') Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*SystemDomainOrganization - Descr: An organization that will be created as part of the seeding process. When the system_domain is not shared with (in the list of) app_domains, this is required as the system_domain will be created as a PrivateDomain in this organization. Default: 
*/
	SystemDomainOrganization interface{} `yaml:"system_domain_organization,omitempty"`

	/*NfsServer - Descr: NFS server for droplets and apps (not used in an AWS deploy, use s3 instead) Default: <nil>
*/
	NfsServer *NfsServer `yaml:"nfs_server,omitempty"`

	/*Uaa - Descr: (DEPRECATED) - Used for generating SSO clients for service brokers. Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Domain - Descr: Deprecated in favor of system_domain. Domain where cloud_controller will listen (api.domain) Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Nats - Descr: Password for cc client to connect to NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*DeaNext - Descr: Memory limit in mb for staging tasks Default: 1024
*/
	DeaNext *DeaNext `yaml:"dea_next,omitempty"`

	/*MetronEndpoint - Descr: The port used to emit messages to the Metron agent Default: 3457
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*Build - Descr: 'build' attribute in the /v2/info endpoint Default: 
*/
	Build interface{} `yaml:"build,omitempty"`

	/*Login - Descr: http or https Default: https
*/
	Login *Login `yaml:"login,omitempty"`

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*Ccdb - Descr: The address of the database server Default: <nil>
*/
	Ccdb *Ccdb `yaml:"ccdb,omitempty"`

	/*Hm9000 - Descr: Port of the hm9000 Api Server Default: 5155
*/
	Hm9000 *Hm9000 `yaml:"hm9000,omitempty"`

	/*LoggerEndpoint - Descr: Port for logger endpoint listed at /v2/info Default: 443
*/
	LoggerEndpoint *LoggerEndpoint `yaml:"logger_endpoint,omitempty"`

	/*Ssl - Descr: specifies that the job is allowed to skip ssl cert verification Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Description - Descr: 'description' attribute in the /v2/info endpoint Default: 
*/
	Description interface{} `yaml:"description,omitempty"`

	/*Name - Descr: 'name' attribute in the /v2/info endpoint Default: 
*/
	Name interface{} `yaml:"name,omitempty"`

	/*SupportAddress - Descr: 'support' attribute in the /v2/info endpoint Default: 
*/
	SupportAddress interface{} `yaml:"support_address,omitempty"`

}