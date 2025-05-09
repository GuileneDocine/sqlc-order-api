package api

import (
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

	r.POST("/customer", h.handleCreateCustomer)
	r.POST("/product", h.handleCreateProduct)
	r.POST("/order", h.handleCreateOrder)

	r.GET("/customer/:id", h.handleGetCustomerById)
	r.GET("/order/:id", h.handleGetOrderById)
	r.GET("/product/:id", h.handleGETProductById)
	r.GET("/products/", h.handleGetAllProducts)

	r.PATCH("/product/:id", h.handleUpdateByIdProduct)
	r.DELETE("/customer/:id", h.handleDeleteCustomer)
	return r
}

func (h *MessageHandler) handleCreateCustomer(c *gin.Context) {
	var req repo.CreateCustomerParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.querier.CreateCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *MessageHandler) handleCreateProduct(c *gin.Context) {
	var req repo.CreateProductParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.querier.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *MessageHandler) handleCreateOrder(c *gin.Context) {
	var req repo.CreateOrderParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.querier.CreateOrder(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *MessageHandler) handleGetCustomerById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id is required"})
		return
	}

	customer, err := h.querier.GetCustomerByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *MessageHandler) handleGetOrderById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id is required"})
		return
	}

	order, err := h.querier.GetOrderById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *MessageHandler) handleGETProductById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id is required"})
		return
	}

	product, err := h.querier.GetProductById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *MessageHandler) handleUpdateByIdProduct(c *gin.Context) {
	var req repo.UpdateByIdProductParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.querier.UpdateByIdProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *MessageHandler) handleGetAllProducts(c *gin.Context) {

	product, err := h.querier.GetAllProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *MessageHandler) handleDeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id is required"})
		return
	}

	err := h.querier.DeleteCustomer(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Customer deleted successfully"})
}
