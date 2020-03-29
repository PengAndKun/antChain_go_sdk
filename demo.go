package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/mychain/mychain-sdk-go/common/codec/contract/abi"
	"os"
	"strings"
	"time"
)

const (
	RestBizTestBizID    = "9b0c522d873a4bf7ac80efdbe27c808c"
	RestBizTestAccount  = "rest_biz_test_account"
	RestBizTestKmsID    = "rest_test:rest_biz_test_account"
	RestBizTestTenantID = "rest_test"
)

const abiJsonStr = `[
  {
    "constant": true,
    "inputs": [
      {
        "name": "b",
        "type": "bytes"
      },
      {
        "name": "s",
        "type": "string"
      }
    ],
    "name": "SayHello",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      },
      {
        "name": "",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "beneficiary",
    "outputs": [
      {
        "name": "",
        "type": "identity"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "say",
    "outputs": [
      {
        "name": "",
        "type": "identity"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "_greeting",
        "type": "uint256"
      },
      {
        "name": "a",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`

func main() {
	TestDeposit()
	TestDepositAsyncWithTransaction()
	TestQueryTransaction()
	TestQueryReceipt()
	TestDeployContractAndCallContract()
	TestCreateAndQueryAccountWithKmsIdAndDeposit()
	TestCreateAndQueryAccountWithKmsIdAndDeployCallContract()
}

func TestDepositAsyncWithTransaction() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-demo/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	if restClient.RestToken == "" {
		panic(fmt.Errorf("rest token:%+v is empty", restClient.RestToken))
	}

	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	content := "我是中国人"
	var gas int64 = 50000
	baseResp, err := restClient.DepositAsyncWithTransaction(RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, content, RestBizTestKmsID, gas)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	jsonObject := make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	innerObject := jsonObject["transactionDO"].(map[string]interface{})
	encodedData := innerObject["data"].(string)
	bytes, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	if data != content {
		panic(fmt.Errorf("origin isn't the same with content,origin:%+v content:%+v", data, content))
	}
}

func TestDeposit() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-demo/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	if restClient.RestToken == "" {
		panic(fmt.Errorf("rest token:%+v is empty", restClient.RestToken))
	}

	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	content := "我是中国人"
	var gas int64 = 50000
	baseResp, err := restClient.Deposit(RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, content, RestBizTestKmsID, gas)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}

	tick := time.Tick(time.Duration(500) * time.Millisecond) // wait 2s for tx finished...

	hash := baseResp.Data
	for i := 0; i < 3; i++ {
		select {
		case <-tick:
			baseResp, err = restClient.QueryReceipt(RestBizTestBizID, hash)
			if !(err == nil && baseResp.Code == "200") {
				panic(fmt.Errorf("no succ receipt baseResp:%+v err:%+v", baseResp, err))
			}
			baseResp, err = restClient.QueryTransaction(RestBizTestBizID, hash)
			if !(err == nil && baseResp.Code == "200") {
				panic(fmt.Errorf("no succ transaction baseResp:%+v err:%+v", baseResp, err))
			}
		}
	}
	jsonObject := make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	innerObject := jsonObject["transactionDO"].(map[string]interface{})
	encodedData := innerObject["data"].(string)
	bytes, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	if data != content {
		panic(fmt.Errorf("origin isn't the same with content,origin:%+v content:%+v", data, content))
	}
}

func TestQueryTransaction() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	hash := "b457afacb11dff49020f70ea1a80059b2d98466a58399d36e5b71389827216b2"
	baseResp, err := restClient.QueryTransaction(RestBizTestBizID, hash)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	fmt.Printf("baseResp:%+v\n", baseResp)
}

func TestQueryReceipt() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	hash := "b457afacb11dff49020f70ea1a80059b2d98466a58399d36e5b71389827216b2"
	baseResp, err := restClient.QueryReceipt(RestBizTestBizID, hash)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	fmt.Printf("baseResp:%+v\n", baseResp)
}

