package fetch_containers_images 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type FetchContainersImagesJob struct {

	/*Broker - Descr: Services and plans offered by the broker Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*FetchContainersImages - Descr: Docker URL (IP/Socket) Default: <nil>
*/
	FetchContainersImages *FetchContainersImages `yaml:"fetch_containers_images,omitempty"`

}