package cf_containers_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*AllocateDockerHostPorts - Descr: Allocate Docker host ports when creating a container Default: true
*/
	AllocateDockerHostPorts interface{} `yaml:"allocate_docker_host_ports,omitempty"`

	/*Unicorn - Descr: Unicorn worker processes. If the number of worker processes > 1 then you must disable the broker.allocate_docker_host_ports property Default: 1
*/
	Unicorn *Unicorn `yaml:"unicorn,omitempty"`

	/*MaxContainers - Descr: Max number of containers Default: 0
*/
	MaxContainers interface{} `yaml:"max_containers,omitempty"`

	/*Username - Descr: Broker's basic auth username Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Host - Descr: Host used to register the broker Default: <nil>
*/
	Host interface{} `yaml:"host,omitempty"`

	/*SessionExpiry - Descr: Session expiry time of the session Default: 86400
*/
	SessionExpiry interface{} `yaml:"session_expiry,omitempty"`

	/*CookieSecret - Descr: A unique secret key, used to sign sessions Default: <nil>
*/
	CookieSecret interface{} `yaml:"cookie_secret,omitempty"`

	/*SslEnabled - Descr: Determines use of https in dashboard url and in callback uri for calls to UAA Default: false
*/
	SslEnabled interface{} `yaml:"ssl_enabled,omitempty"`

	/*FetchImages - Descr: Fetch new/updated container images on restart Default: true
*/
	FetchImages interface{} `yaml:"fetch_images,omitempty"`

	/*ExternalIp - Descr: External IP address used to register the broker Default: <nil>
*/
	ExternalIp interface{} `yaml:"external_ip,omitempty"`

	/*SkipSslValidation - Descr: Determines whether dashboard verifies SSL certificates when communicating with Cloud Controller and UAA Default: true
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

	/*Password - Descr: Broker's basic auth password Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*DockerUrl - Descr: Docker URL (IP/Socket) Default: unix:///var/vcap/sys/run/docker/docker.sock
*/
	DockerUrl interface{} `yaml:"docker_url,omitempty"`

	/*UpdateContainers - Descr: Restart all containers with latest configuration/image on restart Default: true
*/
	UpdateContainers interface{} `yaml:"update_containers,omitempty"`

	/*Services - Descr: Services and plans offered by the broker Default: <nil>
*/
	Services interface{} `yaml:"services,omitempty"`

}