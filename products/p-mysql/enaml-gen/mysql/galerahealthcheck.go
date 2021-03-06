package mysql 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type GaleraHealthcheck struct {

	/*EndpointUsername - Descr: Username used by the sidecar endpoints for Basic Auth Default: <nil>
*/
	EndpointUsername interface{} `yaml:"endpoint_username,omitempty"`

	/*DbPassword - Descr: Password used by the sidecar to connect to the database Default: <nil>
*/
	DbPassword interface{} `yaml:"db_password,omitempty"`

	/*EndpointPassword - Descr: Password used by the sidecar endpoints for Basic Auth Default: <nil>
*/
	EndpointPassword interface{} `yaml:"endpoint_password,omitempty"`

}