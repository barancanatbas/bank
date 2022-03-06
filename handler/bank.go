package handler

import (
	"log"
	"mongoexample/models"
	"mongoexample/repository"
	"mongoexample/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HandlerBank interface {
	Create(c *gin.Context)
	Info(c *gin.Context)
	AddMoney(c *gin.Context)
	ReduceMoney(c *gin.Context)
}

// compile time prof
var _ HandlerBank = handlerBank{}

type handlerBank struct {
	Repo repository.RepoBank
}

func Bank(repo repository.RepoBank) handlerBank {
	return handlerBank{Repo: repo}
}

func (h handlerBank) Create(c *gin.Context) {

	var req models.Account

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// account := models.Account{
	// 	Name:    "deneme",
	// 	Balance: 10,
	// 	Status:  true,
	// }

	err := h.Repo.Create(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (h handlerBank) Info(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Repo.Info(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
}

func (h handlerBank) AddMoney(c *gin.Context) {
	var req request.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := primitive.ObjectIDFromHex(req.ID)
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"balance": req.Balance}}

	err := h.Repo.AddMoney(filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h handlerBank) ReduceMoney(c *gin.Context) {
	var req request.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := primitive.ObjectIDFromHex(req.ID)
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"balance": -req.Balance}}

	err := h.Repo.AddMoney(filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
