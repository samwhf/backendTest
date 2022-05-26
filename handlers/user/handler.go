package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samwhf/backendTest/common/responder"
	"github.com/samwhf/backendTest/objects"
	service "github.com/samwhf/backendTest/services/user"
)

// Create User godoc
// @Summary Create Available User
// @Description Create available user from db
// @Tags sample
// @ID create-user
// @Accept  json
// @Produce  json
// @Param   data body objects.User true "user params"
// @Success 200 {object} responder.ResponseBody
// @Failure 400
// @Router /user [post]
func Create(c *gin.Context) {
	var user *objects.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.Create(c, user)
	if err != nil {
		responder.JsonResponse(c, false, err.Error(), nil)
		return
	}
	responder.JsonResponse(c, true, responder.ResourceCreated, gin.H{
		"id": id,
	})
}

// GetUser godoc
// @Summary Get Available User
// @Description Get available user from db
// @Tags sample
// @ID get-user
// @Produce  json
// @Param   id     path    string     true   "user id"
// @Success 200 {object} responder.ResponseBody
// @Failure 200 {object} responder.ResponseBody
// @Router /user/{id} [get]
func Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := fmt.Errorf("id is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := service.Get(c, id)
	if err != nil {
		responder.JsonResponse(c, false, err.Error(), nil)
		return
	}
	responder.JsonResponse(c, true, responder.ResourceFetched, user)
}

// Update User godoc
// @Summary Update Available User
// @Description Update available user from db
// @Tags sample
// @ID update-user
// @Accept  json
// @Produce  json
// @Param   id     path    string     true   "user id"
// @Param   data body objects.User true "user params"
// @Success 200 {object} responder.ResponseBody
// @Failure 400
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := fmt.Errorf("id is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user *objects.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	err := service.Update(c, user)
	if err != nil {
		responder.JsonResponse(c, false, err.Error(), nil)
		return
	}
	responder.JsonResponse(c, true, responder.ResourceUpdated, nil)
}

// Delete User godoc
// @Summary Delete Available User
// @Description Delete available user from db
// @Tags sample
// @ID delete-user
// @Produce  json
// @Param   id     path    string     true   "user id"
// @Success 200 {object} responder.ResponseBody
// @Failure 400
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := fmt.Errorf("id is invalid")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.Delete(c, id)
	if err != nil {
		responder.JsonResponse(c, false, err.Error(), nil)
		return
	}
	responder.JsonResponse(c, true, responder.ResourceDeleted, nil)
}
