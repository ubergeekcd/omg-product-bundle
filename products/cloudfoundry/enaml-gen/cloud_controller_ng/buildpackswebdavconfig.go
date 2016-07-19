package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BuildpacksWebdavConfig struct {

	/*PrivateEndpoint - Descr: The location of the webdav server eg: https://blobstore.internal Default: https://blobstore.service.cf.internal
*/
	PrivateEndpoint interface{} `yaml:"private_endpoint,omitempty"`

	/*PublicEndpoint - Descr: The location of the webdav server eg: https://blobstore.com Default: 
*/
	PublicEndpoint interface{} `yaml:"public_endpoint,omitempty"`

	/*Password - Descr: The basic auth password that CC uses to connect to the admin endpoint on webdav Default: 
*/
	Password interface{} `yaml:"password,omitempty"`

	/*CaCert - Descr: The ca cert to use when communicating with webdav Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*Username - Descr: The basic auth user that CC uses to connect to the admin endpoint on webdav Default: 
*/
	Username interface{} `yaml:"username,omitempty"`

}
