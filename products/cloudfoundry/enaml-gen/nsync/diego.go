package nsync 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Diego struct {

	/*Nsync - Descr: enable ssl for all communication with the bbs Default: true
*/
	Nsync *Nsync `yaml:"nsync,omitempty"`

	/*Ssl - Descr: when connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}