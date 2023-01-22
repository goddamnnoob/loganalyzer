package conf

type ServerOut struct {
	Name string `json:"name"`
}

type SMIF struct {
	Name string `json:"name"`
}

type Config struct {
	ServerOut         `json:"serverout"`
	SMIF              `json:"SMIF"`
	Logsfolderpath    string `json:"logsfolderpath"`
	Reportsfolderpath string `json:"reportsfolderpath"`
}
