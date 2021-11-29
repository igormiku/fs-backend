package api

import (
	"net/http"
	"strconv"
	"toggle_backend/storage"

	"github.com/gorilla/mux"
)

func (api API) featuresApiHandler(w http.ResponseWriter, r *http.Request) {
	api.setupHeader(&w, r)

	switch r.Method {
	//return all list
	case http.MethodGet:
		var params = mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			result, err := api.featuretoggleService.GetAllFeatures(r.Context())
			api.handleResult(result, err, w)
			return
		}

		result, err := api.featuretoggleService.GetFeature(r.Context(), storage.FeatureToggle{Id: id})
		api.handleResult(result, err, w)

	//update certain Feature
	case http.MethodPut:
		feature, err := api.unmarshalToFeature(w, r)
		if err != nil {
			return
		}

		result, err := api.featuretoggleService.UpdateFeature(r.Context(), feature)
		api.handleResult(result, err, w)
	case http.MethodDelete:
		feature, err := api.unmarshalToFeature(w, r)
		if err != nil {
			return
		}
		if err := api.featuretoggleService.DeleteFeature(r.Context(), feature); err != nil {
			api.writeErrorJSON(w, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	//return features for customerid
	case http.MethodPost:
		feature, err := api.unmarshalToFeatureByCustomerId(w, r)
		if err != nil {
			return
		}
		result, err := api.featuretoggleService.GetAllByCustomerID(r.Context(), feature)
		api.handleResult(result, err, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (api API) featuresOptionsApiHandler(w http.ResponseWriter, r *http.Request) {
	api.setupHeader(&w, r)

	switch r.Method {
	case http.MethodOptions:
		{
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (api API) featuresAddApiHandler(w http.ResponseWriter, r *http.Request) {
	api.setupHeader(&w, r)

	switch r.Method {
	case http.MethodPost:
		feature, err := api.unmarshalToFeature(w, r)
		if err != nil {
			return
		}
		result, err := api.featuretoggleService.CreateFeature(r.Context(), feature)
		api.handleResult(result, err, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
