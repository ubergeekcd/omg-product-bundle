package streaming_mysql_backup_client 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Credentials struct {

	/*Username - Descr: Username used by backup client to stream a backup from the mysql node Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Password - Descr: Password used by backup client to stream a backup from the mysql node Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

}