package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Clients struct {

	/*CcServiceDashboards - Descr: Used to grant scope for SSO clients for service brokers Default: openid,cloud_controller_service_permissions.read
*/
	CcServiceDashboards *CcServiceDashboards `yaml:"cc-service-dashboards,omitempty"`

	/*CcServiceBrokerClient - Descr: (DEPRECATED) - Used to grant scope for SSO clients for service brokers Default: openid,cloud_controller_service_permissions.read
*/
	CcServiceBrokerClient *CcServiceBrokerClient `yaml:"cc_service_broker_client,omitempty"`

}
