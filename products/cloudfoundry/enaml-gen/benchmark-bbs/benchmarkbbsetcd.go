package benchmark_bbs 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BenchmarkBbsEtcd struct {

	/*Machines - Descr: Addresses pointing to the ETCD cluster Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

	/*ClientCert - Descr: PEM-encoded client certificate Default: <nil>
*/
	ClientCert interface{} `yaml:"client_cert,omitempty"`

	/*CaCert - Descr: PEM-encoded root CA certificate Default: <nil>
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*ClientSessionCacheSize - Descr: capacity of the etcd client tls client cache Default: <nil>
*/
	ClientSessionCacheSize interface{} `yaml:"client_session_cache_size,omitempty"`

	/*RequireSsl - Descr: boolean to configure ssl connections with the etcd cluster Default: true
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

	/*ClientKey - Descr: PEM-encoded client key Default: <nil>
*/
	ClientKey interface{} `yaml:"client_key,omitempty"`

}
