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


