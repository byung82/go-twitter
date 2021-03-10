package twitter

type Config struct {
	BearerToken string
	EndPoint    string
}

func NewConfig(bearerToken, endPoint string) *Config {
	return &Config{
		bearerToken,
		endPoint,
	}
}
