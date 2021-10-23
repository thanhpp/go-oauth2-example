package httpserver

import "github.com/gin-gonic/gin"

type FakeHandler struct{}

func (ctrl FakeHandler) A(c *gin.Context) {}
