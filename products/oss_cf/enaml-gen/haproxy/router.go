package haproxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Router struct {

	/*Servers - Descr: Array of router IPs Default: []
*/
	Servers interface{} `yaml:"servers,omitempty"`

	/*Port - Descr: Listening port for Router Default: 80
*/
	Port interface{} `yaml:"port,omitempty"`

}