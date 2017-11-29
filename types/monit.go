package types

type MonitConfig struct {
	Type 	 string	`json:"type"`
	Resource string	`json:"resource"`
	Time	 string	`json:"time"`
}

type Monit struct {
	Id		uint64	     `json:"id"`
	Name 	string	     `json:"name"`
	Created	string	     `json:"created"`
	Config	MonitConfig  `json:"config"`
}	