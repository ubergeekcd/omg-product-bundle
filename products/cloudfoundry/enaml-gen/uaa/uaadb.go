package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaadb struct {

	/*Databases - Descr: The list of databases used in UAA database including tag/name Default: <nil>
*/
	Databases interface{} `yaml:"databases,omitempty"`

	/*DbScheme - Descr: Database scheme for UAA DB Default: <nil>
*/
	DbScheme interface{} `yaml:"db_scheme,omitempty"`

	/*Roles - Descr: The list of database Roles used in UAA database including tag/name/password Default: <nil>
*/
	Roles interface{} `yaml:"roles,omitempty"`

	/*Port - Descr: The UAA database Port Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Address - Descr: The UAA database IP address Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

}