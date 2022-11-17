package request

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
)

// const (
// 	RestBizTestBizID    = "4b4e60e8feaf4ce0a316fc11344b8784"
// 	RestBizTestAccount  = "qiuguowenhua"
// 	RestBizTestKmsID    = "15681561400223758a59e254-7c25-46fd-b1e6-493cf3c28af7"
// 	RestBizTestTenantID = "1568156140022375"

// 	ERC20 = "erc20-test5"
// )

var (
	RestBizBizID    string
	RestBizAccount  string
	RestBizKmsID    string
	RestBizTenantID string

	ERC20   string
	ERC721  string
	ERC1155 string

	Port string
)

var RestClient *client.RestClient

func GetIdentity(name string) {}

type Identity struct {
	Data  string `json:"data"`
	Empty bool   `json:"empty"`
	Value string `json:"value"`
}

type env struct {
	RestBizBizID    string `json:"RestBizBizID"`
	RestBizAccount  string `json:"RestBizAccount"`
	RestBizKmsID    string `json:"RestBizKmsID"`
	RestBizTenantID string `json:"RestBizTenantID"`

	ERC20   string `json:"ERC20"`
	ERC721  string `json:"ERC721"`
	ERC1155 string `json:"ERC1155"`

	Port string `json:"port"`
}

func init() {

	envByte, err := ioutil.ReadFile("env.json")
	if err != nil {
		panic(err)
	}
	envjson := env{}
	err = jsoniter.Unmarshal(envByte, &envjson)
	if err != nil {
		panic(err)
	}

	RestBizBizID = envjson.RestBizBizID
	RestBizAccount = envjson.RestBizAccount
	RestBizKmsID = envjson.RestBizKmsID
	RestBizTenantID = envjson.RestBizTenantID

	ERC20 = envjson.ERC20
	ERC721 = envjson.ERC721
	ERC1155 = envjson.ERC1155

	Port = envjson.Port
}

func InputParamIdentity(identity string) interface{} {
	cmp := strings.Compare("0x", identity[:2]) //比较开头是否是地址还是name
	if cmp == 0 {
		return InputParamIdentityFromIdentity(identity)
	}
	return InputParamIdentityFromName(identity)
}

func InputParamIdentityFromName(name string) interface{} {
	name_btye32 := sha256.Sum256([]byte(name))
	name_base64_string := base64.StdEncoding.EncodeToString(name_btye32[:])
	identity := Identity{name_base64_string, false, name_base64_string}
	return identity
}
func InputParamIdentityFromIdentity(name string) interface{} {
	_identity, _ := hex.DecodeString(name[2:])
	_identity_string := base64.StdEncoding.EncodeToString(_identity[:])
	identity := Identity{_identity_string, false, _identity_string}
	return identity
}
func IsName(str string) bool {
	cmp := strings.Compare("0x", str[:2]) //比较开头是否是地址还是name
	return cmp != 0
}
