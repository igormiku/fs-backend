package api

import (
	"fmt"
	"log"
	"net/http"
	"toggle_backend/services"

	"github.com/gorilla/mux"
)

type API struct {
	featuretoggleService services.FeatureToggleService
	Router               *mux.Router
}

func NewAPI(featuretoggleService services.FeatureToggleService) API {
	return API{
		featuretoggleService: featuretoggleService,
		Router:               mux.NewRouter().StrictSlash(true),
	}
}

func (api API) SetupAPI(apiBasePath string) {

	api.newAPI(apiBasePath, "v1", "features", "GET", api.featuresApiHandler)
	api.newAPI(apiBasePath, "v1", "features/{id}", "GET", api.featuresApiHandler)
	//CHROME CORS compliance
	api.newAPI(apiBasePath, "v1", "features/{id}", "OPTIONS", api.featuresOptionsApiHandler)
	api.newAPI(apiBasePath, "v1", "features", "POST", api.featuresApiHandler)
	//CHROME CORS compliance
	api.newAPI(apiBasePath, "v1", "features", "OPTIONS", api.featuresOptionsApiHandler)
	api.newAPI(apiBasePath, "v1", "features/add", "POST", api.featuresAddApiHandler)
	api.newAPI(apiBasePath, "v1", "features", "PUT", api.featuresApiHandler)
	api.newAPI(apiBasePath, "v1", "features", "DELETE", api.featuresApiHandler)

}

func (api API) newAPI(basePath string, apiVersion string, entityName string, method string, hf http.HandlerFunc) {
	urlPath := fmt.Sprintf("/%s/%s/%s",
		basePath,
		apiVersion,
		entityName,
	)

	api.Router.HandleFunc(urlPath, hf).Methods(method)
	log.Println("Registering endpoint  ", urlPath, " ", method)

}