/** 下面为测试合约的abi
[
  {
    "constant": true,
    "inputs": [
      {
        "name": "b",
        "type": "bytes"
      },
      {
        "name": "s",
        "type": "string"
      }
    ],
    "name": "SayHello",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      },
      {
        "name": "",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "beneficiary",
    "outputs": [
      {
        "name": "",
        "type": "identity"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "say",
    "outputs": [
      {
        "name": "",
        "type": "identity"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "_greeting",
        "type": "uint256"
      },
      {
        "name": "a",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]
*/
func TestDeployContractAndCallContract() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	u := uuid.New()
	contractName := fmt.Sprintf("test_biz_deploy_contract_%v", u.String())
	orderId := fmt.Sprintf("order_%v", u.String())
	//deploy contract
	baseResp, err := restClient.DeployContract(RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, RestBizTestKmsID, contractName, "608060405234801561001057600080fd5b506040516102ef3803806102ef833981018060405281019080805190602001909291908051820192919050505081600081905550600060018190555050506102928061005d6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631002aecd1461005c57806338af3eed146101f0578063954ab4b21461021b575b600080fd5b34801561006857600080fd5b50610109600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610246565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561014d578082015181840152602081019050610132565b50505050905090810190601f16801561017a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101b3578082015181840152602081019050610198565b50505050905090810190601f1680156101e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156101fc57600080fd5b50610205610256565b6040518082815260200191505060405180910390f35b34801561022757600080fd5b5061023061025c565b6040518082815260200191505060405180910390f35b6060808383915091509250929050565b60015481565b60006001549050905600a165627a7a72305820ac9ff0ce4f83f475e39f7a8ecdfeb0b16673a328ca1af858b2ce81ccbe75837c0029")
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
	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	baseResp, err = restClient.CallContract(RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, contractName, "SayHello(bytes,string)", string(inputParamListBytes), `["bytes","string"]`, RestBizTestKmsID, false)
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
	if arg2 != output2 {
		panic(fmt.Errorf("input arg2:%s is not same with output2:%s", arg2, output2))
	}
	//local call
	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	baseResp, err = restClient.CallContract(RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, contractName, "SayHello(bytes,string)", string(inputParamListBytes), `["bytes","string"]`, RestBizTestKmsID, true)
	if !(err == nil && baseResp.Success && baseResp.Code == "0") {
		panic(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	}
	err = json.Unmarshal([]byte(baseResp.Data), &outputs)
	if err != nil {
		panic(err)
	}
	output1, err = base64.StdEncoding.DecodeString(outputs.OutRes[0].(string))
	if err != nil {
		panic(err)
	}
	output2 = outputs.OutRes[1].(string)
	if !isBytesSame(arg1, output1) {
		panic(fmt.Errorf("intput arg1:%+v is not same with output1:%+v", arg1, output1))
	}
	if arg2 != output2 {
		panic(fmt.Errorf("input arg2:%s is not same with output2:%s", arg2, output2))
	}
	//async call and query receipts
	abi, err := abi.JSON(strings.NewReader(abiJsonStr))
	if err != nil {
		panic(err)
	}
	sayHelloResp := &[]interface{}{&[]byte{}, new(string)}
	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	orderId = fmt.Sprintf("order_%v", u.String())
	inputParamListStr := string(inputParamListBytes)
	resp, err := restClient.CallSolcContractAsyncWithReceipt(abi, RestBizTestBizID, orderId, RestBizTestAccount, RestBizTestTenantID, RestBizTestKmsID, contractName, "SayHello(bytes,string)", inputParamListStr, `["bytes","string"]`, sayHelloResp)
	if !(err == nil && resp.Success && resp.Code == "200") {
		panic(fmt.Errorf("callSolcContractAsyncWithReceipt failed resp:%+v err:%+v", resp, err))
	}
	output1 = *(*sayHelloResp)[0].(*[]byte)
	output2 = *(*sayHelloResp)[1].(*string)
	if !isBytesSame(arg1, output1) {
		panic(fmt.Errorf("intput arg1:%+v is not same with output1:%+v", arg1, output1))
	}
	if arg2 != output2 {
		panic(fmt.Errorf("input arg2:%s is not same with output2:%s", arg2, output2))
	}
}

