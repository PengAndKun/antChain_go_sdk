package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type ContractERC20Mint struct {
	//ContractName string `json:"contractName"`
	Account string `json:"account"`
	Amount  string `json:"amount"`
}
type ContractERC20Admin struct {
	Account string `json:"account"`
}
type ContractERC20Burn struct {
	//ContractName string `json:"contractName"`
	Account string `json:"account"`
	Amount  string `json:"amount"`
}

type ContractERC20BalanceOf struct {
	Account string `json:"account"`
}
type ContractERC20Transfer struct {
	Sender  string `json:"sender"`
	KmsID   string `json:"kmsID"`
	Account string `json:"account"`
	Amount  string `json:"amount"`
}
type ContractERC20TransferFrom struct {
	Sender      string `json:"sender"`
	AccountFrom string `json:"accountFrom"`
	AccountTo   string `json:"accountTo"`
	Amount      string `json:"amount"`
}

type ContractERC20Approve struct {
	Sender  string `json:"sender"`
	KmsID   string `json:"kmsID"`
	Account string `json:"account"`
	Amount  string `json:"amount"`
}
type ContractERC20DecreaseAllowance struct {
	Sender          string `json:"sender"`
	Spender         string `json:"spender"`
	SubtractedValue string `json:"subtractedValue"`
}

type ContractERC20IncreaseAllowance struct {
	Sender     string `json:"sender"`
	Spender    string `json:"spender"`
	AddedValue string `json:"addedValue"`
}

type ContractERC20Allowance struct {
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}

func CallContractERC20(r *gin.Engine) {
	//ERC20积分合约
	//生成积分
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
		var gas int64 = 180000
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
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//销毁积分
	r.POST("/callContract/ERC20/burn", func(c *gin.Context) {
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
			ERC20, "burn(identity,uint256,uint256)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//增加管理员权限
	r.POST("/callContract/ERC20/addAdmin", func(c *gin.Context) {
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
		adminData := ContractERC20Admin{}
		jsoniter.Unmarshal(buf, &adminData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(adminData.Account))
		inputParamListBytes, err := jsoniter.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "addAdmin(identity)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//查看当前账户是否是管理员
	r.POST("/callContract/ERC20/isAdmin", func(c *gin.Context) {
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
		adminData := ContractERC20Admin{}
		jsoniter.Unmarshal(buf, &adminData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(adminData.Account))
		inputParamListBytes, err := jsoniter.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "isAdmin(identity)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//总积分数量
	r.POST("/callContract/ERC20/totalSupply", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		var gas int64 = 50000
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		log.Println(orderId)
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "totalSupply()",
			"[]", `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		log.Println(baseResp)
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//积分合约名称
	r.POST("/callContract/ERC20/name", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		var gas int64 = 50000
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		log.Println(orderId)
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "name()",
			"[]", `["string"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		log.Println(baseResp)
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//积分合约标识符
	r.POST("/callContract/ERC20/symbol", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		var gas int64 = 50000
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		log.Println(orderId)
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "symbol()",
			"[]", `["string"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		log.Println(baseResp)
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//积分合约单位
	r.POST("/callContract/ERC20/decimals", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		var gas int64 = 50000
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		log.Println(orderId)
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "decimals()",
			"[]", `["uint8"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		log.Println(baseResp)
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//获取某一账户余额
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
	//获取某一账户实际可用余额
	r.POST("/callContract/ERC20/balanceOfAvailable", func(c *gin.Context) {
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
			ERC20, "balanceOfAvailable(identity)",
			string(inputParamListBytes), `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//积分转移（from to _to）
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
		jsonArr = append(jsonArr, InputParamIdentity(TransferFromData.AccountTo))
		jsonArr = append(jsonArr, TransferFromData.Amount)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())

		sender := TransferFromData.Sender
		if sender == "" {
			sender = RestBizAccount
		}

		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			ERC20, "transferFrom(identity,identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//sender积分转移（sender to account）
	r.POST("/callContract/ERC20/transfer", func(c *gin.Context) {
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
		transferData := ContractERC20Transfer{}
		jsoniter.Unmarshal(buf, &transferData)
		var gas int64 = 100000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(transferData.Account)) //transfer( identity account,uint256 amount)
		jsonArr = append(jsonArr, transferData.Amount)                      //transfer( identity account,uint256 amount)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		sender := transferData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		kmsId := transferData.KmsID
		if kmsId == "" {
			kmsId = RestBizKmsID
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			ERC20, "transfer(identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			kmsId, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//sender积分授权（sender approve account）
	r.POST("/callContract/ERC20/approve", func(c *gin.Context) {
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
		ApproveData := ContractERC20Approve{}
		jsoniter.Unmarshal(buf, &ApproveData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(ApproveData.Account)) //Approve( identity account,uint256 amount)
		jsonArr = append(jsonArr, ApproveData.Amount)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		sender := ApproveData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		kmsId := ApproveData.KmsID
		if kmsId == "" {
			kmsId = RestBizKmsID
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			ERC20, "approve(identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			kmsId, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//sender 减少积分授权（sender decreaseAllowance account ）
	r.POST("/callContract/ERC20/decreaseAllowance", func(c *gin.Context) {
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
		DecreaseAllowance := ContractERC20DecreaseAllowance{}
		jsoniter.Unmarshal(buf, &DecreaseAllowance)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(DecreaseAllowance.Spender)) //DecreaseAllowance( identity spender,uint256 subtractedValue)
		jsonArr = append(jsonArr, DecreaseAllowance.SubtractedValue)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		sender := DecreaseAllowance.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			ERC20, "decreaseAllowance(identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//sender 增加积分授权（sender increaseAllowance account ）
	r.POST("/callContract/ERC20/increaseAllowance", func(c *gin.Context) {
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
		IncreaseAllowance := ContractERC20IncreaseAllowance{}
		jsoniter.Unmarshal(buf, &IncreaseAllowance)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(IncreaseAllowance.Spender)) //IncreaseAllowance( identity spender,uint256 addedValue)
		jsonArr = append(jsonArr, IncreaseAllowance.AddedValue)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}

		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		sender := IncreaseAllowance.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			ERC20, "increaseAllowance(identity,uint256)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//查看积分授权（sender allowance account）
	r.POST("/callContract/ERC20/allowance", func(c *gin.Context) {
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
		AllowanceData := ContractERC20Allowance{}
		jsoniter.Unmarshal(buf, &AllowanceData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(AllowanceData.Owner)) //Allowance( identity owner,identity sender)
		jsonArr = append(jsonArr, InputParamIdentity(AllowanceData.Spender))

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())

		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			ERC20, "allowance(identity,identity)",
			string(inputParamListBytes), `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

}
