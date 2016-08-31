package pmysql

const (
	innodbBufferPoolSize       int = 2147483648
	maxConnections             int = 1500
	databaseStartupTimeout     int = 600
	wsrepDebug                 int = 1
	backupServerPort           int = 8081
	seededDBUser                   = "repcanary"
	seededDBName                   = "canary_db"
	adminUsername                  = "root"
	notificationClientUsername     = "mysql-monitoring"
	natsPort                       = "4222"
	natsUser                       = "nats"
)