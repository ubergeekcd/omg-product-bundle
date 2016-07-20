package proxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*User - Descr: Username to register a route via NATS Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Port - Descr: IP port of Cloud Foundry NATS server Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Password - Descr: Password to register a route via NATS Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Machines - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

}
