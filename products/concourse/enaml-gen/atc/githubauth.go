package atc 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type GithubAuth struct {

	/*Authorize - Descr: An array of different criteria to check for when authorizing a GitHub
user. If empty, GitHub authorization is effectively disabled.
 Default: []
*/
	Authorize interface{} `yaml:"authorize,omitempty"`

	/*ApiUrl - Descr: Override default API endpoint URL for Github Enterprise. Must end in a
trailing slash.
 Default: <nil>
*/
	ApiUrl interface{} `yaml:"api_url,omitempty"`

	/*ClientSecret - Descr: GitHub client secret to use for OAuth.

The application must be configured with its callback URL as
`{external_url}/auth/github/callback` (replacing `{external_url}`
with the actual value).
 Default: 
*/
	ClientSecret interface{} `yaml:"client_secret,omitempty"`

	/*TokenUrl - Descr: Override default access token endpoint for Github Enterprise.
 Default: <nil>
*/
	TokenUrl interface{} `yaml:"token_url,omitempty"`

	/*ClientId - Descr: GitHub client ID to use for OAuth.

The application must be configured with its callback URL as
`{external_url}/auth/github/callback` (replacing `{external_url}`
with the actual value).
 Default: 
*/
	ClientId interface{} `yaml:"client_id,omitempty"`

	/*AuthUrl - Descr: Override default OAuth endpoint for Github Enterprise.
 Default: <nil>
*/
	AuthUrl interface{} `yaml:"auth_url,omitempty"`

}
