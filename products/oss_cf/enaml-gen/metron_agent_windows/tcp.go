package metron_agent_windows 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tcp struct {

	/*BatchingBufferFlushIntervalMilliseconds - Descr: The maximum time that a message can stay in the batching buffer before being flushed Default: 100
*/
	BatchingBufferFlushIntervalMilliseconds interface{} `yaml:"batching_buffer_flush_interval_milliseconds,omitempty"`

	/*BatchingBufferBytes - Descr: The number of bytes which can be buffered prior to TCP writes (applies to TLS over TCP) Default: 10240
*/
	BatchingBufferBytes interface{} `yaml:"batching_buffer_bytes,omitempty"`

}