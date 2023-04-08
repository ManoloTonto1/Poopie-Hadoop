package hadoop

type Product struct {
	URl     string   `json:"url"`
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Reviews []string `json:"reviews"`
}
