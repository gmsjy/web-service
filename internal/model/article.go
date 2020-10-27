package model

//Article 文章模型
type Article struct {
	*Model
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Desc          string `json:"desc"`
	Tags          []Tag  `gorm:"many2many:blog_tags;"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

//Tag 文章标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
