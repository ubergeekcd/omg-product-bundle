package rep 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Garden struct {

	/*Address - Descr: Garden server listening address. Default: /var/vcap/data/garden/garden.sock
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Network - Descr: Network type for the garden server connection (tcp or unix). Default: unix
*/
	Network interface{} `yaml:"network,omitempty"`

}
