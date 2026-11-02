package users_repository

import postgres_pool "actualiteco/actualite/internal/core/repository/postgres"

type UsersRepository struct {
	ConnPool postgres_pool.Pool
}

func NewAuthRepository(conn postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		ConnPool: conn,
	}
}
