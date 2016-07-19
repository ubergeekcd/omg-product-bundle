package cf_redis_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CfRedisBrokerJob struct {

	/*Redis - Descr: Set TCP_NODELAY on the slave socket after SYNC. Default: no
*/
	Redis *Redis `yaml:"redis,omitempty"`

	/*Cf - Descr: The password to use when authenticating with NATS Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

}
