package models

// Account defines a server account
type Account struct {
	ID       int64
	Name     string
	Password string
	Type     int
	Premdays int
	Lastday  int64
	Email    string
	Creation int64
}
