package admin

import (
	"ca-reservaksin/businesses/admin"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/admin/request"
	"ca-reservaksin/controllers/admin/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService admin.Service
}

func NewAdminController(service admin.Service) *AdminController {
	return &AdminController{
		AdminService: service,
	}
}

func (ctrl *AdminController) Register(c echo.Context) error {
	req := request.Admin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.AdminService.Register(req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate data") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *AdminController) Login(c echo.Context) error {
	req := request.AdminLogin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.AdminService.Login(req.Username, req.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *AdminController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	admin, err := ctrl.AdminService.GetByID(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(admin))
}
