package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"toggle_backend/storage"
)

func (api API) writeErrorJSON(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Printf("Failed with: %+v", err)
	errorJSON, err := json.Marshal(err)
	if err != nil {
		log.Printf("Error while marshalling error: %+v", err)
	}
	w.Write(errorJSON)
}

func (api API) handleResult(i interface{}, err error, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		return
	}
	jsonResponse, err := api.marshalObject(i)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		return
	}
	//log.Printf("jsonResponse with: %s\n", string(jsonResponse))
	w.Write(jsonResponse)

}

func (api API) marshalObject(i interface{}) ([]byte, error) {
	resultJSON, err := json.Marshal(&i)
	if err != nil {
		return nil, err
	}
	return resultJSON, nil
}

func (api API) unmarshallFeature(body []byte) (storage.FeatureToggle, error) {
	var feature storage.FeatureToggle
	if err := json.Unmarshal(body, &feature); err != nil {
		return storage.FeatureToggle{}, err
	}
	return feature, nil
}

func (api API) unmarshallFeatureByCustomerId(body []byte) (storage.FeatureByCustomerIdRequestAPI, error) {
	var feature storage.FeatureByCustomerIdRequestAPI
	if err := json.Unmarshal(body, &feature); err != nil {
		return storage.FeatureByCustomerIdRequestAPI{}, err
	}
	return feature, nil
}

func (api API) unmarshalToFeature(w http.ResponseWriter, r *http.Request) (storage.FeatureToggle, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		log.Printf("Cannot parse the body: %+v\n", err)
		return storage.FeatureToggle{}, err
	}
	ds, err := api.unmarshallFeature(body)
	if err != nil {
		log.Printf("Cannot unmarshal the body: %+v %s \n", err, string(body))
		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		return storage.FeatureToggle{}, err
	}
	return ds, nil
}

func (api API) unmarshalToFeatureByCustomerId(w http.ResponseWriter, r *http.Request) (storage.FeatureByCustomerIdRequestAPI, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		log.Printf("Cannot parse the body: %+v\n", err)
		return storage.FeatureByCustomerIdRequestAPI{}, err
	}
	ds, err := api.unmarshallFeatureByCustomerId(body)
	if err != nil {
		log.Printf("Cannot unmarshal the body: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		api.writeErrorJSON(w, err)
		return storage.FeatureByCustomerIdRequestAPI{}, err
	}
	return ds, nil
}
