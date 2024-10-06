package store

import (
	"time"

	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
)

type ChatRepository struct{
	store *Store
}

func (r *ChatRepository) GetAll()(*[]model.Chat,error){
	chats := make([]model.Chat,0,10)

	rows,err := r.store.db.Query(
		"SELECT id,user1_id,user2_id,created_at FROM chats")

	if err!= nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var chat model.Chat

		err := rows.Scan(&chat.ID,&chat.User1Id,&chat.User2Id,&chat.CreatedTime)

		if err!= nil{
			return nil,err
		}

		chats = append(chats, chat)
	}

	return &chats,nil
}

func (r *ChatRepository) Create(m *model.Chat)(*model.Chat,error){
	if err := r.store.db.QueryRow(
		"INSERT INTO chats(user1_id,user2_id) VALUES($1,$2) RETURNING id, created_at",
	m.User1Id,m.User2Id).Scan(&m.ID,&m.CreatedTime);err != nil{
		return nil,err
	}

	return m, nil
}

func (r *ChatRepository) GetByChatById(id int)(*model.Chat,error){
	var chat model.Chat

	if err := r.store.db.QueryRow(
		"SELECT id, user1_id, user2_id, created_at FROM chats WHERE id = $1",
	id).Scan(&chat.ID,&chat.User1Id,&chat.User2Id,&chat.CreatedTime);err != nil{
		return nil,err
	}

	return &chat, nil
}

func (r *ChatRepository) GetByChatsUser1Id(user1Id int)(*[]model.Chat,error){
	chats := make([]model.Chat,0,10)

	rows,err := r.store.db.Query(
		"SELECT id,user1_id,user2_id,created_at FROM chats WHERE user1_id = $1",
		user1Id)

	if err!= nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var chat model.Chat

		err := rows.Scan(&chat.ID,&chat.User1Id,&chat.User2Id,&chat.CreatedTime)

		if err!= nil{
			return nil,err
		}

		chats = append(chats, chat)
	}

	return &chats,nil
}

func (r *ChatRepository) GetByChatsUser2Id(user2Id int)(*[]model.Chat,error){
	chats := make([]model.Chat,0,10)

	rows,err := r.store.db.Query(
		"SELECT id,user1_id,user2_id,created_at FROM chats WHERE user2_id = $1",
		user2Id)

	if err!= nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var chat model.Chat

		err := rows.Scan(&chat.ID,&chat.User1Id,&chat.User2Id,&chat.CreatedTime)

		if err!= nil{
			return nil,err
		}

		chats = append(chats, chat)
	}

	return &chats,nil
}

func (r *ChatRepository) DeleteById(id int) error{
	if _,err := r.store.db.Exec(
		"DELETE FROM chats WHERE id = $1;",
		id); err != nil{
			return err
		}
	return nil
}


func (r *ChatRepository) ModifyCreatedTimeToCurrent(id int) error{
	if _,err := r.store.db.Exec(
		"UPDATE chats SET created_at = $1 WHERE id = $2",
		time.Now(), id); err != nil{
			return err
		}
	return nil
}
