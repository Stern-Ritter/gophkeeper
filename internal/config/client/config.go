package client

// ClientConfig holds the configuration for the client.
type ClientConfig struct {
	ServerURL    string // ServerURL is the URL of the server.
	BuildVersion string // Version of the application
	BuildDate    string // Application build date
	LoggerLvl    string // The logging level
}
