package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/holiman/uint256"
	jsoniter "github.com/json-iterator/go"
)

type ContractERC20Mint struct {
	//ContractName string `json:"contractName"`
	Account string      `json:"contractName"`
	Amount  uint256.Int `json:"amount"`
}

type ContractERC20BalanceOf struct {
	Account string `json:"account"`
}
type ContractERC20Transfer struct {
	Account string      `json:"account"`
	Amount  uint256.Int `json:"amount"`
}
type ContractERC20TransferFrom struct {
	AccountFrom string      `json:"accountFrom"`
	AccountTo   string      `json:"accountTo"`
	Amount      uint256.Int `json:"amount"`
}

func CallContractERC20(r *gin.Engine) {
	//ERC20积分合约
	r.POST("/callContract/ERC20/mint", func(c *gin.Context) {
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
		mintData := ContractERC20Mint{}
		jsoniter.Unmarshal(buf, &mintData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(mintData.Account))
		jsonArr = append(jsonArr, mintData.Amount)
		inputParamListBytes, err := jsoniter.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "mint(identity,uint256)",
			string(inputParamListBytes), "(bool)",
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	r.POST("/callContract/ERC20/balanceOf", func(c *gin.Context) {
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
		BalanceOfData := ContractERC20BalanceOf{}
		jsoniter.Unmarshal(buf, &BalanceOfData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(BalanceOfData.Account)) //balanceOf(identity)
		//inputParamListBytes, err := json.Marshal(&jsonArr)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "balanceOf(identity)",
			string(inputParamListBytes), `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	r.POST("/callContract/ERC20/transferFrom", func(c *gin.Context) {
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
		TransferFromData := ContractERC20TransferFrom{}
		jsoniter.Unmarshal(buf, &TransferFromData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(TransferFromData.AccountFrom)) //transferFrom( identity from,identity to,uint256 amount)
		jsonArr = append(jsonArr, InputParamIdentity(TransferFromData.AccountTo))   //transferFrom( identity from,identity to,uint256 amount)
		jsonArr = append(jsonArr, TransferFromData.Amount)                          //transferFrom( identity from,identity to,uint256 amount)
		//inputParamListBytes, err := json.Marshal(&jsonArr)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "transferFrom(identity,identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
}
