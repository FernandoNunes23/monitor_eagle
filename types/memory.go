package types

type Memory struct {
	Total 	  	string	`json:"total"`
	Used      	string  `json:"used"`
	Free	 	string	`json:"free"`
	Available 	string	`json:"available"`
	Buffers		string	`json:"buffers"`
	Cache		string	`json:"cache"`
}