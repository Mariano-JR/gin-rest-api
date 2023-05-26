package routes

import (
	"github.com/Mariano-JR/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/alunos", controllers.ExibeTodosALunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarALunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run()
}
