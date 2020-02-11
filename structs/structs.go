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
	Id string
	Title string
	Price float64
	Amount int
}
