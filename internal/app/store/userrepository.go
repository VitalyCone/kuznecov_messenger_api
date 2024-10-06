package store

import "github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error){
	if err := r.store.db.QueryRow(
		"INSERT INTO users(username) VALUES($1) RETURNING id",
	u.Username).Scan(&u.ID);err != nil{
		return nil,err
	}

	return u, nil
}

func (r *UserRepository) DeleteById(id int) (error){
	if _,err := r.store.db.Exec(
		"DELETE FROM users WHERE id = $1;",
	id); err != nil{
		return err
	}

	return nil
}

func (r *UserRepository) GetUsers() (*[]model.User,error){
	usersArray := make([]model.User,0,10)

	rows,err := r.store.db.Query(
		"SELECT id,username FROM users")

	if err!= nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var user model.User

		err := rows.Scan(&user.ID,&user.Username)

		if err!= nil{
			return nil,err
		}

		usersArray = append(usersArray, user)

	}

	return &usersArray,nil
}

func (r *UserRepository) GetUserByID(id int) (*model.User,error){

	var user model.User

	if err := r.store.db.QueryRow(
		"SELECT id, username FROM users WHERE id = $1",
	id).Scan(&user.ID,&user.Username);err != nil{
		return nil,err
	}

	return &user, nil
}

func (r *UserRepository) ModifyUser(user *model.User) (error){
	if _,err := r.store.db.Exec(
		"UPDATE users SET username = $1 WHERE id = $2 RETURNING id",
		user.Username, user.ID); err != nil{
			return err
		}
	return nil
}
