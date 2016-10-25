package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BuildpacksCdn struct {

	/*PrivateKey - Descr: Private key for signing download URIs Default: 
*/
	PrivateKey interface{} `yaml:"private_key,omitempty"`

	/*KeyPairId - Descr: Key pair name for signed download URIs Default: 
*/
	KeyPairId interface{} `yaml:"key_pair_id,omitempty"`

	/*Uri - Descr: URI for a CDN to used for buildpack downloads Default: 
*/
	Uri interface{} `yaml:"uri,omitempty"`

}