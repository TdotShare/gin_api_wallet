package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"root/config"
	"root/models"

	"github.com/gin-gonic/gin"
)

type HistoryController struct{}

func (h HistoryController) ActionHistoryGet(c *gin.Context) {
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	history := []models.History{}
	err = db.Select(&history, "SELECT * FROM wt_history")

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "fun ActionHistoryGet", "data": history})
}

func (h HistoryController) ActionCreate(c *gin.Context) {

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	mapbody := make(map[string]string)
	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(jsonData, &mapbody)

	bank := models.Bank{}
	err = db.Get(&bank, "SELECT * FROM wt_bank WHERE id = ?", mapbody["bank_id"])

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "log": "ก่อนเข้า INSERT"})
		return
	}

	_, err = db.NamedExec(`INSERT INTO wt_history (bank_id , price , status) VALUES (:bank_id , :price , :status)`,
		map[string]interface{}{
			"bank_id": mapbody["bank_id"],
			"price":   mapbody["price"],
			"status":  mapbody["status"],
		})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create Data Success !", "data": mapbody})

}

func (h HistoryController) ActionDelete(c *gin.Context) {

	id := c.Param("id")

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`DELETE FROM wt_history WHERE id= ?`, id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Data Success !"})

}
