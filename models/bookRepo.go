package models

type RepoBookModel struct {
	BookID string `json:"book_id" bson:"book_id"`
	Title  string `json:"title" bson:"title"`
	Price  int    `json:"price" bson:"price"`
	Stock  int    `json:"stock" bson:"stock"`
}

type RepoBookUpdateModel struct {
	Title string `json:"title" bson:"title,omitempty"`
	Price int    `json:"price" bson:"price,omitempty"`
	Stock int    `json:"stock" bson:"stock,omitempty"`
}
