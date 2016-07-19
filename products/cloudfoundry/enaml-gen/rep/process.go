package rep 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Process struct {

	/*User - Descr: User to use while performing a container healthcheck Default: vcap
*/
	User interface{} `yaml:"user,omitempty"`

	/*Dir - Descr: Directory to run the healthcheck process from Default: <nil>
*/
	Dir interface{} `yaml:"dir,omitempty"`

	/*Path - Descr: Path of the command to run to perform a container healthcheck Default: /bin/sh
*/
	Path interface{} `yaml:"path,omitempty"`

	/*Env - Descr: Environment variables to use when running the garden health check Default: <nil>
*/
	Env interface{} `yaml:"env,omitempty"`

	/*Args - Descr: List of command line args to pass to the garden health check process Default: -c, ls > /tmp/test
*/
	Args interface{} `yaml:"args,omitempty"`

}
