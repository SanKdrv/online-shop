package domain

type Message struct {
	ChatOwnerID int64  `db:"chat_owner_id" json:"chat_id"`
	UserID      int64  `db:"user_id" json:"user_id"`
	Message     string `db:"message" json:"message"`
}

func (Message) TableName() string {
	return "chat_messages"
}
