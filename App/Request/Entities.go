package Request

type Article struct {
	ArticleId	string	`json:"id,omitempty"`
	Title		string	`json:"title,omitempty"`
	Price  		string	`json:"price,omitempty"`
}
type ArticleCart struct{
	Element		*Article	`json:"article,omitempty"`
	Amount 		int			`json:"quantity,omitempty"`
}

type PostArticle struct{
	CartId		string	`json:"CartId,omitempty"`
	ArticleId	string	`json:"id,omitempty"`
	Title		string	`json:"title,omitempty"`
	Price  		string	`json:"price,omitempty"`
	Amount 		int		`json:"quantity,omitempty"`
}

type TestResponse struct {
	CartId string `json:"CartId,omitempty"`
	Response string `json:"Response,omitempty"`
}

type BadResponseAddItem struct {
}
//Example
type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
  }
  type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
  }