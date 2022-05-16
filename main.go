package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brackbk/purchase-service/controller"
	"github.com/brackbk/purchase-service/logger"
	"github.com/brackbk/purchase-service/repository"
	"github.com/brackbk/purchase-service/service"
	"github.com/brackbk/purchase-service/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var client *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client = utils.ConnectDB()
}

func main() {
	logger.Info("Purchase-service starting...")
	router := mux.NewRouter()

	// CONTROLLERS

	purchaseController := controller.PurchaseController{
		Service: service.NewDefaultPurchaseService(repository.NewPurchaseRepositoryAdapter(client)),
	}
	purchaseItemController := controller.PurchaseItemController{
		Service:         service.NewDefaultPurchaseItemService(repository.NewPurchaseItemRepositoryAdapter(client)),
		PurchaseService: service.NewDefaultPurchaseService(repository.NewPurchaseRepositoryAdapter(client)),
	}

	// ROUTERS

	/*router.HandleFunc("/refresh", userController.Refresh).Methods(http.MethodPost)*/

	router.HandleFunc("/purchase", purchaseController.Create).Methods(http.MethodPost)
	router.HandleFunc("/purchase", purchaseController.Get).Methods(http.MethodGet)
	router.HandleFunc("/purchase", purchaseController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/purchase", purchaseController.Update).Methods(http.MethodPut)
	router.HandleFunc("/purchase/list", purchaseController.List).Methods(http.MethodPost)

	router.HandleFunc("/purchaseitem", purchaseItemController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/purchaseitem", purchaseItemController.Update).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("Port")), router))
}
