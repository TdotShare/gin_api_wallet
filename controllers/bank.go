package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/config"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct{}

// ActionBankGet godoc
// @tags Bank
// @summary ข้อมูลแบงค์ทั้งหมด (กระเป๋าเงิน)
// @description bank all data
// @id ActionBankGet
// @produce json
// @response 200 {object} models.Bank "Success"
// @router /bank/all [get]
func (b BankController) ActionBankGet(c *gin.Context) {
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	bank := []models.Bank{}
	err = db.Select(&bank, "SELECT * FROM wt_bank")

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "fun ActionBankGet", "data": bank})
}

// ActionView godoc
// @tags Bank
// @summary ดูรายละเอียดข้อมูลกระเป๋าเงิน
// @description bank view id
// @id ActionViewBank
// @param id path int true "id bank"
// @produce json
// @response 200 {object} models.BankAccount "Success"
// @router /bank/view/{id} [get]
func (b BankController) ActionView(c *gin.Context) {

	id := c.Param("id")

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	bank := models.Bank{}
	cmd := fmt.Sprintf("SELECT * FROM wt_bank WHERE id=%s", id)
	err = db.Get(&bank, cmd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	bank_join_account := models.BankAccount{}

	cmd = fmt.Sprintf(
		`SELECT b.* , a.id 'account.id' , a.username 'account.username' , a.password 'account.password',
		a.name 'account.name' FROM wt_bank b , wt_account a  WHERE b.id = '%s' AND  b.account_id = a.id`, id)

	err = db.Get(&bank_join_account, cmd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	history_model := []models.History{}
	cmd = fmt.Sprintf("SELECT * FROM wt_history WHERE bank_id=%s", id)
	err = db.Select(&history_model, cmd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	bank_join_account.History = history_model

	c.JSON(http.StatusOK, gin.H{"message": "fun ActionView", "data": bank_join_account})

}

// ActionCreate godoc
// @tags Bank
// @summary สร้างกระเป๋าเงิน
// @description create bank
// @id ActionCreateBank
//@param Bank body models.Bank true "data bank"
// @produce json
// @router /bank/create [post]
func (b BankController) ActionCreate(c *gin.Context) {

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	mapbody := make(map[string]string)
	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(jsonData, &mapbody)

	var countRows int
	cmd := fmt.Sprintf("SELECT count(*) FROM wt_bank WHERE account_id='%s'", mapbody["account_id"])
	_ = db.Get(&countRows, cmd)

	if countRows > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "มีข้อมูล account_id นี้อยู่แล้ว !"})
		return
	}

	_, err = db.NamedExec(`INSERT INTO wt_bank (name , account_id) VALUES (:name ,:account_id)`,
		map[string]interface{}{
			"name":       mapbody["name"],
			"account_id": mapbody["account_id"],
		})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create Data Success !", "data": nil})

}

// ActionDelete godoc
// @tags Bank
// @id ActionDeleteBank
// @summary ลบข้อมูลข้อมูลกระเป๋าเงิน
// @description bank delete id
// @param id path int true "id bank"
// @produce json
// @router /bank/delete/{id} [get]
func (b BankController) ActionDelete(c *gin.Context) {

	id := c.Param("id")

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`DELETE FROM wt_bank WHERE id= ?`, id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Data Success !"})

}
