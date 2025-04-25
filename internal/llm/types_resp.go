package llm

type ProductRespType struct {
	Rank        int      `json:"rank"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ProductURL  string   `json:"product_url"`
	Categories  []string `json:"categories"`
}
