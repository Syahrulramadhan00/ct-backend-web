package Model

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Place       string `json:"place"`
	Description string `json:"description"`
	Telephone   string `json:"telephone"`
}
