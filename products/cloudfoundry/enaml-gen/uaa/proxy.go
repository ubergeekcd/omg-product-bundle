package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Proxy struct {

	/*Servers - Descr: Array of the router IPs acting as the first group of HTTP/TCP backends. These will be added to the proxy_ips_regex as exact matches. When using spiff, these will be router_z1 and router_z2 static IPs from cf-jobs.yml Default: []
*/
	Servers interface{} `yaml:"servers,omitempty"`

}
