package exception

type Exception struct {
	Name         string `json:"name"`
	FirstLine    string `json:"firstline"`
	First10Lines string `json:"first10lines"`
	Time         string `json:"time"`
}
