package web

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	logger "gitlab.alipay-inc.com/antchain/restclient-go-demo/logger"
	"gitlab.alipay-inc.com/antchain/restclient-go-demo/request"
	"gitlab.alipay-inc.com/antchain/restclient-go-sdk/client"
)

type ServerPort struct {
	Port string `json:"port"`
}

func init() {
	var err error
	//configFilePath := os.Getenv("GOPATH") + "src/gitlab.alipay-inc.com/antchain/restclient-go-demo/rest-config.json"
	configFilePath := "rest-config.json"
	request.RestClient, err = client.NewRestClient(configFilePath)
	if err != nil {
		logger.L().Debug(fmt.Errorf("failed to NewRestClient err:%+v", err))
	}
	if request.RestClient.RestToken == "" {
		logger.L().Debug(fmt.Errorf("rest token:%+v is empty", request.RestClient.RestToken))
	}
}

func Service() {
	r := gin.Default()
	r.Use(GinRecovery(true))
	request.Query(r)
	request.CallContractERC20(r)
	r.Run(getPort())
}

func getPort() string {
	return request.Port
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.L().Errorf("%v %v %v %v %v (%v) %v", c.Writer.Status(), c.Request.Method, c.Request.URL.RequestURI(), "", c.Request.UserAgent(), c.ClientIP())
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.L().Errorf("[Recovery from panic]\n%v%v\n%v", string(httpRequest), err, string(debug.Stack()))
				} else {
					logger.L().Errorf("[Recovery from panic]\n%v%v", string(httpRequest), err)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
