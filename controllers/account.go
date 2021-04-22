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

type AccountController struct{}

func (a AccountController) ActionUserall(c *gin.Context) {

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	//rows, err := db.Query("SELECT * FROM wt_account")

	// tableData := make([]map[string]interface{}, 0)

	// for rows.Next() {
	// 	entry := make(map[string]interface{})

	// 	//err := rows.MapScan(entry)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	tableData = append(tableData, entry)
	// }

	account := []models.Account{}
	err = db.Select(&account, "SELECT * FROM wt_account")

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "fun ActionUserall", "data": account})

}

func (a AccountController) ActionView(c *gin.Context) {
	id := c.Param("id")

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	account := models.Account{}
	cmd := fmt.Sprintf("SELECT * FROM wt_account WHERE id=%s", id)
	err = db.Get(&account, cmd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "fun ActionView", "data": account})
}

func (a AccountController) ActionCreate(c *gin.Context) {

	mapbody := make(map[string]string)
	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(jsonData, &mapbody)

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	var countRows int
	cmd := fmt.Sprintf("SELECT count(*) FROM wt_account WHERE username='%s'", mapbody["username"])
	_ = db.Get(&countRows, cmd)

	if countRows > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "มีข้อมูล username นี้อยู่แล้ว !"})
		return
	}

	_, err = db.NamedExec(`INSERT INTO wt_account (username,password,name) VALUES (:username,:password,:name)`,
		map[string]interface{}{
			"username": mapbody["username"],
			"password": mapbody["password"],
			"name":     mapbody["name"],
		})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create Data Success !", "data": nil})
}

func (a AccountController) ActionUpdate(c *gin.Context) {

	mapbody := make(map[string]string)
	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(jsonData, &mapbody)

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	account := models.Account{}
	err = db.Get(&account, "SELECT * FROM wt_account WHERE id= ? ", mapbody["id"])

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	_, err = db.Exec(`UPDATE wt_account SET username = ? , password = ? , name = ? WHERE id = ?`,
		ShortcutsIF(mapbody["username"], account.Username),
		ShortcutsIF(mapbody["password"], account.Passowrd),
		ShortcutsIF(mapbody["name"], account.Name),
		mapbody["id"],
	)

	c.JSON(http.StatusOK, gin.H{"message": "Update Data Success !", "data": mapbody})
}

func (a AccountController) ActionDelete(c *gin.Context) {

	id := c.Param("id")

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`DELETE FROM wt_account WHERE id= ?`, id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Data Success !"})
}

func ShortcutsIF(textDef string, textNot string) string {
	if textDef == "" {
		return textNot
	} else {
		return textDef
	}
}
