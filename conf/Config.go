package conf

type ServerOut struct {
	Name string `json:"name"`
}

type Config struct {
	ServerOut         `json:"serverout"`
	Logsfolderpath    string `json:"logsfolderpath"`
	Reportsfolderpath string `json:"reportsfolderpath"`
}
