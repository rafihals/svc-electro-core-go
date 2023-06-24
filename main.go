package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"svc-boilerplate-golang/utils/database"

	_boilerplateHttpDeliver "svc-boilerplate-golang/domain/boilerplate/delivery/http"
	_boilerplateRepository "svc-boilerplate-golang/domain/boilerplate/repository"
	_boilerplateUsecase "svc-boilerplate-golang/domain/boilerplate/usecase"
)

func main() {
	routers := gin.Default()

	mysql := ConnectMySQL()
	//orcl:= ConnectOracle()

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
	mux := http.NewServeMux()
	mux.HandleFunc("/plm/cors", Cors)
	http.ListenAndServe(":80", mux)
}

func ConnectMySQL() (mysql *sql.DB) {
	mysql, err := database.SetupMysqlDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func ConnectOracle() (oracle *sql.DB) {
	oracle, err := database.SetupOracleDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}
