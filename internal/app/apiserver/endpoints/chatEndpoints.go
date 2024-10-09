package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary create chat
// @Schemes
// @Description create chat
// @Tags chat
// @Accept json
// @Produce json
// @Param chat body dtos.CreateChatDto true "new chat data"
// @Success 200 {string} Helloworld
// @Router /chat [post]
func (ep *Endpoints) CreateChat(g *gin.Context) {
	var newChat dtos.CreateChatDto

	err := g.BindJSON(&newChat)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modelChat,err := newChat.CreateChatDtoToModel(ep.store)

	if err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat, err := ep.store.Chat().Create(modelChat)

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
// @Router /chats/all [get]
func (ep *Endpoints) GetAllChats(g *gin.Context) {
	chats, err := ep.store.Chat().GetAll()

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, chats)
}

// PingExample godoc
// @Summary Get chat by id
// @Schemes
// @Description Get chat by 1
// @Tags chat
// @Accept json
// @Produce json
// @Param id path int false "id"
// @Success 200 {string} Helloworld
// @Router /chat/{id} [GET]
func (ep *Endpoints) GetChatById(g *gin.Context){
	id := g.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	chats, err := ep.store.ChatMessage().GetByChatId(num)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	g.JSON(http.StatusOK, chats)
}

// PingExample godoc
// @Summary Find chats for user
// @Schemes
// @Description Get chats for user
// @Tags chat
// @Accept json
// @Produce json
// @Param user_id path int false "user_id"
// @Success 200 {string} Helloworld
// @Router /chats/{user_id} [GET]
func (ep *Endpoints) GetChatsForUser(g *gin.Context){
	id := g.Param("user_id")

	num, err := strconv.Atoi(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	chats, err := ep.store.Chat().GetByChatsUser1Id(num)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	g.JSON(http.StatusOK, chats)
}

// PingExample godoc
// @Summary delete chat
// @Schemes
// @Description delete chat
// @Tags chat
// @Accept json
// @Produce json
// @Param id path integer true "chat id"
// @Success 200 {string} Helloworld
// @Router /chat/{id} [delete]
func (ep *Endpoints) DeleteChat(g *gin.Context) {

	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := ep.store.Chat().DeleteById(id); err != nil {
		g.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusNoContent, nil)
}
