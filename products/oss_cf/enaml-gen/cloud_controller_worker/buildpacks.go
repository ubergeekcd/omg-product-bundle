package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Buildpacks struct {

	/*WebdavConfig - Descr: The basic auth password that CC uses to connect to the admin endpoint on webdav Default: 
*/
	WebdavConfig *BuildpacksWebdavConfig `yaml:"webdav_config,omitempty"`

	/*Cdn - Descr: URI for a CDN to used for buildpack downloads Default: 
*/
	Cdn *BuildpacksCdn `yaml:"cdn,omitempty"`

	/*BuildpackDirectoryKey - Descr: Directory (bucket) used store buildpacks.  It does not have be pre-created. Default: cc-buildpacks
*/
	BuildpackDirectoryKey interface{} `yaml:"buildpack_directory_key,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

}