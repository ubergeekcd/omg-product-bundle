package destroy_service_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*OrgName - Descr: Org that hosts the Service Broker Default: system
*/
	OrgName interface{} `yaml:"org_name,omitempty"`

	/*Password - Descr: Broker basic auth password Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*SpaceName - Descr: Space that hosts the Service Broker Default: p-spring-cloud-services
*/
	SpaceName interface{} `yaml:"space_name,omitempty"`

	/*User - Descr: Broker basic auth user Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*PlanNames - Descr: Broker plan names  Default: p-service-registry,p-config-server,p-circuit-breaker-dashboard
*/
	PlanNames interface{} `yaml:"plan_names,omitempty"`

}