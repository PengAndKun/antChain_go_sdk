package request

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type Account struct {
	Account string `json:"account"`
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
}
