package book

//CreateModel - модель создания книги

type CreateModel struct {
	Name string `json:"name" form:"name"`
}
