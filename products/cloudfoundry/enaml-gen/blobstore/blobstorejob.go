package blobstore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BlobstoreJob struct {

	/*Blobstore - Descr: TCP port on which the blobstore server (nginx) listens Default: 80
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Domain - Descr: The system domain.  The public server will listen on host 'blobstore.system-domain.tld' Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

}
