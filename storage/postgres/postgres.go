package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type Store struct {
	db         *pgxpool.Pool
	department *DepartmentRepo
	employee   *EmployeeRepo
	role       *RoleRepo
	user       *UserRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConn

	pool, err := pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return &Store{
		db:         pool,
		department: NewDepartmentRepo(pool),
		employee:   NewEmployeeRepo(pool),
		role:       NewRoleRepo(pool),
		user:       NewUserRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Department() storage.DepartmentRepoI {

	if s.department == nil {
		s.department = NewDepartmentRepo(s.db)
	}

	return s.department
}

func (s *Store) Employee() storage.EmployeeRepoI {

	if s.employee == nil {
		s.employee = NewEmployeeRepo(s.db)
	}

	return s.employee
}

func (s *Store) Role() storage.RoleRepoI {

	if s.role == nil {
		s.role = NewRoleRepo(s.db)
	}

	return s.role
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
