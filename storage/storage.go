package storage

import (
	"app/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Department() DepartmentRepoI
	Employee() EmployeeRepoI
	Role() RoleRepoI
	User() UserRepoI
}

type DepartmentRepoI interface {
	Insert(context.Context, *models.CreateDepartment) (string, error)
	GetByID(context.Context, *models.DepartmentPrimaryKey) (*models.Department, error)
	GetList(context.Context, *models.GetListDepartmentRequest) (*models.GetListDepartmentResponse, error)
	Update(context.Context, *models.UpdateDepartment) error
	Delete(context.Context, *models.DepartmentPrimaryKey) error
}

type EmployeeRepoI interface {
	Insert(context.Context, *models.CreateEmployee) (string, error)
	GetByID(context.Context, *models.EmployeePrimaryKey) (*models.Employee, error)
	GetByDepartmentId(ctx context.Context, req *models.EmployeeDepartmentId) (*models.GetListEmployeeResponse, error)
	GetList(context.Context, *models.GetListEmployeeRequest) (*models.GetListEmployeeResponse, error)
	Update(context.Context, *models.UpdateEmployee) (int64, error)
	Delete(context.Context, *models.EmployeePrimaryKey) error
}

type RoleRepoI interface {
	Insert(context.Context, *models.CreateRole) (string, error)
	GetByID(context.Context, *models.RolePrimaryKey) (*models.Role, error)
	GetList(context.Context, *models.GetListRoleRequest) (*models.GetListRoleResponse, error)
	Update(context.Context, *models.UpdateRole) error
	Delete(context.Context, *models.RolePrimaryKey) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByPKey(ctx context.Context, req *models.UserPrimarKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	// Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	// Delete(ctx context.Context, req *models.UserPrimarKey) error
}
