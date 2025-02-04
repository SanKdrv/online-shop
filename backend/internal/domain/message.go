package domain

type Message struct {
	ChatOwnerId int64  `db:"chat_owner_id" json:"chat_id"`
	UserId      int64  `db:"user_id" json:"user_id"`
	Message     string `db:"message" json:"message"`
}

func (Message) TableName() string {
	return "chat_messages"
}
