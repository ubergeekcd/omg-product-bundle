package swarm_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SwarmAgentJob struct {

	/*Docker - Descr: TCP port where Docker daemon will listen to (if not set, TCP will not be available) Default: 4243
*/
	Docker *Docker `yaml:"docker,omitempty"`

	/*Env - Descr: HTTPS proxy that Docker should use Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

	/*SwarmAgent - Descr: Sets the expiration of an ephemeral node Default: 60s
*/
	SwarmAgent *SwarmAgent `yaml:"swarm_agent,omitempty"`

	/*Swarm - Descr: Swarm discovery string (ie: consul://<ip>/<path>, etcd://<ip1>,<ip2>/<path>, zk://<ip1>,<ip2>/<path>, [nodes://]<ip1>,<ip2> Default: <nil>
*/
	Swarm *Swarm `yaml:"swarm,omitempty"`

}