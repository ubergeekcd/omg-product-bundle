package replication_canary 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ReplicationCanary struct {

	/*CanaryPassword - Descr: Canary user password Default: <nil>
*/
	CanaryPassword interface{} `yaml:"canary_password,omitempty"`

	/*NotificationsClientSecret - Descr: Secret of the UAA client used to send mysql monitoring notifications. This will be used when creating the client. Default: <nil>
*/
	NotificationsClientSecret interface{} `yaml:"notifications_client_secret,omitempty"`

	/*WriteReadDelay - Descr: This property configures how long the canary waits to read the chirps after writing (in seconds) Default: 1
*/
	WriteReadDelay interface{} `yaml:"write_read_delay,omitempty"`

	/*SwitchboardCount - Descr: The number of switchboard proxies Default: <nil>
*/
	SwitchboardCount interface{} `yaml:"switchboard_count,omitempty"`

	/*UaaAdminClientSecret - Descr: Secret of the UAA client used to create the notifications client Default: <nil>
*/
	UaaAdminClientSecret interface{} `yaml:"uaa_admin_client_secret,omitempty"`

	/*PollFrequency - Descr: Configure how frequently the canary polls the cluster for replication failure (in seconds) Default: 30
*/
	PollFrequency interface{} `yaml:"poll_frequency,omitempty"`

	/*NotificationsClientUsername - Descr: Username of the UAA client used to send mysql monitoring notifications. This will be created using the uaa admin client. Default: mysql-monitoring
*/
	NotificationsClientUsername interface{} `yaml:"notifications_client_username,omitempty"`

	/*CanaryUsername - Descr: Canary username Default: repcanary
*/
	CanaryUsername interface{} `yaml:"canary_username,omitempty"`

	/*MysqlPort - Descr: Database port for contacting mysql Default: 3306
*/
	MysqlPort interface{} `yaml:"mysql_port,omitempty"`

	/*CanaryDatabase - Descr: Canary database Default: canary_db
*/
	CanaryDatabase interface{} `yaml:"canary_database,omitempty"`

	/*UaaAdminClientUsername - Descr: Username of the UAA client used to create the notifications client Default: admin
*/
	UaaAdminClientUsername interface{} `yaml:"uaa_admin_client_username,omitempty"`

	/*ClusterIps - Descr: List of IP addresses of servers used to read data from Default: <nil>
*/
	ClusterIps interface{} `yaml:"cluster_ips,omitempty"`

	/*SwitchboardUsername - Descr: Basic Auth username to contact the Switchboard API Default: <nil>
*/
	SwitchboardUsername interface{} `yaml:"switchboard_username,omitempty"`

	/*SwitchboardPassword - Descr: Basic Auth password to contact the Switchboard API Default: <nil>
*/
	SwitchboardPassword interface{} `yaml:"switchboard_password,omitempty"`

}