package config

type App struct {
	Name    string
	Version float64
}

type Server struct {
	Port string
	Host string
}

type Config struct {
	App    App
	Server Server
}

func CreateConfiguration() Config {
	return Config{
		App: App{
			"Crypto Garage Project",
			1.0,
		},
		Server: Server{
			Port: ":3000",
			Host: "localhost",
		},
	}
}
