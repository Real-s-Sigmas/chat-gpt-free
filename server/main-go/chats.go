package main_go

type Message struct {
	Id       string `json:"-" db:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Chat_id  string `json:"chat_id"`
}

type Chat struct {
	Id      string `json:"-" db:"id"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	User_id string `json:"user_id"`
}

type User struct {
	Id       string `json:"-" db:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
