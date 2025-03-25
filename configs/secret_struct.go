package configs

type Secrets struct {
	Mongo Mongo `mapstructure:",squash"`
	Redis Redis `mapstructure:",squash"`
}

type Mongo struct {
	Connection string `mapstructure:"mongo_connection"`
	Options    string `mapstructure:"mongo_options"`
	DbName     string `mapstructure:"mongo_db_name"`
	IsDebug    bool   `mapstructure:"mongo_is_debug"`
}

type Redis struct {
	Host         string `mapstructure:"redis_host"`
	Port         int    `mapstructure:"redis_port"`
	Password     string `mapstructure:"redis_password"`
	InstanceName string `mapstructure:"redis_instance_name"`
}
