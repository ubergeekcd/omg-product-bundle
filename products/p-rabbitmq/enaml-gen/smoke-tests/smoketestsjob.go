package smoke_tests 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SmokeTestsJob struct {

	/*Broker - Descr: Service name displayed in the catalog metadata Default: p-rabbitmq
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*SmokeTests - Descr: App domain for Cloud Foundry. Defaults to cf.domain. Default: 
*/
	SmokeTests *SmokeTests `yaml:"smoke_tests,omitempty"`

	/*Cf - Descr: Full URL of Cloud Foundry API Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

}