package dedicated_node 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Snapshotting struct {

	/*RdbChecksum - Descr: RDB files created with checksum disabled have a checksum of zero that will tell the loading code to skip the check. Default: yes
*/
	RdbChecksum interface{} `yaml:"rdb_checksum,omitempty"`

	/*RdbCompression - Descr: Compress string objects using LZF when dump .rdb databases Default: yes
*/
	RdbCompression interface{} `yaml:"rdb_compression,omitempty"`

	/*StopWritesOnBgsaveError - Descr: This will make the user aware (in an hard way) that data is not persisting on disk properly Default: yes
*/
	StopWritesOnBgsaveError interface{} `yaml:"stop_writes_on_bgsave_error,omitempty"`

	/*Save - Descr: save <seconds> <changes>; Will save the DB if both the given number of seconds and the given number of write operations against the DB occurred. Default: [900 1 300 10 60 10000]
*/
	Save interface{} `yaml:"save,omitempty"`

}
