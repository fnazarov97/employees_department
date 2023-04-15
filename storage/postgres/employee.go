package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/models"
	"app/pkg/helper"
)

type EmployeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) *EmployeeRepo {
	return &EmployeeRepo{
		db: db,
	}
}

func (r *EmployeeRepo) Insert(ctx context.Context, employee *models.CreateEmployee) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO employee (
			id,
			f_name,
			l_name,
			address,
			phone,
			role_id,
			department_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		employee.FirstName,
		employee.LastName,
		employee.Address,
		employee.PhoneNumber,
		employee.RoleId,
		employee.DepartmentId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *EmployeeRepo) GetByID(ctx context.Context, req *models.EmployeePrimaryKey) (*models.Employee, error) {

	var (
		id           sql.NullString
		firstName    sql.NullString
		lastName     sql.NullString
		address      sql.NullString
		phoneNumber  sql.NullString
		RoleId       sql.NullString
		DepartmentId sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query := `
		SELECT
			id,
			f_name,
			l_name,
			address,
			phone,
			role_id,
			department_id,
			created_at,
			updated_at
		FROM employee
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&address,
		&phoneNumber,
		&RoleId,
		&DepartmentId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	resp := &models.Employee{
		Id:           id.String,
		FirstName:    firstName.String,
		LastName:     lastName.String,
		Address:      address.String,
		PhoneNumber:  phoneNumber.String,
		RoleId:       RoleId.String,
		DepartmentId: DepartmentId.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}

	return resp, err
}

func (r *EmployeeRepo) GetByDepartmentId(ctx context.Context, req *models.EmployeeDepartmentId) (*models.GetListEmployeeResponse, error) {
	var resp = &models.GetListEmployeeResponse{}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			f_name,
			l_name,
			address,
			phone,
			role_id,
			department_id,
			created_at,
			updated_at
		FROM employee WHERE department_id = $1
	`
	rows, err := r.db.Query(ctx, query, req.DepartmentId)

	if err != nil {
		return nil, err
	}
	var (
		id           sql.NullString
		firstName    sql.NullString
		lastName     sql.NullString
		address      sql.NullString
		phoneNumber  sql.NullString
		roleId       sql.NullString
		departmentId sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&address,
			&phoneNumber,
			&roleId,
			&departmentId,
			&createdAt,
			&updatedAt,
		)
		fmt.Println(resp)

		resp.Employees = append(resp.Employees, &models.Employee{
			Id:           id.String,
			FirstName:    firstName.String,
			LastName:     lastName.String,
			Address:      address.String,
			PhoneNumber:  phoneNumber.String,
			RoleId:       roleId.String,
			DepartmentId: departmentId.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return resp, err
}

func (r *EmployeeRepo) GetList(ctx context.Context, req *models.GetListEmployeeRequest) (*models.GetListEmployeeResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = ""
		resp   = &models.GetListEmployeeResponse{}
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			f_name,
			l_name,
			address,
			phone,
			role_id,
			department_id,
			created_at,
			updated_at
		FROM employee
	`
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	query += offset + " " + limit

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	var (
		id           sql.NullString
		firstName    sql.NullString
		lastName     sql.NullString
		address      sql.NullString
		phoneNumber  sql.NullString
		roleId       sql.NullString
		departmentId sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	for rows.Next() {
		err = rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&address,
			&phoneNumber,
			&roleId,
			&departmentId,
			&createdAt,
			&updatedAt,
		)
		fmt.Println(resp)

		resp.Employees = append(resp.Employees, &models.Employee{
			Id:           id.String,
			FirstName:    firstName.String,
			LastName:     lastName.String,
			Address:      address.String,
			PhoneNumber:  phoneNumber.String,
			RoleId:       roleId.String,
			DepartmentId: departmentId.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return resp, err
}

func (r *EmployeeRepo) Update(ctx context.Context, employee *models.UpdateEmployee) (int64, error) {

	var (
		params map[string]interface{}
		query  string
	)

	query = `
		UPDATE
			employee
		SET
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			phone_number = :phone_number,
			role_id = :role_id,
			department_id = :department_id,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"first_name":    employee.FirstName,
		"last_name":     employee.LastName,
		"address":       employee.Address,
		"phone_number":  employee.PhoneNumber,
		"role_id":       employee.RoleId,
		"department_id": employee.DepartmentId,
		"id":            employee.Id,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *EmployeeRepo) Delete(ctx context.Context, req *models.EmployeePrimaryKey) error {

	_, err := r.db.Exec(ctx, "delete from employee where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
