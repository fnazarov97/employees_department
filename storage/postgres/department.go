package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type DepartmentRepo struct {
	db *pgxpool.Pool
}

func NewDepartmentRepo(db *pgxpool.Pool) *DepartmentRepo {
	return &DepartmentRepo{
		db: db,
	}
}

func (r *DepartmentRepo) Insert(ctx context.Context, Department *models.CreateDepartment) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO department (
			id,
			d_number,
			name,
			updated_at
		) VALUES ($1, $2, $3, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		Department.D_Number,
		Department.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *DepartmentRepo) GetByID(ctx context.Context, req *models.DepartmentPrimaryKey) (*models.Department, error) {

	var (
		id        sql.NullString
		d_number  sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			id,
			d_number,
			name,
			created_at,
			updated_at
		FROM department
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&d_number,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Department{
		Id:        id.String,
		D_Number:  d_number.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, err
}

func (r *DepartmentRepo) GetList(ctx context.Context, req *models.GetListDepartmentRequest) (*models.GetListDepartmentResponse, error) {

	var (
		resp   models.GetListDepartmentResponse
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		search = req.Search
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			d_number,
			name,
			created_at,
			updated_at
		FROM Department
	`

	if search != "" {
		search = fmt.Sprintf("where name like  '%s%s' ", req.Search, "%")
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
		return &models.GetListDepartmentResponse{}, err
	}

	var (
		id        sql.NullString
		d_number  sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&d_number,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return &models.GetListDepartmentResponse{}, err
		}

		department := models.Department{
			Id:        id.String,
			D_Number:  d_number.String,
			Name:      name.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		}

		resp.Departments = append(resp.Departments, &department)
	}

	return &resp, nil
}

func (r *DepartmentRepo) Update(ctx context.Context, Department *models.UpdateDepartment) error {
	query := `
		UPDATE
			department
		SET
			d_number = $3,
			name = $2,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		Department.Id,
		Department.D_Number,
		Department.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *DepartmentRepo) Delete(ctx context.Context, req *models.DepartmentPrimaryKey) error {

	_, err := r.db.Exec(ctx, "delete from department where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
