package models

type DepartmentPrimaryKey struct {
	Id string `json:"id"`
}

type CreateDepartment struct {
	D_Number string `json:"d_number"`
	Name     string `json:"name"`
}

type Department struct {
	Id        string `json:"id"`
	D_Number  string `json:"d_number"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateDepartment struct {
	Id       string `json:"id"`
	D_Number string `json:"d_number"`
	Name     string `json:"name"`
}

type UpdateDepartmentSwag struct {
	D_Number string `json:"d_number"`
	Name     string `json:"name"`
}

type GetListDepartmentRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListDepartmentResponse struct {
	Count       int64         `json:"count"`
	Departments []*Department `json:"departments"`
}

type Empty struct{}
