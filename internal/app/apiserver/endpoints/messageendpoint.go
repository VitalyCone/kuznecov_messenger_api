package endpoints

import (
	"net/http"
	"strconv"

	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apiserver/dtos"
	"github.com/gin-gonic/gin"
)

// PingExample message
// @Summary Get message by id
// @Schemes
// @Description Get message by id
// @Tags message
// @Accept json
// @Produce json
// @Param id path int false "id"
// @Router /message/{id} [GET]
func (ep *Endpoints) GetMessage(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
	}

	message, err := ep.store.ChatMessage().Get(id)

	if err != nil {
		g.JSON(http.StatusNotFound, nil)
	}

	g.JSON(http.StatusOK, message)
}

// PingExample message
// @Summary Create message by id
// @Schemes
// @Description Create message by id
// @Tags message
// @Accept json
// @Produce json
// @Param message body dtos.CreateChatMessageDto false "message"
// @Router /message [POST]
func (ep *Endpoints) CreateMessage(g *gin.Context) {
	var newMessage dtos.CreateChatMessageDto

	err := g.BindJSON(&newMessage)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	modelMessage, err := newMessage.CreateChatMessageDtoToModel(ep.store)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := ep.store.ChatMessage().Create(modelMessage)

	if err != nil {
		g.JSON(http.StatusNotFound, nil)
	}

	g.JSON(http.StatusOK, message)
}

// PingExample message
// @Summary Delete message by id
// @Schemes
// @Description Delete message by id
// @Tags message
// @Accept json
// @Produce json
// @Param id path int false "id"
// @Router /message/{id} [DELETE]
func (ep *Endpoints) DeleteMessage(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
	}

	if err := ep.store.ChatMessage().Delete(id); err != nil {
		g.JSON(http.StatusNotFound, nil)
	}

	g.JSON(http.StatusNoContent, nil)
}
