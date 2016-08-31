package cf_mysql_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type QuotaEnforcer struct {

	/*Password - Descr: The password for the quota-enforcer user Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Pause - Descr: In seconds, the interval that the Quota Enforcer pauses between checks for violators and reformers Default: 1
*/
	Pause interface{} `yaml:"pause,omitempty"`

}