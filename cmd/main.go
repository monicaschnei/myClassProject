package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"log"
	"myclass/api"
	db "myclass/db/sqlc"
	"myclass/internal/controller"
	"myclass/internal/core/service"
	"myclass/util"
)

func main() {
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())
	config, err := util.LoadConfig(".")
	tokenMaker, _ := token.NewPasetoMaker(config.AccesSymetryTokenKey)
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	//Inicializa o db connection
	connection, err := sql.Open(config.DBDrive, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	//Cria o repositório
	store := db.NewStore(connection)

	//Cria Service
	professionalUserService := usecase.NewProfessionalUserService(store, api.NewServer(instance, config))
	professionalAvailabilityService := usecase.NewProfessionalAvailabilityService(store, api.NewServer(instance, config), professionalUserService)

	//Cria Controller
	professionalUserController := controller.NewProfessionalUserController(instance, professionalUserService)

	//Inicializa as rotas do controller
	professionalUserController.InitRouter()
	availabilityProfessionalUserController.InitRouter()

	//Cria o httpServer
	server := api.NewServer(instance, config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	//Start httpServer
	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	log.Printf("listening port %s\n", config.Port)
}
