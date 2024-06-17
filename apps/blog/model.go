package blog

type CreateBlogRequest struct {
	Title    string            `json:"title"`
	Author   string            `json:"author"`
	Content  string            `json:"content"`
	Abstract string            `json:"abstract"`
	Tags     map[string]string `json:"tags"`
}

type Blog struct {
	CreatedBy string `json:"created_by"`

	CreatedAt int64 `json:"created_at"`

	UpdatedAt int64 `json:"updated_at"`

	PublishedAt int64 `json:"published_at"`

	Status Status `json:"status"`

	*CreateBlogRequest
}
