package cf_mysql_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cf struct {

	/*ApiUrl - Descr: URL of the CloudFoundry Cloud Controller Default: <nil>
*/
	ApiUrl interface{} `yaml:"api_url,omitempty"`

	/*SkipSslValidation - Descr: Determines whether dashboard verifies SSL certificates when communicating with Cloud Controller and UAA Default: false
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

}