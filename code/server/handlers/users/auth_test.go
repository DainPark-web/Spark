package users

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestSignUp(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", SignUp)

}
