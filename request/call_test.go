package request

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"gitlab.alipay-inc.com/antchain/restclient-go-demo/logger"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
)

func init() {
	var err error
	//configFilePath := os.Getenv("GOPATH") + "src/gitlab.alipay-inc.com/antchain/restclient-go-demo/rest-config.json"
	configFilePath := "rest-config.json"
	RestClient, err = client.NewRestClient(configFilePath)
	if err != nil {
		logger.L().Debug(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	if RestClient.RestToken == "" {
		logger.L().Debug(fmt.Errorf("rest token:%+v is empty", RestClient.RestToken))
	}
}
func Test(t *testing.T) {
	accountIdentity := "0x7c262b618eb2a90d92ab903f4f48ea37a9af4c4ce5a4b4c039b28e1eeddb6689"
	_accountName := "qiuguowenhua"

	accountName := sha256.Sum256([]byte(_accountName))
	println(string(accountName[:]))
	a := hex.EncodeToString(accountName[:])
	b := base64.StdEncoding.EncodeToString(accountName[:])
	println(string(a))
	println("b is ", b)
	// var gas int64 = 50000
	jsonArr := make([]interface{}, 0)
	// a := string(InputParamIdentity(_accountName)[:])
	// fmt.Println(a)
	jsonArr = append(jsonArr, InputParamIdentityFromName(_accountName))

	inputParamListBytes, err := json.Marshal(&jsonArr)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("inputParamListBytes is ", string(inputParamListBytes[:]))
	// u := uuid.New()
	// orderId := fmt.Sprintf("order_%v", u.String())
	// fmt.Println(orderId)
	// in := string(inputParamListBytes[:])
	// fmt.Println("in is ", in)

	// fmt.Println(strings.Compare("0x", accountIdentity[:2]))
	//in2 := "[{\"data\":\"fCYrYY6yqQ2Sq5A/T0jqN6mvTEzlpLTAObKOHu3bZok=\",\"empty\":false,\"value\":\"fCYrYY6yqQ2Sq5A/T0jqN6mvTEzlpLTAObKOHu3bZok=\"}]"
	// baseResp, err := RestClient.CallContract(RestBizTestBizID,
	// 	orderId, RestBizTestAccount, RestBizTestTenantID,
	// 	ERC20, "balanceOf(identity)",
	// 	in, `["uint256"]`,
	// 	RestBizTestKmsID, false, gas)
	// fmt.Println(baseResp, "     err :", err)

	// hash := "0xa341faed8638bad30c0afdbcb1c734cfbabd9ac6f78da9689d48c0ae9f7a1c43"
	// baseResp, err := RestClient.QueryReceipt(RestBizTestBizID, hash)
	// if !(err == nil && baseResp.Code == "200") {
	// 	fmt.Println(fmt.Errorf("no succ resp baseResp:%+v err:%+v", baseResp, err))
	// }
	// fmt.Printf("baseResp:%+v\n", baseResp)
	// b1, _ := base64.StdEncoding.DecodeString("TbWK2pIhMEsWGrCqEbAypi1g2asp5BxtI49EAcgcgYk=")
	// fmt.Println(b)
	// fmt.Println(string(b))
	// data := hex.EncodeToString(b)
	// fmt.Println(data)
	c, _ := hex.DecodeString(accountIdentity[2:])
	fmt.Println("type:", reflect.TypeOf(c))
	fmt.Println(accountIdentity[2:])
	c1 := base64.StdEncoding.EncodeToString(c[:])
	fmt.Println("type:", reflect.TypeOf(c1))

}
