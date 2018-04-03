package propublica

const (
	apiEndpoint = "https://api.propublica.org"
)

type Config struct {
	APIKey string
}

type Client struct {
	config *Config
}

func NewClient(cfg *Config) (*Client, error) {
	return &Client{
		config: cfg,
	}, nil
}
