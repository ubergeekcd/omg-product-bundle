package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type PackagesWebdavConfig struct {

	/*Username - Descr: The basic auth user that CC uses to connect to the admin endpoint on webdav Default: 
*/
	Username interface{} `yaml:"username,omitempty"`

	/*BlobstoreTimeout - Descr: The timeout in seconds for requests to the blobstore Default: 5
*/
	BlobstoreTimeout interface{} `yaml:"blobstore_timeout,omitempty"`

	/*Password - Descr: The basic auth password that CC uses to connect to the admin endpoint on webdav Default: 
*/
	Password interface{} `yaml:"password,omitempty"`

	/*PrivateEndpoint - Descr: The location of the webdav server eg: https://blobstore.internal Default: https://blobstore.service.cf.internal:4443
*/
	PrivateEndpoint interface{} `yaml:"private_endpoint,omitempty"`

	/*PublicEndpoint - Descr: The location of the webdav server eg: https://blobstore.com Default: 
*/
	PublicEndpoint interface{} `yaml:"public_endpoint,omitempty"`

	/*CaCert - Descr: The ca cert to use when communicating with webdav Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

}