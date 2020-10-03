package config

type Midware struct {
	Db Db `mapstructure:"database"  yaml:"database"`
	Redis Redis `mapstructure:"redis"  yaml:"redis"`
}

type Db struct {
	Username     string `mapstructure:"username"  yaml:"username"`
	Password     string `mapstructure:"password" yaml:"password"`
	Path         string `mapstructure:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" yaml:"db-name"`
	Config       string `mapstructure:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

