package controller

import (
	"fmt"
	"log"
	"net/http"
	"qkeruen/models"
	"qkeruen/service"

	"github.com/gin-gonic/gin"
)

type securityController struct {
	SecurityService service.SecurityService
	JWTService      service.JWTService
}

func NewSecurityController(security service.SecurityService, jwtService service.JWTService) *securityController {
	controller := &securityController{
		SecurityService: security,
		JWTService:      jwtService,
	}

	return controller
}

func (s *securityController) Add(ctx *gin.Context) {
	var data models.Security

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}
	log.Println(data)

	if err := s.SecurityService.Create(data); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err.Error()),
			},
		)
		return
	}

	ctx.JSON(200, "created.")
}
