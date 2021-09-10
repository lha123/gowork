package result

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code""`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	obj := &Response{http.StatusOK, "success", data}
	marshal, err := json.Marshal(obj)
	if err != nil {
		return
	}
	json.Unmarshal(marshal, obj)
	c.JSON(http.StatusOK, obj)
}

func ResponseError(c *gin.Context) {
	obj := &Response{Code: http.StatusInternalServerError, Msg: "success"}
	marshal, err := json.Marshal(obj)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, string(marshal))
}
