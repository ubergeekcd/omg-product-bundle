package gorouter 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tracing struct {

	/*EnableZipkin - Descr: Enables the addition of the X-B3-Trace-Id header to incoming requests. If the header already exists on the incoming request, it will not be overwritten. Default: false
*/
	EnableZipkin interface{} `yaml:"enable_zipkin,omitempty"`

}