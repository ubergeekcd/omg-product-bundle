package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MetronEndpoint struct {

	/*Port - Descr: The port used to emit messages to the Metron agent Default: 3457
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Host - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	Host interface{} `yaml:"host,omitempty"`

}
