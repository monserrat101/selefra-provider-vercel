package vercel_client

type Config struct {
	APIToken string `yaml:"api_token"  mapstructure:"api_token"`
	Team     string `yaml:"team"  mapstructure:"team"`
}
