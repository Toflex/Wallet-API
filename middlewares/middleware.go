package middlewares

import (
	"bytes"
	"fmt"
	"github.com/Toflex/Wallet-API/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type APIRequest struct {
	TimeStamp string
	Method string
	Path string
	Request string

}

type Middleware struct {
	Serv services.Service
}

func (m *Middleware) LogRequestMiddleware(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	apiResponse:= &APIRequest{

		TimeStamp: time.Now().Format("02-01-2006 15:01:05"),
		Method:    c.Request.Method,
		Path:      c.Request.RequestURI,
		Request:  string(body),
	}

	fmt.Printf("[%s] %s %s\n Request: %s\n", apiResponse.TimeStamp, apiResponse.Method, apiResponse.Path, apiResponse.Request)
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {

	auth:= c.Request.Header.Get("Authorization")
	jwt:=strings.Split(auth, " ")

	token, Tokenerr := m.Serv.ValidateToken(jwt[len(jwt)-1])
	if Tokenerr != nil {
		log.Error("An error occurred when verifying jwt token. %s", Tokenerr.Error())
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	if token.Valid == false {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
}
