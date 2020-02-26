package Request

type Article struct {
	ArticleId	string	`json:"id,omitempty" bson:"id"`
	Title		string	`json:"title,omitempty" bson:"title"`
	Price  		string	`json:"price,omitempty" bson:"price"`
}
type ArticleCart struct{
	Element		Article	`json:"article,omitempty" bson:"article"`
	Amount 		int		`json:"quantity,omitempty" bson:"quantity"`
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


