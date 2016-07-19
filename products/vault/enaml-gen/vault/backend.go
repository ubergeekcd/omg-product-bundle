package vault 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Backend struct {

	/*UseFile - Descr: Use File backend Default: false
*/
	UseFile interface{} `yaml:"use_file,omitempty"`

	/*UseConsul - Descr: Use Cosul for data store Default: false
*/
	UseConsul interface{} `yaml:"use_consul,omitempty"`

	/*UseZookeeper - Descr: Use Zookeeper as data store Default: false
*/
	UseZookeeper interface{} `yaml:"use_zookeeper,omitempty"`

	/*File - Descr: Path for File backend Default: /var/vcap/store/
*/
	File *File `yaml:"file,omitempty"`

	/*Consul - Descr: Path for Consul Default: vault/
*/
	Consul *Consul `yaml:"consul,omitempty"`

	/*S3 - Descr: AWS access key Default: <nil>
*/
	S3 *S3 `yaml:"s3,omitempty"`

	/*Zookeeper - Descr: Address for Zookeeper Default: localhost:2181
*/
	Zookeeper *Zookeeper `yaml:"zookeeper,omitempty"`

	/*UseInmem - Descr: Use In Memory backend Default: false
*/
	UseInmem interface{} `yaml:"use_inmem,omitempty"`

	/*UseS3 - Descr: Use S3 backend Default: false
*/
	UseS3 interface{} `yaml:"use_s3,omitempty"`

}
