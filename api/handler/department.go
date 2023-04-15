package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"app/models"

	"github.com/gin-gonic/gin"
)

// CreateDepartment godoc
// @ID create_Department
// @Router /department [POST]
// @Summary Create Department
// @Description Create Department
// @Tags Department
// @Accept json
// @Produce json
// @Param Department body models.CreateDepartment true "CreateDepartmentRequestBody"
// @Success 201 {object} models.Department "GetDepartmentBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateDepartment(c *gin.Context) {

	var department models.CreateDepartment

	err := c.ShouldBindJSON(&department)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Department().Insert(context.Background(), &department)
	if err != nil {
		log.Println("error whiling create department:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id department:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDDepartment godoc
// @ID get_by_id_department
// @Router /department/{id} [GET]
// @Summary Get By ID Department
// @Description Get By ID Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Department "GetDepartmentBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdDepartment(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id department:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListDepartment godoc
// @ID get_list_Department
// @Router /department [GET]
// @Summary Get List Department
// @Description Get List Department
// @Tags Department
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListDepartmentResponse "GetDepartmentListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListDepartment(c *gin.Context) {
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

	res, err := h.storage.Department().GetList(context.Background(), &models.GetListDepartmentRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list department:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateDepartment godoc
// @ID update_Department
// @Router /department/{id} [PUT]
// @Summary Update Department
// @Description Update Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Department body models.UpdateDepartmentSwag true "UpdateDepartmentRequestBody"
// @Success 202 {object} models.Department "UpdateDepartmentBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateDepartment(c *gin.Context) {

	var (
		department models.UpdateDepartment
	)

	department.Id = c.Param("id")

	err := c.ShouldBindJSON(&department)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Department().Update(context.Background(), &models.UpdateDepartment{
		Id:       department.Id,
		D_Number: department.D_Number,
		Name:     department.Name,
	})

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{
		Id: department.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteDepartment godoc
// @ID delete_Department
// @Router /department/{id} [DELETE]
// @Summary Delete Department
// @Description Delete Department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteDepartmentBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Department().Delete(context.Background(), &models.DepartmentPrimaryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  department:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "department deleted")
}
