package service_metrics 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ServiceMetrics struct {

	/*ExecutionIntervalSeconds - Descr: Interval to repeatedly obtain and emit metrics, in seconds Default: 60
*/
	ExecutionIntervalSeconds interface{} `yaml:"execution_interval_seconds,omitempty"`

	/*MetricsCommandArgs - Descr: Arguments to be passed to the metrics command (see service_metrics.metrics_command) Default: []
*/
	MetricsCommandArgs interface{} `yaml:"metrics_command_args,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose mode Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

	/*Origin - Descr: Used when initialising dropsonde. Should be set to something descriptive of the deployment (e.g. p-redis) Default: <nil>
*/
	Origin interface{} `yaml:"origin,omitempty"`

	/*MetricsCommand - Descr: Command to obtain metrics in JSON format Default: <nil>
*/
	MetricsCommand interface{} `yaml:"metrics_command,omitempty"`

	/*MonitDependencies - Descr: Array of monit service dependencies Default: []
*/
	MonitDependencies interface{} `yaml:"monit_dependencies,omitempty"`

}