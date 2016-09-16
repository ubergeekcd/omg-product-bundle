package register_service_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RegisterServiceBrokerJob struct {

	/*Domain - Descr: The CF top-level domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*AppDomains - Descr: CloudFoundry application domains Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*SpringCloudBroker - Descr: Product Name Default: p-spring-cloud-services
*/
	SpringCloudBroker *SpringCloudBroker `yaml:"spring_cloud_broker,omitempty"`

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}