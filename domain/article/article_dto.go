package article

type Article struct {
	Titre string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
	Date string  `json:"date"`
}
