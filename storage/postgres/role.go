package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type RoleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

func (r *RoleRepo) Insert(ctx context.Context, role *models.CreateRole) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO role (
			id,
			r_title,
			description,
			updated_at
		) VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		role.R_Title,
		role.Discription,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *RoleRepo) GetByID(ctx context.Context, req *models.RolePrimaryKey) (*models.Role, error) {

	var (
		id          sql.NullString
		r_title     sql.NullString
		description sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query := `
		SELECT
			id,
			r_title,
			description,
			created_at,
			updated_at
		FROM role
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&r_title,
		&description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Role{
		Id:          id.String,
		R_Title:     r_title.String,
		Description: description.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, err
}

func (r *RoleRepo) GetList(ctx context.Context, req *models.GetListRoleRequest) (*models.GetListRoleResponse, error) {

	var (
		resp   models.GetListRoleResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			r_title,
			description,
			created_at,
			updated_at
		FROM role
	`

	if search != "" {
		search = fmt.Sprintf("where r_title like  '%s%s' ", req.Search, "%")
		query += search
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return &models.GetListRoleResponse{}, err
	}

	var (
		id          sql.NullString
		r_title     sql.NullString
		description sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&r_title,
			&description,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return &models.GetListRoleResponse{}, err
		}

		role := models.Role{
			Id:          id.String,
			R_Title:     r_title.String,
			Description: description.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}

		resp.Roles = append(resp.Roles, &role)
	}

	return &resp, nil
}

func (r *RoleRepo) Update(ctx context.Context, role *models.UpdateRole) error {
	query := `
		UPDATE
			role
		SET
			r_title = $2,
			description = $3,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		role.Id,
		role.R_Title,
		role.Discription,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RoleRepo) Delete(ctx context.Context, req *models.RolePrimaryKey) error {

	_, err := r.db.Exec(ctx, "delete from role where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
