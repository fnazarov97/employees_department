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

// CreateRole godoc
// @ID create_Role
// @Router /role [POST]
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body models.CreateRole true "CreateRoleRequestBody"
// @Success 201 {object} models.Role "GetRoleBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateRole(c *gin.Context) {

	var role models.CreateRole

	err := c.ShouldBindJSON(&role)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Role().Insert(context.Background(), &role)
	if err != nil {
		log.Println("error whiling create Branc:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Role().GetByID(context.Background(), &models.RolePrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Role:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDRole godoc
// @ID get_by_id_Role
// @Router /role/{id} [GET]
// @Summary Get By ID Role
// @Description Get By ID Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Role "GetRoleBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdRole(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Role().GetByID(context.Background(), &models.RolePrimaryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Role:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListRole godoc
// @ID get_list_Role
// @Router /role [GET]
// @Summary Get List Role
// @Description Get List Role
// @Tags Role
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetListRoleResponse "GetRoleListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListRole(c *gin.Context) {
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

	res, err := h.storage.Role().GetList(context.Background(), &models.GetListRoleRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list Role:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateRole godoc
// @ID update_Role
// @Router /role/{id} [PUT]
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Role body models.UpdateRoleSwag true "UpdateRoleRequestBody"
// @Success 202 {object} models.Role "UpdateRoleBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateRole(c *gin.Context) {

	var (
		role models.UpdateRole
	)

	role.Id = c.Param("id")

	err := c.ShouldBindJSON(&role)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Role().Update(context.Background(), &models.UpdateRole{
		Id:          role.Id,
		R_Title:     role.R_Title,
		Discription: role.Discription,
	})

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Role().GetByID(context.Background(), &models.RolePrimaryKey{
		Id: role.Id,
	})
	if err != nil {
		log.Printf("error whiling get by id Role: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteRole godoc
// @ID delete_Role
// @Router /role/{id} [DELETE]
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteRoleBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Role().Delete(context.Background(), &models.RolePrimaryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Role:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "Role deleted")
}
