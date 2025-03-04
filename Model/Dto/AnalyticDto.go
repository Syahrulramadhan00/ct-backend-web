package Dto

type ChartData struct {
	Labels   []string   `json:"labels"`  
	Datasets []Dataset `json:"datasets"` 
}

type Dataset struct {
	Label string  `json:"label"` 
	Data  []float64   `json:"data"` 
}
