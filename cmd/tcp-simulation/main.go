package main

import (
	"fmt"
	"log"
	"net/http"
	"tcp-simulation/pkg/api"
	"tcp-simulation/pkg/service"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func main() {
	a := App{}
	a.routes()
	a.run(":8000")
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) routes() {
	a.Router = mux.NewRouter()
	postAPI := InitPackageAPI()
	aimdAPI := InitAimdAPI()
	slowStartAPI := InitSlowStartAPI()
	analysisAPI := InitAnalysisAPI()
	a.Router.HandleFunc("/packages", postAPI.GetAllPackage()).Methods("GET")
	a.Router.HandleFunc("/packages/generate", postAPI.GenerateLossPackages()).Methods("GET")
	a.Router.HandleFunc("/aimd", aimdAPI.GenerateGrafic()).Methods("GET")
	a.Router.HandleFunc("/slowstart", slowStartAPI.GenerateGrafic()).Methods("GET")
	a.Router.HandleFunc("/analysis", analysisAPI.GenerateGrafic()).Methods("GET")
}

func InitPackageAPI() api.PackageAPI {
	postService := service.NewPackageService()
	packageAPI := api.NewPackageAPI(postService)
	return packageAPI
}

func InitAimdAPI() api.AimdAPI {
	aimdService := service.NewAimdService()
	aimdAPI := api.NewAimdAPI(aimdService)
	return aimdAPI
}

func InitSlowStartAPI() api.SlowStartAPI {
	slowstartService := service.NewSlowStartService()
	slowstartAPI := api.NewSlowStartAPI(slowstartService)
	return slowstartAPI
}

func InitAnalysisAPI() api.AnalysisAPI {
	slowstartService := service.NewSlowStartService()
	aimdService := service.NewAimdService()
	analysisAPI := api.NewAnalysisAPI(aimdService, slowstartService)
	return analysisAPI
}
