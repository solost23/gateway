package configs

type ServerConfig struct {
	Name         string  `mapstructure:"name"`
	Mode         string  `mapstructure:"mode"`
	TimeLocation string  `mapstructure:"time_location"`
	Addr         string  `mapstructure:"addr"`
	ConfigPath   string  `mapstructure:"config_path"`
	Consul       Consul  `mapstructure:"consul"`
	Routes       []Route `mapstructure:"routes"`
}

type Consul struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Route struct {
	Path   string `mapstructure:"path"`
	Scheme string `mapstructure:"scheme"`
	Host   string `mapstructure:"host"`
}
