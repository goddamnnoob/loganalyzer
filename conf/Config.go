package conf

type Config struct {
	ServerOut struct {
		name string `name:"name"`
	} `json:"serverout"`
	Logsfolderpath string `json:"logsfolderpath"`
}
