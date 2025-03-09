package users

import "github.com/gin-gonic/gin"

type Response struct {
	Ok bool `json:"ok"`
}

type SignUpResponse struct {
	Ok   bool  `json:"ok"`
	Data gin.H `json:"data"`
}
