package models

type HandBookModel struct {
	BookID string `json:"book_id" bson:"book_id"`
	Title  string `json:"title" bson:"title"`
	Price  int    `json:"price" bson:"price"`
	Stock  int    `json:"stock" bson:"stock"`
}

type HandBookUpdateModel struct {
	Title string `json:"title" bson:"title"`
	Price int    `json:"price" bson:"price"`
	Stock int    `json:"stock" bson:"stock"`
}
