package v0Service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (s *service) RegisterAccount(c *gin.Context) {
	reqMsg, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read from http body err:%s\n", err)
		return nil
	}
}
