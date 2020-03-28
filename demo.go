package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/model"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/response"
	"os"
)

func main() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(err)
	}
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	kmsId := fmt.Sprintf("%s_%s", restClient.RestClientProperties.TenantId, u.String())
	account := fmt.Sprintf("myaccount_%s", u.String())
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			BizId:  restClient.RestClientProperties.BizId,
			Method: model.CREATEACCOUNT,
		},
		OrderId:    orderId,
		Account:    account,
		MykmsKeyId: kmsId,
	}
	baseResp, err := restClient.ChainCallForBiz(callRestBizParam)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("create account with kmsId failed,resp:%+v err:%+v", baseResp, err))
	}

	clientParam, err := restClient.CreateQueryAccountParam(account)
	if err != nil {
		panic(err)
	}

	baseResp, err = restClient.ChainCall(clientParam.Hash, restClient.RestClientProperties.BizId, clientParam.SignData, model.QUERYACCOUNT)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("query account failed,resp:%+v err:%+v", baseResp, err))
	}

	jsonObject := make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	status := int64(jsonObject["status"].(float64))
	if !(status == 0) {
		panic(fmt.Errorf("account status is wrong,status:%v", status))
	}
	fmt.Printf("%+v\n", jsonObject)

	u = uuid.New()
	contractName := fmt.Sprintf("test_biz_deploy_contract_%v", u.String())

	//deploy contract
	baseResp, err = deployContract(restClient, account, kmsId, contractName, "608060405234801561001057600080fd5b506040516102ef3803806102ef833981018060405281019080805190602001909291908051820192919050505081600081905550600060018190555050506102928061005d6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631002aecd1461005c57806338af3eed146101f0578063954ab4b21461021b575b600080fd5b34801561006857600080fd5b50610109600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610246565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561014d578082015181840152602081019050610132565b50505050905090810190601f16801561017a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101b3578082015181840152602081019050610198565b50505050905090810190601f1680156101e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156101fc57600080fd5b50610205610256565b6040518082815260200191505060405180910390f35b34801561022757600080fd5b5061023061025c565b6040518082815260200191505060405180910390f35b6060808383915091509250929050565b60015481565b60006001549050905600a165627a7a72305820ac9ff0ce4f83f475e39f7a8ecdfeb0b16673a328ca1af858b2ce81ccbe75837c0029")
	if !(err == nil && baseResp.Code == "0") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	//call contract
	arg1 := make([]byte, 13)
	for i := 0; i < 13; i++ {
		arg1[i] = byte(i)
	}
	arg2 := "hello"
	jsonArr := make([]interface{}, 0)
	jsonArr = append(jsonArr, arg1)
	jsonArr = append(jsonArr, arg2)
	inputParamListBytes, err := json.Marshal(&jsonArr)
	if err != nil {
		panic(err)
	}
	baseResp, err = callContract(restClient, account, contractName, "SayHello(bytes,string)", string(inputParamListBytes), `["bytes","string"]`, kmsId, false)
	if !(err == nil && baseResp.Success && baseResp.Code == "0") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	type Output struct {
		OutRes []interface{} `json:"outRes"`
	}
	outputs := Output{}
	err = json.Unmarshal([]byte(baseResp.Data), &outputs)
	if err != nil {
		panic(err)
	}
	output1, err := base64.StdEncoding.DecodeString(outputs.OutRes[0].(string))
	if err != nil {
		panic(err)
	}
	output2 := outputs.OutRes[1].(string)
	if !isBytesSame(arg1, output1) {
		panic(fmt.Errorf("intput arg1:%+v is not same with output1:%+v", arg1, output1))
	}
	if !(arg2 == output2) {
		panic(fmt.Errorf("input arg2:%s is not same with output2:%s", arg2, output2))
	}
}

func callContract(client *client.RestClient, account, contractName, methodSignature, inputParamListStr, outTypes, kmsId string, isLocal bool) (response.BaseResp, error) {
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			BizId:  client.RestClientProperties.BizId,
			Method: model.CALLCONTRACTBIZ,
		},
		OrderId:            orderId,
		Account:            account,
		ContractName:       contractName,
		MethodSignature:    methodSignature,
		InputParamListStr:  inputParamListStr,
		OutTypes:           outTypes,
		MykmsKeyId:         kmsId,
		IsLocalTransaction: isLocal,
	}
	return client.ChainCallForBiz(callRestBizParam)
}

func deployContract(client *client.RestClient, account, kmsId, contractName, contractCode string) (response.BaseResp, error) {
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	//deploy contract
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			BizId:  client.RestClientProperties.BizId,
			Method: model.DEPLOYCONTRACTFORBIZ,
		},
		OrderId:      orderId,
		Account:      account,
		MykmsKeyId:   kmsId,
		ContractName: contractName,
		ContractCode: contractCode,
	}
	return client.ChainCallForBiz(callRestBizParam)
}

func deposit(client *client.RestClient, account, content, mykmsKeyId string) (response.BaseResp, error) {
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			BizId:  client.RestClientProperties.BizId,
			Method: model.DEPOSIT,
		},
		OrderId:    orderId,
		Account:    account,
		Content:    content,
		MykmsKeyId: mykmsKeyId,
		Gas:        50000, // 可选
	}
	return client.ChainCallForBiz(callRestBizParam)
}

func queryReceipt(client *client.RestClient, hash string) (response.BaseResp, error) {
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			Hash:   hash,
			Method: model.QUERYRECEIPT,
		},
	}
	return client.ChainCallForBiz(callRestBizParam)
}

func queryTransaction(client *client.RestClient, hash string) (response.BaseResp, error) {
	callRestBizParam := model.CallRestBizParam{
		BaseParam: model.BaseParam{
			Hash:   hash,
			Method: model.QUERYTRANSACTION,
		},
	}
	return client.ChainCallForBiz(callRestBizParam)
}

func isBytesSame(a, b []byte) bool {
	if a == nil && b != nil || a != nil && b == nil || len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
