package api

import (
	"database/sql"
	"net/http"

	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	querier repo.Querier
}

func NewMessageHandler(querier repo.Querier) *MessageHandler {
	return &MessageHandler{
		querier: querier,
	}
}

func (h *MessageHandler) WireHttpHandler() http.Handler {

	r := gin.Default()
	r.Use(gin.CustomRecovery(func(c *gin.Context, _ any) {
		c.String(http.StatusInternalServerError, "Internal Server Error: panic")
		c.AbortWithStatus(http.StatusInternalServerError)
	})) //prevents the server from crashing if an error occurs in any route

	r.POST("/message", h.handleCreateMessage)
	r.GET("/message/:id", h.handleGetMessage)
	r.GET("/thread/:id/messages", h.handleGetThreadMessages)
	r.POST("/thread", h.handleCreateThread)
	r.DELETE("/message/:id", h.handleDeleteMessage)
	r.PATCH("/message", h.handleUpdateMessage)

	return r
}

type threadParam struct {
	Topic string `json:"topic"`
}

func (h *MessageHandler) handleCreateThread(c *gin.Context) {
	var req threadParam
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := h.querier.CreateThread(c, &req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

func (h *MessageHandler) handleCreateMessage(c *gin.Context) {
	var req repo.CreateMessageParams
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the thread exist
	_, err = h.querier.GetThreadById(c, *req.ThreadID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No row found for this thread id"})
			return
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.querier.CreateMessage(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetMessage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	message, err := h.querier.GetMessageByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) handleGetThreadMessages(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	messages, err := h.querier.GetMessagesByThread(c, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"thread":   id,
		"topic":    "example",
		"messages": messages,
	})
}

func (h *MessageHandler) handleDeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := h.querier.DeleteMessage(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete message"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "message deleted successfully"})

}

func (h *MessageHandler) handleUpdateMessage(c *gin.Context) {
	var req repo.UpdateMessageParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.querier.UpdateMessage(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Message updated successfully"})
}
