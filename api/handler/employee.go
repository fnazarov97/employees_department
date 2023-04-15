package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"app/models"

	"github.com/gin-gonic/gin"
)

// CreateEmployee godoc
// @ID create_employee
// @Router /employee [POST]
// @Summary Create Employee
// @Description Create Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param Employee body models.CreateEmployee true "CreateEmployeeRequestBody"
// @Success 201 {object} models.Employee "GetEmployeeBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateEmployee(c *gin.Context) {

	var employee models.CreateEmployee

	err := c.ShouldBindJSON(&employee)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//check department exists
	_, err = h.storage.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{
		Id: employee.DepartmentId,
	})
	if err != nil {
		log.Println("Department not found:", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	//check role exists
	_, err = h.storage.Role().GetByID(context.Background(), &models.RolePrimaryKey{
		Id: employee.RoleId,
	})
	if err != nil {
		log.Println("Role not found:", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	id, err := h.storage.Employee().Insert(context.Background(), &employee)
	if err != nil {
		log.Println("error whiling create employee:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Employee().GetByID(context.Background(), &models.EmployeePrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id employee:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDEmployee godoc
// @ID get_by_id_employee
// @Router /employee/{id} [GET]
// @Summary Get By ID Employee
// @Description Get By ID Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Employee "GetEmployeeBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdEmployee(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Employee().GetByID(context.Background(), &models.EmployeePrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Employee:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetByDepartmentId godoc
// @ID get_list_employee_by_departmentId
// @Router /employee [GET]
// @Summary Get List Employee By DepartmentId
// @Description Get List Employee By DepartmentId
// @Tags Employee
// @Accept json
// @Produce json
// @Param department_id path string true "department_id"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListEmployeeResponse "GetEmployeeListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByDepartmentId(c *gin.Context) {
	var departmentId = c.Param("department_id")

	res, err := h.storage.Employee().GetByDepartmentId(context.Background(), &models.EmployeeDepartmentId{
		DepartmentId: departmentId,
	})

	if err != nil {
		log.Println("error whiling get list employee:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListEmployee godoc
// @ID get_list_employee
// @Router /employee [GET]
// @Summary Get List Employee
// @Description Get List Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListEmployeeResponse "GetEmployeeListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListEmployee(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := h.storage.Employee().GetList(context.Background(), &models.GetListEmployeeRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list employee:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateEmployee godoc
// @ID update_employee
// @Router /employee/{id} [PUT]
// @Summary Update Employee
// @Description Update Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Employee body models.UpdateEmployeeSwag true "UpdateEmployeeRequestBody"
// @Success 202 {object} models.Employee "UpdateEmployeeBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateEmployee(c *gin.Context) {

	var (
		employee models.UpdateEmployee
	)

	employee.Id = c.Param("id")

	err := c.ShouldBindJSON(&employee)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//check department exists
	_, err = h.storage.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{
		Id: employee.DepartmentId,
	})
	if err != nil {
		log.Println("Department not found:", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	//check role exists
	_, err = h.storage.Role().GetByID(context.Background(), &models.RolePrimaryKey{
		Id: employee.RoleId,
	})
	if err != nil {
		log.Println("Role not found:", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	rowsAffected, err := h.storage.Employee().Update(context.Background(), &employee)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected <= 0 {
		log.Printf("error whiling update: %v", sql.ErrNoRows)
		c.JSON(http.StatusInternalServerError, sql.ErrNoRows.Error())
		return
	}

	resp, err := h.storage.Employee().GetByID(context.Background(), &models.EmployeePrimaryKey{
		Id: employee.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteEmployee godoc
// @ID delete_employee
// @Router /employee/{id} [DELETE]
// @Summary Delete Employee
// @Description Delete Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteEmployeeBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Employee().Delete(context.Background(), &models.EmployeePrimaryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  employee:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, "Employee deleted")
}
