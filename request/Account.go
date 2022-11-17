package request

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type Account struct {
	Account string `json:"account"`
}
type CreateAccount struct {
	CreateAccountName string `json:"createAccountName"`
}
type QueryTransaction struct {
	Hash string `json:"hash"`
}

func Query(r *gin.Engine) {
	r.POST("/queryAccount", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		jsonType := Account{}
		jsoniter.Unmarshal(buf, &jsonType)
		//orderId := fmt.Sprintf("order_%v", u.String())
		bizid := RestBizBizID
		account := jsonType.Account
		baseResp, err := RestClient.QueryAccount(bizid, account)
		if !(err == nil && baseResp.Code == "200") {
			//((fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err)).Error())
			//log.Panicf(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err).Error())
			log.Panicln(baseResp)
		}
		result := baseResp
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.POST("/createAccount", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		jsonType := CreateAccount{}
		jsoniter.Unmarshal(buf, &jsonType)
		u := uuid.New()

		orderId := fmt.Sprintf("order_%v", u.String())
		bizid := RestBizBizID
		accountNew := jsonType.CreateAccountName
		baseResp, err := RestClient.CreateAccountWithKmsId(bizid, orderId, accountNew, RestBizTenantID, RestBizKmsID)
		if !(err == nil && baseResp.Code == "200") {
			log.Panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//QueryTransaction
	r.POST("/queryTransaction", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		jsonType := QueryTransaction{}
		jsoniter.Unmarshal(buf, &jsonType)

		hash := jsonType.Hash

		baseResp, err := RestClient.QueryTransaction(RestBizBizID, hash)

		if !(err == nil && baseResp.Code == "200") {
			log.Panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//QueryReceipt
	r.POST("/queryReceipt", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		jsonType := QueryTransaction{}
		jsoniter.Unmarshal(buf, &jsonType)

		hash := jsonType.Hash

		baseResp, err := RestClient.QueryReceipt(RestBizBizID, hash)

		if !(err == nil && baseResp.Code == "200") {
			log.Panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

}
