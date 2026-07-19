package repository

import (
	"context"
	"database/sql"
	"kunalavghade/learning-go/section-10/6-repository/model"
	"log"
)

type UserRepo interface {
	CreateUser(name, email, password, avatar string) (int, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUsers() ([]*model.User, error)
}

type SQLUserRepo struct {
	db *sql.DB
}

func NewSQLUserRepo(db *sql.DB) UserRepo {
	return &SQLUserRepo{
		db: db,
	}
}

func (r *SQLUserRepo) CreateUser(name, email, password, avatar string) (int, error) {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	query := `insert into user(name, email, hashed_password) values (?, ?, ?)`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, password)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	id, _ := result.LastInsertId()

	profileStmt, err := tx.PrepareContext(ctx, `insert into profile(user_id, avatar) values (?,?)`)
	if err != nil {
		log.Fatal(err)
		err = tx.Rollback()
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(id, avatar)
	if err != nil {
		err = tx.Rollback()
		log.Fatal(err)
		return 0, err
	}
	return int(id), tx.Commit()
}

func (r *SQLUserRepo) GetUserByEmail(email string) (*model.User, error) {
	query := `
		SELECT 
			u.id, u.name, u.email, u.hashed_password, u.created_at, p.user_id, p.avatar
		FROM user u
		JOIN profile p
		ON u.id = p.user_id
		WHERE u.email = ?`
	row := r.db.QueryRow(query, email)

	var usr model.User

	err := row.Scan(
		&usr.ID,
		&usr.Name,
		&usr.Email,
		&usr.Password,
		&usr.CreatedAt,
		&usr.Profile.UserID,
		&usr.Profile.Avatar,
	)

	if err != nil {
		log.Fatal(err)
	}
	return &usr, nil
}

func (r *SQLUserRepo) GetUsers() ([]*model.User, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM user`
	row, err := r.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var users []*model.User
	for row.Next() {
		var usr model.User
		err := row.Scan(
			&usr.ID,
			&usr.Name,
			&usr.Email,
			&usr.Password,
			&usr.CreatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &usr)
	}
	return users, nil
}
