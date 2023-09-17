package module

type Config struct {
	Environment uint8    `yaml:"environment" validate:"gte=1,lte=2"`
	LogLevel    uint32   `yaml:"log_level" validate:"required"`
	Address     string   `yaml:"address" validate:"required"`
	FrontendUrl string   `yaml:"frontend_url" validate:"required"`
	BackendUrl  string   `yaml:"backend_url" validate:"required"`
	Cors        []string `yaml:"cors" validate:"required"`
	MongoUrl    string   `yaml:"mongo_url" validate:"required"`
	MongoDbName string   `yaml:"mongo_db_name" validate:"required"`
	ForwardLink string   `yaml:"forward_link" validate:"required"`
}