func TestCreateAndQueryAccountWithKmsIdAndDeposit() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(err)
	}
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	kmsId := fmt.Sprintf("%s_%s", RestBizTestTenantID, u.String())
	account := fmt.Sprintf("myaccount_%s", u.String())
	baseResp, err := restClient.CreateAccountWithKmsId(RestBizTestBizID, orderId, account, RestBizTestTenantID, kmsId)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("create account with kmsId failed,resp:%+v err:%+v", baseResp, err))
	}

	baseResp, err = restClient.QueryAccount(RestBizTestBizID, account)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("query account failed,resp:%+v err:%+v", baseResp, err))
	}

	jsonObject := make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	status := jsonObject["status"].(float64)
	if status != 0 {
		panic(fmt.Errorf("account status is wrong,status:%v", status))
	}
	fmt.Printf("%+v\n", jsonObject)

	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	content := "我是中国人"
	var gas int64 = 50000
	baseResp, err = restClient.Deposit(RestBizTestBizID, orderId, account, RestBizTestTenantID, content, kmsId, gas)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ deposit resp,resp:%+v err:%+v", baseResp, err))
	}
	hash := baseResp.Data
	time.Sleep(2 * time.Second) // wait for some time
	baseResp, err = restClient.QueryTransaction(RestBizTestBizID, hash)
	if !(err == nil && baseResp.Code == "200") {
		panic(fmt.Errorf("no succ transaction baseResp:%+v err:%+v", baseResp, err))
	}
	jsonObject = make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	innerObject := jsonObject["transactionDO"].(map[string]interface{})
	encodedData := innerObject["data"].(string)
	bytes, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	if data != content {
		panic(fmt.Errorf("origin isn't the same with content,origin:%+v content:%+v", data, content))
	}
}

func TestCreateAndQueryAccountWithKmsIdAndDeployCallContract() {
	configFilePath := os.Getenv("GOPATH") + "/src/gitlab.alipay-inc.com/antchain/restclient-go-sdk/test/rest-config.json"
	restClient, err := client.NewRestClient(configFilePath)
	if err != nil {
		panic(err)
	}
	u := uuid.New()
	orderId := fmt.Sprintf("order_%v", u.String())
	kmsId := fmt.Sprintf("%s_%s", RestBizTestTenantID, u.String())
	account := fmt.Sprintf("myaccount_%s", u.String())
	baseResp, err := restClient.CreateAccountWithKmsId(RestBizTestBizID, orderId, account, RestBizTestTenantID, kmsId)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("create account with kmsId failed,resp:%+v err:%+v", baseResp, err))
	}

	baseResp, err = restClient.QueryAccount(RestBizTestBizID, account)
	if !(err == nil && baseResp.Success && baseResp.Code == "200") {
		panic(fmt.Errorf("query account failed,resp:%+v err:%+v", baseResp, err))
	}

	jsonObject := make(map[string]interface{})
	err = json.Unmarshal([]byte(baseResp.Data), &jsonObject)
	if err != nil {
		panic(err)
	}
	status := jsonObject["status"].(float64)
	if status != 0 {
		panic(fmt.Errorf("account status is wrong,status:%v", status))
	}
	fmt.Printf("%+v\n", jsonObject)

	u = uuid.New()
	contractName := fmt.Sprintf("test_biz_deploy_contract_%v", u.String())

	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	//deploy contract
	baseResp, err = restClient.DeployContract(RestBizTestBizID, orderId, account, RestBizTestTenantID, kmsId, contractName, "608060405234801561001057600080fd5b506040516102ef3803806102ef833981018060405281019080805190602001909291908051820192919050505081600081905550600060018190555050506102928061005d6000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631002aecd1461005c57806338af3eed146101f0578063954ab4b21461021b575b600080fd5b34801561006857600080fd5b50610109600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610246565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561014d578082015181840152602081019050610132565b50505050905090810190601f16801561017a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156101b3578082015181840152602081019050610198565b50505050905090810190601f1680156101e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156101fc57600080fd5b50610205610256565b6040518082815260200191505060405180910390f35b34801561022757600080fd5b5061023061025c565b6040518082815260200191505060405180910390f35b6060808383915091509250929050565b60015481565b60006001549050905600a165627a7a72305820ac9ff0ce4f83f475e39f7a8ecdfeb0b16673a328ca1af858b2ce81ccbe75837c0029")
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
	u = uuid.New()
	orderId = fmt.Sprintf("order_%v", u.String())
	baseResp, err = restClient.CallContract(RestBizTestBizID, orderId, account, RestBizTestTenantID, contractName, "SayHello(bytes,string)", string(inputParamListBytes), `["bytes","string"]`, kmsId, false)
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
	if arg2 != output2 {
		panic(fmt.Errorf("input arg2:%s is not same with output2:%s", arg2, output2))
	}
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
