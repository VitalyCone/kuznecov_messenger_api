package store

import (
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
)

type ChatMessageRepository struct {
	store *Store
}

func (r *ChatMessageRepository) Get(messageId int) (*model.ChatMessage, error) {
	var chatMes model.ChatMessage

	chatMes.ID = messageId

	if err := r.store.db.QueryRow(
		"SELECT chat_id, user_id, text, created_at FROM chats WHERE id = $1",
		messageId).Scan(&chatMes.ChatID, &chatMes.UserID, &chatMes.Text, &chatMes.CreatedTime); err != nil {
		return nil, err
	}

	return &chatMes, nil
}

func (r *ChatMessageRepository) GetByChatId(chatId int) (*[]model.ChatMessage, error) {
	var messages []model.ChatMessage

	rows, err := r.store.db.Query("SELECT id, user_id, text, created_at FROM chat_messages WHERE chat_id = $1", chatId)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var message model.ChatMessage

		err = rows.Scan(&message.ID,&message.UserID,&message.Text,&message.CreatedTime)

		if err!= nil{
			return nil,err
		}
		message.ChatID = chatId

		messages = append(messages,message)
	}

	return &messages, nil
}

func (r *ChatMessageRepository) Create(m *model.ChatMessage) (*model.ChatMessage, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO chat_messages(chat_id,user_id,text) VALUES($1,$2,$3) RETURNING id, created_at",
		m.ChatID, m.UserID,m.Text).Scan(&m.ID, &m.CreatedTime); err != nil {
		return nil, err
	}

	return m, nil
}

func (r *ChatMessageRepository) Delete(messageId int) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM chat_messages WHERE id = $1;",
		messageId); err != nil {
		return err
	}
	return nil
}

// func (r *ChatMessageRepository) Modify(*model.ChatMessage) error {
// 	//...
// }
