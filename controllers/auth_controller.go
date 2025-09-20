package controllers

import (
	"net/http"

	"github.com/Vis7044/GinCrud2/models"
	"github.com/Vis7044/GinCrud2/services"
	"github.com/Vis7044/GinCrud2/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	serv *services.AuthService
}

func NewAuthService(s *services.AuthService) *AuthController {
	return &AuthController{
		serv: s,
	}
}

func (au *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status:false,Data: "Pleasse Provide user fields"})
		return
	}
	if err := au.serv.Register(c.Request.Context(),&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status:false,Data: "failed to register user"})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[string]{Status: true, Data: "Register Sucessfull"})

}

func (au *AuthController) Login(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status:false,Data: err.Error()})
		return
	}	
	if (input.Email == "" || input.Password == "") {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status:false,Data: "Please provide Email or Password"})
		return
	}
	token, err := au.serv.Login(c.Request.Context(),input.Email,input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status:false,Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[string]{Status: true, Data: token})

}

func (au *AuthController) Profile(c *gin.Context) {
	email := c.GetString("email")
	c.JSON(http.StatusOK, utils.ResponseHandler[string]{Status: true, Data: email})
}
