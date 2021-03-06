package service_metrics 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ServiceMetricsJob struct {

	/*Etcd - Descr: Number of concurrent requests to ETCD Default: 10
*/
	Etcd *Etcd `yaml:"etcd,omitempty"`

	/*ServiceMetrics - Descr: Interval to repeatedly obtain and emit metrics, in seconds Default: 60
*/
	ServiceMetrics *ServiceMetrics `yaml:"service_metrics,omitempty"`

	/*Nats - Descr: Password for cc client to connect to NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*SyslogDaemonConfig - Descr: Transport to be used when forwarding logs (tcp|udp|relp). Default: tcp
*/
	SyslogDaemonConfig *SyslogDaemonConfig `yaml:"syslog_daemon_config,omitempty"`

	/*MetronAgent - Descr: Incoming port for dropsonde log messages Default: 3457
*/
	MetronAgent *MetronAgent `yaml:"metron_agent,omitempty"`

}