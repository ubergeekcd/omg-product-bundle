package flannel 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type FlannelJob struct {

	/*Apiserver - Descr: ip of api server Default: <nil>
*/
	Apiserver *Apiserver `yaml:"apiserver,omitempty"`

	/*Kubernetes - Descr: master Default: false
*/
	Kubernetes *Kubernetes `yaml:"kubernetes,omitempty"`

}