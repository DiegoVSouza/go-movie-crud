package dto

type MovieInputDTO struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Writer   string  `json:"writer"`
	Link     string  `json:"link"`
	Duration float64 `json:"duration"`
}

type MovieOutputDTO struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Writer   string  `json:"writer"`
	Link     string  `json:"link"`
	Duration float64 `json:"duration"`
}

type MovieDeleteDTO struct {
	ID string `json:"id"`
}
