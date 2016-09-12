package rabbitmq_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Rabbitmq struct {

	/*Administrator - Descr: The username used to access an admin account of the Rabbit servers Default: <nil>
*/
	Administrator *Administrator `yaml:"administrator,omitempty"`

	/*OperatorSetPolicy - Descr: Operator policy definition Default: <nil>
*/
	OperatorSetPolicy *OperatorSetPolicy `yaml:"operator_set_policy,omitempty"`

	/*DnsHost - Descr: External load balancer host Default: <nil>
*/
	DnsHost interface{} `yaml:"dns_host,omitempty"`

	/*Hosts - Descr: List of IPs of the Rabbit servers (maybe proxied) upon which to allocate resources Default: <nil>
*/
	Hosts interface{} `yaml:"hosts,omitempty"`

	/*Ssl - Descr: Set to true if RabbitMQ cluster is configured to use TLS/SSL Default: false
*/
	Ssl interface{} `yaml:"ssl,omitempty"`

	/*ManagementDomain - Descr: Domain that will be used to access RabbitMQ management UI, typically the system CF/CC domain. Default: <nil>
*/
	ManagementDomain interface{} `yaml:"management_domain,omitempty"`

}