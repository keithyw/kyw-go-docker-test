package repositories

import (
	"errors"
	"time"
	"github.com/keithyw/kyw-go-docker-test/database"
	"github.com/keithyw/kyw-go-docker-test/models"
)

type UserRepositoryImpl struct {
	Conn *database.MysqlDB
}

func NewUserRepository(db *database.MysqlDB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) CreateUser(user models.User) (*models.User, error) {
	var newUser models.User
	stmt, err := r.Conn.DB.Prepare("INSERT INTO users(id, username, email, passwd, first_name, last_name, created, modified) values(NULL, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(user.Username, user.Email, user.Passwd, user.FirstName, user.LastName, time.Now().UnixMilli(), time.Now().UnixMilli())
	if err != nil {
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	newUser.Username = user.Username
	newUser.Passwd = user.Passwd
	newUser.Email = user.Email
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Created = user.Created
	newUser.Modified = user.Modified
	newUser.ID = lastId
	return &newUser, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int) error {
	stmt, err := r.Conn.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(id int, user models.User) (*models.User, error) {
	stmt, err := r.Conn.DB.Prepare("UPDATE users SET username = ?, email = ?, first_name = ?, last_name = ?, modified = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(user.Username, user.Email, user.FirstName, user.LastName, time.Now().UnixMilli(), id)
	if err != nil {
		return nil, err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowCount == 0 {
		return nil, errors.New("nothing updated")
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindUserById(id int) (*models.User, error) {
	var user models.User
	stmt, err := r.Conn.DB.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &user.Passwd, &user.FirstName, &user.LastName, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindUserByName(name string) (*models.User, error) {
	var user models.User
	stmt, err := r.Conn.DB.Prepare("SELECT * FROM users WHERE name = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(name).Scan(&user.ID, &user.Username, &user.Email, &user.Passwd, &user.FirstName, &user.LastName, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	stmt, err := r.Conn.DB.Prepare("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Passwd, &user.FirstName, &user.LastName, &user.Created, &user.Modified); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}