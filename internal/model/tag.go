package model

//Tag 文章标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
