package postgresdb

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(p *PostgresDB) *UserRepository {
	return &UserRepository{
		DB: p.DB,
	}
}

func (r *UserRepository) CreateTable() {
	sql := "create table if not exists users (id int not null primary key, name varchar(255), email varchar(255))"
	_, err := r.DB.Exec(context.TODO(), sql)
	if err != nil {
		log.Panic("Cannot create user table", err)
	}
}

func (r *UserRepository) Insert(user User) error {
	_, err := r.DB.Exec(context.TODO(), "insert into users (id,name,email) values ($1,$2,$3)", user.ID, user.Name, user.Email)
	return err
}

func (r *UserRepository) Count() int {

	row := r.DB.QueryRow(context.TODO(), "select count(1) from users")
	var result int
	err := row.Scan(&result)
	if err != nil {
		log.Panic("Cannot read the value. ", err)
	}

	return result
}
