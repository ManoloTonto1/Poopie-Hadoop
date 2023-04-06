package hadoop

type Product struct {
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Reviews []string `json:"reviews"`
}
