package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"yuque-webhook-wecom/config"
	"yuque-webhook-wecom/storage"

	"github.com/labstack/echo/v4"
)

type Hook struct {
	parser *storage.YuqueParser
}

func NewHook(parser *storage.YuqueParser) *Hook {
	return &Hook{parser: parser}
}

func (h *Hook) Handler(c echo.Context) (err error) {
	token := c.Param("token") // 验证 token
	log.Printf("token is %s\n", token)
	if token != config.Config.GetString("token.webhook") {
		log.Printf("token 校验失败")
		return c.JSON(http.StatusInternalServerError, "")
	}
	tMap := make(map[string]interface{})

	if err = json.NewDecoder(c.Request().Body).Decode(&tMap); err != nil {
		log.Printf("decode err: %v\n", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	if err = h.parser.Parser(tMap); err != nil {
		log.Printf("parser err: %v\n", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "ok")
}
