package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_getProduct(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	getProduct(ctx)
	t.Run("cannot find any product", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
