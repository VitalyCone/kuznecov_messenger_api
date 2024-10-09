package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apiserver/dtos"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary create user
// @Schemes
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Param username body dtos.CreateUserDto true "Account username"
// @Success 200 {string} Helloworld
// @Router /user [post]
func (ep *Endpoints)CreateUser(g *gin.Context){
	var newUser dtos.CreateUserDto

	err := g.BindJSON(&newUser); if err!= nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ep.store.User().Create(newUser.CreateUserDtoToModel())

	if err != nil{
		g.JSON(http.StatusNoContent, err)
		return
	}

	g.JSON(http.StatusCreated, user)
}

// PingExample godoc
// @Summary delete user
// @Schemes
// @Description delete user
// @Tags user
// @Accept json
// @Produce json
// @Param id path integer true "User id"
// @Success 200 {string} Helloworld
// @Router /user/{id} [delete]
func (ep *Endpoints)DeleteUser(g *gin.Context){

	id, err := strconv.Atoi(g.Param("id"))

	if err != nil{
		g.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}

	if err := ep.store.User().DeleteById(id); err != nil{
		g.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusNoContent,nil)
}

// PingExample godoc
// @Summary get user by id
// @Schemes
// @Description get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path integer false "User id"
// @Success 200 {string} Helloworld
// @Router /user/{id} [get]
func (ep *Endpoints)GetUser(g *gin.Context){
	var user *model.User

	stringId := g.Param("id")

	id, err := strconv.Atoi(stringId)

	if err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err = ep.store.User().GetUserByID(id)
	
	if err!= nil{
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK,user)
}

// PingExample godoc
// @Summary get all users
// @Schemes
// @Description get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user [get]
func (ep *Endpoints)GetAllUsers(g *gin.Context){
	var users *[]model.User

	users, err := ep.store.User().GetUsers()
	
	if err!= nil{
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK,users)
}

// PingExample godoc
// @Summary modify user by id
// @Schemes
// @Description modify user by id
// @Tags user
// @Accept json
// @Produce json
// @Param user body dtos.ModifyUserDto true "User"
// @Success 200 {string} Helloworld
// @Router /user [put]
func (ep *Endpoints)ModifyUser(g *gin.Context){
	var modUser dtos.ModifyUserDto

	if err := g.BindJSON(&modUser); err!= nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ep.store.User().ModifyUser(modUser.ModifyUserDtoToModel()); err != nil{
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusNoContent, nil)
}



