package controllers

import (
	"net/http"
	"strconv"

	"github.com/Vis7044/GinCrud2/models"
	"github.com/Vis7044/GinCrud2/services"
	"github.com/Vis7044/GinCrud2/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Testcontroller struct {
	service *services.TestService
}
func Init(s *services.TestService) *Testcontroller {
	return &Testcontroller{
		service: s,
	}
}

func (ctrl *Testcontroller) CreateTest(c *gin.Context) {
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
		return
	}

	created, err := ctrl.service.Create(test)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[models.Test]{Status: false,Data:*created})
}

func (ctrl *Testcontroller) GetTest(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	result, err := ctrl.service.GetAll(limit, skip)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[[]models.Test]{Status: false,Data:*result})
}

func (ctrl *Testcontroller) GetOne(c *gin.Context) {
	idParams := c.Param("id")
	ObjId, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
	}
	result, err := ctrl.service.GetOne(ObjId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[models.Test]{Status: false,Data:*result})
}

func (ctrl *Testcontroller) UpdateTest(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
		return
	}
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
		return
	}
	result, err := ctrl.service.UpdateOne(objID, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[*int64]{Status: false,Data:result})
}

func (ctrl *Testcontroller) DeleteTest(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
		return
	}
	result, err := ctrl.service.DeleteOne(objID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseHandler[string]{Status: true, Data: err.Error()})
	}
	c.JSON(http.StatusOK, utils.ResponseHandler[*int64]{Status: false,Data:result})
}