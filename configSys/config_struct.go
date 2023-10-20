package configSys

type Config struct {
	Listener ListenerConfig `toml:"listener"`
	Proxy    ProxyConfig    `toml:"proxy"`
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
}

type ListenerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type ProxyConfig struct {
	Rewrite bool   `toml:"rewrite"`
	Pass    string `toml:"pass"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DbName   string `toml:"db_name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
}
