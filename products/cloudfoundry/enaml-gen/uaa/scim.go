package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Scim struct {

	/*UseridsEnabled - Descr: Enables the endpoint `/ids/Users` that allows consumers to translate user ids to name Default: true
*/
	UseridsEnabled interface{} `yaml:"userids_enabled,omitempty"`

	/*User - Descr: If true override users defined in uaa.scim.users found in the database. Default: true
*/
	User *ScimUser `yaml:"user,omitempty"`

	/*Users - Descr: A list of users to be bootstrapped with authorities.
Each entry supports the following format:
  Short Pipe: username|password|comma,separated,groups
  Long Pipe: username|password|email|firstName|lastName|comma,separated,groups|origin
  Short OpenStruct:
    - name: username
      password: password
      groups:
        - group1
        - group2
  Long OpenStruct:
    - name: username
      password: password
      groups:
        - group1
        - group2
      firstName: first name
      lastName: lastName
      email: email
      origin: origin-value - most commonly uaa
 Default: <nil>
*/
	Users interface{} `yaml:"users,omitempty"`

	/*ExternalGroups - Descr: A list of external group mappings. Pipe delimited. A value may look as '- internal.read|cn=developers,ou=scopes,dc=test,dc=com' Default: <nil>
*/
	ExternalGroups interface{} `yaml:"external_groups,omitempty"`

	/*Groups - Descr: Contains a hash of group names and their descriptions. These groups will be added to the UAA database for the default zone but not associated with any user.
Example:
  uaa:
    scim:
      groups:
        my-test-group: 'My test group description'
        another-group: 'Another group description'
Deprecated format(still supported, but may be removed in the future): 
Comma separated list of groups that should be added to the UAA db, but not assigned to a user by default.
 Default: <nil>
*/
	Groups interface{} `yaml:"groups,omitempty"`

}
