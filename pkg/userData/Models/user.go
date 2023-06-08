package Model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	DOB     string `json:"dob"`
	Address string `json:"address"`
	Country string `json:"country"`
}



