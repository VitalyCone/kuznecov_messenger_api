package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary create chat
// @Schemes
// @Description create chat
// @Tags chat
// @Accept json
// @Produce json
// @Param chat body model.Chat true "new chat data"
// @Success 200 {string} Helloworld
// @Router /chat [post]
func (ep *Endpoints) CreateChat(g *gin.Context) {
	var newChat model.Chat

	err := g.BindJSON(&newChat)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat, err := ep.store.Chat().Create(&newChat)

	if err != nil {
		g.JSON(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusCreated, chat)
}

// PingExample godoc
// @Summary get all chats
// @Schemes
// @Description get all chats
// @Tags chat
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /chat/all [get]
func (ep *Endpoints) GetAllChats(g *gin.Context) {
	chats, err := ep.store.Chat().GetAll()

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, chats)
}

// PingExample godoc
// @Summary find chat by id/user1_id/user2_id
// @Schemes
// @Description create chat
// @Tags chat
// @Accept json
// @Produce json
// @Param id query int false "id"
// @Param user1_id query int false "id"
// @Param user2_id query int false "id"
// @Success 200 {string} Helloworld
// @Router /chat [GET]
func (ep *Endpoints) GetChats(g *gin.Context) {
	valArray := []string{g.Query("user1_id"), g.Query("user2_id"), g.Query("id")}

	if valArray[0] != "" {
		num, err := strconv.Atoi(valArray[0])
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		chats, err := ep.store.Chat().GetByChatsUser1Id(num)

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		g.JSON(http.StatusOK, chats)

	} else if valArray[1] != "" {
		num, err := strconv.Atoi(valArray[1])
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		chats, err := ep.store.Chat().GetByChatsUser2Id(num)

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		g.JSON(http.StatusOK, chats)

	} else if valArray[2] != "" {
		num, err := strconv.Atoi(valArray[2])
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		chats, err := ep.store.Chat().GetByChatById(num)

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		g.JSON(http.StatusOK, chats)
	} else {
		g.JSON(http.StatusNotFound, nil)
	}
}

// PingExample godoc
// @Summary delete chat
// @Schemes
// @Description delete chat
// @Tags chat
// @Accept json
// @Produce json
// @Param id query integer true "chat id"
// @Success 200 {string} Helloworld
// @Router /chat [delete]
func (ep *Endpoints) DeleteChat(g *gin.Context) {

	id, err := strconv.Atoi(g.Query("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := ep.store.Chat().DeleteById(id); err != nil {
		g.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusNoContent, nil)
}
