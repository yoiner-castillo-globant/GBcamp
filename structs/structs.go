package structs

//Element _
type Element struct {
	Id    string
	Title string
	Price float64
}

type ApiStruct struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

type ResponseStruct struct {
	Id     string
	Title  string
	Price  float64
	Amount int
}

type ResponseNewCartWs struct {
	CartId string `json:"CartId"`
}

type BadResponseAddItem struct {
	Response string `json:"Response"`
}

type RequestAddItem struct {
	CartID    string `json:"CartId"`
	ArticleId string `json:"ArticleId"`
}
