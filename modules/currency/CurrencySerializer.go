package currency_converter
type GetAllRatesResponse struct {
	Rates []Rate `json:"rates"`
}

type Query struct {
	From string `json:"from"`
	To string `json:"to"`
	Amount float64 `json:"amount"`
}

type ConvertResponse struct {
	Query Query `json:"query"`
	Result float64 `json:"result"`
}