package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Droplets struct {

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

	/*WebdavConfig - Descr: The basic auth password that CC uses to connect to the admin endpoint on webdav Default: 
*/
	WebdavConfig *DropletsWebdavConfig `yaml:"webdav_config,omitempty"`

	/*Cdn - Descr: URI for a CDN to used for droplet downloads Default: 
*/
	Cdn *DropletsCdn `yaml:"cdn,omitempty"`

	/*DropletDirectoryKey - Descr: Directory (bucket) used store droplets.  It does not have be pre-created. Default: cc-droplets
*/
	DropletDirectoryKey interface{} `yaml:"droplet_directory_key,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

}