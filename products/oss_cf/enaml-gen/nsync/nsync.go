package nsync 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nsync struct {

	/*ListenerDebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17006
*/
	ListenerDebugAddr interface{} `yaml:"listener_debug_addr,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*LifecycleBundles - Descr: List of lifecycle bundles arguments for different stacks in form 'lifecycle-name:path/to/bundle' Default: [buildpack/cflinuxfs2:buildpack_app_lifecycle/buildpack_app_lifecycle.tgz buildpack/windows2012R2:windows_app_lifecycle/windows_app_lifecycle.tgz docker:docker_app_lifecycle/docker_app_lifecycle.tgz]
*/
	LifecycleBundles interface{} `yaml:"lifecycle_bundles,omitempty"`

	/*ListenAddr - Descr: Address from which nsync serves requests Default: 0.0.0.0:8787
*/
	ListenAddr interface{} `yaml:"listen_addr,omitempty"`

	/*BulkerDebugAddr - Descr: address at which to serve debug info Default: 0.0.0.0:17007
*/
	BulkerDebugAddr interface{} `yaml:"bulker_debug_addr,omitempty"`

	/*Bbs - Descr: capacity of the tls client cache Default: <nil>
*/
	Bbs *Bbs `yaml:"bbs,omitempty"`

	/*Cc - Descr: How long to wait for completion of requests to CC in seconds. Default: 30
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*FileServerUrl - Descr: URL of file server Default: http://file-server.service.cf.internal:8080
*/
	FileServerUrl interface{} `yaml:"file_server_url,omitempty"`

	/*ConsulAgentPort - Descr: local consul agent's port Default: 8500
*/
	ConsulAgentPort interface{} `yaml:"consul_agent_port,omitempty"`

	/*DiegoPrivilegedContainers - Descr: Whether or not to use privileged containers for  buildpack based LRPs and tasks. Containers with a docker-image-based rootfs will continue to always be unprivileged and cannot be changed. Default: false
*/
	DiegoPrivilegedContainers interface{} `yaml:"diego_privileged_containers,omitempty"`

}