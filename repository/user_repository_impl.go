package repository

import (
	"database/sql"
	"errors"
	"go_fiber/helper"
	"go_fiber/model/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserRepositoryImpl struct{
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(c *fiber.Ctx, tx *sql.Tx, user entity.User) entity.User  {
	query := `
	
	INSERT INTO users (name, email, password, address, verifyCode, phone) VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := tx.ExecContext(c.Context(), query, user.Name, user.Email, user.Password, user.Address, user.VerifyCode, user.Phone)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	
	return user
}

func (repository *UserRepositoryImpl) Update(c *fiber.Ctx, tx *sql.Tx, user entity.User) entity.User   {
	
	query := `
		UPDATE users SET name = ?, email = ?, address = ?, phone = ?, updated_at = ? WHERE id = ? 
	`
	_ , err := tx.ExecContext(c.Context(), query, user.Name, user.Email, user.Address, user.Phone, time.Now(), user.Id)
	helper.PanicIfError(err)
	
	return user
}

func (repository *UserRepositoryImpl) FindAll(c *fiber.Ctx, tx *sql.Tx) []entity.User  {
	
	query := `SELECT id, name, email, address, isVerified, phone, created_at, updated_at FROM users`

	rows, err := tx.QueryContext(c.Context(), query)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []entity.User

	for rows.Next(){
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Address, &user.IsVerified, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	
	return users
}

func (repository *UserRepositoryImpl) FindById(c *fiber.Ctx, tx *sql.Tx, userId int) (entity.User, error)   {
	
	query := `
		SELECT id, name, email, address, phone, created_at, updated_at 
		FROM users
		WHERE id = ?
	`
	
	rows, err := tx.QueryContext(c.Context(), query, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}

	if rows.Next(){
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Address, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	}else{
		return user, errors.New("User Not Found")
	}
}

func (repository *UserRepositoryImpl) Delete(c *fiber.Ctx, tx *sql.Tx, user entity.User) {
	
	query := `
		DELETE FROM users WHERE id = ?
	`
	_, err := tx.ExecContext(c.Context(), query, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByEmail(c *fiber.Ctx, tx *sql.Tx, email string) (entity.User, error) {
	
	query := `
		SELECT name, email, password, verifyCode, isVerified FROM users 
		WHERE email = ?
		LIMIT 1
	`
	rows, err := tx.QueryContext(c.Context(), query, email)
	helper.PanicIfError(err)

	user := entity.User{}

	if rows.Next(){
		err = rows.Scan(&user.Name, &user.Email, &user.Password, &user.VerifyCode, &user.IsVerified)
		helper.PanicIfError(err)
		return user, nil
	}else{
		return user, errors.New("Email Not Registered")
	}
}

func (repository *UserRepositoryImpl) SetStatusIsVerified(c *fiber.Ctx, tx *sql.Tx, email string){
	
	query := `
		UPDATE users SET isVerified = 'y' 
		WHERE email = ?
	`

	_, err := tx.ExecContext(c.Context(), query, email)
	helper.PanicIfError(err)
}