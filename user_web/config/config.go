package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	Signingkey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name          string        `mapstructure:"name"`
	Port          int           `mapstructure:"port"`
	Usersrvconfig UserSrvConfig `mapstructure:"user_srv"`
	JWTInfo       JWTConfig     `mapstructure:"jwt"`
}
