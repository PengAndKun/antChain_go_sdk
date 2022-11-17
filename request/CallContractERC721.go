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

type ContractERC721BalanceOf struct {
	ERC721Identity string `json:"ERC721Identity"`
	Account        string `json:"account"`
}

type ContractERC721OwnerOf struct {
	ERC721Identity string `json:"ERC721Identity"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721TotalSupply struct {
	ERC721Identity string `json:"ERC721Identity"`
}
type ContractERC721TokenURI struct {
	ERC721Identity string `json:"ERC721Identity"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721SetBaseURI struct {
	ERC721Identity string `json:"ERC721Identity"`
	BaseURI_       string `json:"baseURI_"`
}
type ContractERC721TransferFrom struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	AccountFrom    string `json:"accountFrom"`
	AccountTo      string `json:"accountTo"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721SetApprovalForAll struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	Operator       string `json:"operator"`
	Approved       bool   `json:"_approved"`
}
type ContractERC721IsApprovedForAll struct {
	ERC721Identity string `json:"ERC721Identity"`
	Owner          string `json:"owner"`
	Operator       string `json:"operator"`
}
type ContractERC721Approve struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	To             string `json:"to"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721GetApproved struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721SafeMint struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	To             string `json:"to"`
	TokenId        string `json:"tokenId"`
}
type ContractERC721SafeMintBatch struct {
	ERC721Identity string `json:"ERC721Identity"`
	Sender         string `json:"sender"`
	To             string `json:"to"`
	Quantity       string `json:"quantity"`
}

