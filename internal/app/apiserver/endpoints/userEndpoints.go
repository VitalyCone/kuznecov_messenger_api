package endpoints

import (
	"net/http"
	"strconv"

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
// @Param username body model.User true "Account username"
// @Success 200 {string} Helloworld
// @Router /user [post]
func (ep *Endpoints)CreateUser(g *gin.Context){
	var newUser model.User

	err := g.BindJSON(&newUser); if err!= nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ep.store.User().Create(&newUser)

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
// @Param id query integer true "User id"
// @Success 200 {string} Helloworld
// @Router /user [delete]
func (ep *Endpoints)DeleteUser(g *gin.Context){

	//g.JSON(http.StatusOK,g.Param("id"))
	id, err := strconv.Atoi(g.Query("id"))

	if err != nil{
		g.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}

	if err := ep.store.User().DeleteById(id); err != nil{
		g.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusNoContent,nil)
}

// PingExample godoc
// @Summary get user by id / get all users
// @Schemes
// @Description get user by id or get all users if id == "" (empty)
// @Tags user
// @Accept json
// @Produce json
// @Param id query integer false "User id"
// @Success 200 {string} Helloworld
// @Router /user [get]
func (ep *Endpoints)GetUsers(g *gin.Context){
	var user *model.User

	stringId := g.Query("id")

	if stringId == ""{
		ep.getAllUsers(g)
		return
	}

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

func (ep *Endpoints)getAllUsers(g *gin.Context){
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
// @Param username body model.User true "User"
// @Success 200 {string} Helloworld
// @Router /user [put]
func (ep *Endpoints)ModifyUser(g *gin.Context){
	var modUser model.User

	if err := g.BindJSON(&modUser); err!= nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ep.store.User().ModifyUser(&modUser); err != nil{
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusNoContent, nil)
}



