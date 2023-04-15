package models

type EmployeePrimaryKey struct {
	Id string `json:"id"`
}

type EmployeeDepartmentId struct {
	DepartmentId string `json:"department_id"`
}

type CreateEmployee struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	RoleId       string `json:"role_id"`
	DepartmentId string `json:"department_id"`
}

type Employee struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	RoleId       string `json:"role_id"`
	DepartmentId string `json:"department_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateEmployee struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	RoleId       string `json:"role_id"`
	DepartmentId string `json:"department_id"`
}

type UpdateEmployeeSwag struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	RoleId       string `json:"role_id"`
	DepartmentId string `json:"department_id"`
}

type GetListEmployeeRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListEmployeeResponse struct {
	Count     int64       `json:"count"`
	Employees []*Employee `json:"employees"`
}