func CallContractERC721(r *gin.Engine) {
	//ERC721藏品合约
	//指定账户藏品数量
	r.POST("/callContract/ERC721/balanceOf", func(c *gin.Context) {
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
		BalanceOfData := ContractERC721BalanceOf{}
		jsoniter.Unmarshal(buf, &BalanceOfData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(BalanceOfData.Account)) //balanceOf(identity)
		//inputParamListBytes, err := json.Marshal(&jsonArr)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := BalanceOfData.ERC721Identity
		if BalanceOfData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "balanceOf(identity)",
			string(inputParamListBytes), `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//查询指定藏品id的拥有者
	r.POST("/callContract/ERC721/ownerOf", func(c *gin.Context) {
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
		ownerOfData := ContractERC721OwnerOf{}
		jsoniter.Unmarshal(buf, &ownerOfData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, ownerOfData.TokenId) //ownerOfData(uint256)
		//inputParamListBytes, err := json.Marshal(&jsonArr)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := ownerOfData.ERC721Identity
		if ownerOfData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "ownerOf(uint256)",
			string(inputParamListBytes), `["identity"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			log.Println(err)
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//总合约藏品数量
	r.POST("/callContract/ERC721/totalSupply", func(c *gin.Context) {
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
		totalSupplyData := ContractERC721TotalSupply{}
		jsoniter.Unmarshal(buf, &totalSupplyData)

		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := totalSupplyData.ERC721Identity
		if totalSupplyData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "totalSupply()",
			"[]", `["uint256"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
		}
		log.Println(baseResp)
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//获得藏品http地址
	r.POST("/callContract/ERC721/tokenURI", func(c *gin.Context) {
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
		tokenURIData := ContractERC721TokenURI{}
		jsoniter.Unmarshal(buf, &tokenURIData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, tokenURIData.TokenId) //tokenURI(uint256)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := tokenURIData.ERC721Identity
		if tokenURIData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "tokenURI(uint256)",
			string(inputParamListBytes), `["string"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			log.Println(err)
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//设置藏品http地址(需要管理员权限)
	r.POST("/callContract/ERC721/setBaseURI", func(c *gin.Context) {
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
		setBaseURIData := ContractERC721SetBaseURI{}
		jsoniter.Unmarshal(buf, &setBaseURIData)

		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, setBaseURIData.BaseURI_) //setBaseURI(string)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := setBaseURIData.ERC721Identity
		if setBaseURIData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "setBaseURI(string)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			log.Println(err)
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//藏品转移（from to _to ）
	r.POST("/callContract/ERC721/transferFrom", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		buf := make([]byte, 1024)
		u := uuid.New()
		c.Request.Body.Read(buf)
		transferFromData := ContractERC721TransferFrom{}
		jsoniter.Unmarshal(buf, &transferFromData)
		var gas int64 = 150000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(transferFromData.AccountFrom)) //transferFrom( identity from,identity to,uint256 tokenId)
		jsonArr = append(jsonArr, InputParamIdentity(transferFromData.AccountTo))
		jsonArr = append(jsonArr, transferFromData.TokenId)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		sender := transferFromData.Sender
		if sender == "" {
			sender = RestBizAccount
		}

		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := transferFromData.ERC721Identity
		if transferFromData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "transferFrom(identity,identity,uint256)",
			string(inputParamListBytes), `[""]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//授权所有的nft给某一地址
	r.POST("/callContract/ERC721/setApprovalForAll", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		u := uuid.New()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		setApprovalForAllData := ContractERC721SetApprovalForAll{}
		jsoniter.Unmarshal(buf, &setApprovalForAllData)
		var gas int64 = 50000
		orderId := fmt.Sprintf("order_%v", u.String())
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(setApprovalForAllData.Operator)) //setApprovalForAll(identity operator, bool _approved)
		jsonArr = append(jsonArr, setApprovalForAllData.Approved)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}

		sender := setApprovalForAllData.Sender
		erc721Identity := setApprovalForAllData.ERC721Identity
		if setApprovalForAllData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "setApprovalForAll(identity,bool)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//查看授权信息
	r.POST("/callContract/ERC721/isApprovedForAll", func(c *gin.Context) {
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
		isApprovedForAllData := ContractERC721IsApprovedForAll{}
		jsoniter.Unmarshal(buf, &isApprovedForAllData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(isApprovedForAllData.Owner)) //isApprovedForAll(identity owner, identity operator)
		jsonArr = append(jsonArr, InputParamIdentity(isApprovedForAllData.Operator))

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		u := uuid.New()

		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := isApprovedForAllData.ERC721Identity
		if isApprovedForAllData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, RestBizAccount, RestBizTenantID,
			erc721Identity, "isApprovedForAll(identity,identity)",
			string(inputParamListBytes), `["bool"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//单独授权某一nft权利
	r.POST("/callContract/ERC721/approve", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		u := uuid.New()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		ApproveData := ContractERC721Approve{}
		jsoniter.Unmarshal(buf, &ApproveData)
		//fmt.Println(ApproveData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(ApproveData.To)) //Approve(identity to,uint256 tokenId)
		jsonArr = append(jsonArr, ApproveData.TokenId)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))

		orderId := fmt.Sprintf("order_%v", u.String())
		erc721Identity := ApproveData.ERC721Identity
		if ApproveData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		sender := ApproveData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "approve(identity,uint256)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//获得某一NFT的权限信息(只能有一个权限)
	r.POST("/callContract/ERC721/getApproved", func(c *gin.Context) {
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
		getApprovedData := ContractERC721GetApproved{}
		jsoniter.Unmarshal(buf, &getApprovedData)
		var gas int64 = 50000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, getApprovedData.TokenId) //getApproved(uint256 tokenId)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))
		u := uuid.New()
		orderId := fmt.Sprintf("order_%v", u.String())
		sender := getApprovedData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		erc721Identity := getApprovedData.ERC721Identity
		if getApprovedData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "getApproved(uint256)",
			string(inputParamListBytes), `["identity"]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

	//生成指定藏品id的藏品(需要有管理员)
	r.POST("/callContract/ERC721/safeMint", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		u := uuid.New()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		safeMintData := ContractERC721SafeMint{}
		jsoniter.Unmarshal(buf, &safeMintData)
		//fmt.Println(ApproveData)
		var gas int64 = 500000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(safeMintData.To)) //safeMint(identity to,uint256 tokenId)
		jsonArr = append(jsonArr, safeMintData.TokenId)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))

		orderId := fmt.Sprintf("order_%v", u.String())
		sender := safeMintData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		erc721Identity := safeMintData.ERC721Identity
		if safeMintData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "safeMint(identity,uint256)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})
	//批量生成指定数量的藏品（需要有管理员）
	r.POST("/callContract/ERC721/safeMintBatch", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("catch exception")
				c.JSON(http.StatusOK, gin.H{
					"result": err,
				})
			}
		}()
		u := uuid.New()
		buf := make([]byte, 1024)
		c.Request.Body.Read(buf)
		safeMintBatchData := ContractERC721SafeMintBatch{}
		jsoniter.Unmarshal(buf, &safeMintBatchData)
		//fmt.Println(ApproveData)
		var gas int64 = 200000
		jsonArr := make([]interface{}, 0)
		jsonArr = append(jsonArr, InputParamIdentity(safeMintBatchData.To)) //safeMintBatch(identity to,uint256 quantity)
		jsonArr = append(jsonArr, safeMintBatchData.Quantity)

		inputParamListBytes, err := json.Marshal(&jsonArr)
		if err != nil {
			log.Panicln(err)
		}
		//fmt.Println(string(inputParamListBytes))

		orderId := fmt.Sprintf("order_%v", u.String())
		sender := safeMintBatchData.Sender
		if sender == "" {
			sender = RestBizAccount
		}
		erc721Identity := safeMintBatchData.ERC721Identity
		if safeMintBatchData.ERC721Identity == "" {
			erc721Identity = ERC721
		}
		baseResp, err := RestClient.CallContract(RestBizBizID,
			orderId, sender, RestBizTenantID,
			erc721Identity, "safeMintBatch(identity,uint256)",
			string(inputParamListBytes), `[]`,
			RestBizKmsID, false, gas)
		if !(err == nil && baseResp.Success && baseResp.Code == "200") {
			panic(baseResp)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": baseResp,
		})
	})

}
