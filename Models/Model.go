package models

//Todo is ...
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//TableName is ... explicitly set tablestruct
func (b *Todo) TableName() string {
	return "todo"
}