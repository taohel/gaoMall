package conf

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

type Redis struct {
	Write RedisConf `yaml:"write"`
	Read  RedisConf `yaml:"read"`
}
