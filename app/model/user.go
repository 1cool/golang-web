package model

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name" validate:"required,NameValid"`
	Phone  string `json:"phone"`
	Gender int    `json:"gender"`
}

func (u *User) List() []User {
	return []User{
		{Id: 1, Name: "小梦", Phone: "13512348526", Gender: 1},
		{Id: 2, Name: "金沙江看", Phone: "13512348526", Gender: 1},
		{Id: 3, Name: "萨达萨达", Phone: "13512348526", Gender: 1},
		{Id: 4, Name: "撒打算", Phone: "13512348526", Gender: 1},
	}
}
