package main

import (
	"fmt"
	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ziggornif.xyz/ddd-gym/domains/formules/adapter"
	test_adapter "ziggornif.xyz/ddd-gym/domains/formules/test/adapter"
	"ziggornif.xyz/ddd-gym/domains/formules/usecase"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	repository := test_adapter.NewSpyRepository()
	dispatcher := adapter.NewJetStreamDispatcher()
	useCase := usecase.NewCreateFormuleUseCase(repository, dispatcher)

	router.GET("/formules", func(c *gin.Context) {
		genID, _ := uuid.NewUUID()
		useCase.Execute(genID.String(), fmt.Sprintf("%v-%v", "formule", genID.String()), rand.Int())
		c.JSON(200, gin.H{
			"message": "Formule created",
			"id":      genID.String(),
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
