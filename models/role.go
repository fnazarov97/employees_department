package models

type RolePrimaryKey struct {
	Id string `json:"id"`
}

type CreateRole struct {
	R_Title     string `json:"r_title"`
	Discription string `json:"discription"`
}

type Role struct {
	Id          string `json:"id"`
	R_Title     string `json:"r_title"`
	Description string `json:"discription"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateRole struct {
	Id          string `json:"id"`
	R_Title     string `json:"r_title"`
	Discription string `json:"discription"`
}

type UpdateRoleSwag struct {
	R_Title     string `json:"r_title"`
	Discription string `json:"discription"`
}

type GetListRoleRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListRoleResponse struct {
	Count int64   `json:"count"`
	Roles []*Role `json:"roles"`
}
