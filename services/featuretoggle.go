package services

import (
	"context"
	"strconv"
	"time"
	"toggle_backend/storage"
)

type FeatureToggleService struct {
	DB storage.DataAccessLayer
}

func NewFeatureToggleService(db storage.DataAccessLayer) FeatureToggleService {
	return FeatureToggleService{
		DB: db,
	}
}

func (fts FeatureToggleService) GetAllFeatures(ctx context.Context) ([]storage.FeatureToggle, error) {
	result, err := fts.DB.SelectAll(ctx)
	if err != nil {
		return []storage.FeatureToggle{}, err
	}
	return result, nil
}

func (fts FeatureToggleService) UpdateFeature(ctx context.Context, ft storage.FeatureToggle) (storage.FeatureToggle, error) {
	result, err := fts.DB.Update(ctx, ft)
	if err != nil {
		return storage.FeatureToggle{}, err
	}
	return result, nil
}

func (fts FeatureToggleService) GetFeature(ctx context.Context, ft storage.FeatureToggle) (storage.FeatureToggle, error) {
	result, err := fts.DB.Select(ctx, ft)
	if err != nil {
		return storage.FeatureToggle{}, err
	}
	return result, nil
}

func (fts FeatureToggleService) DeleteFeature(ctx context.Context, ft storage.FeatureToggle) error {
	err := fts.DB.Delete(ctx, ft)
	if err != nil {
		return err
	}
	return nil
}

func (fts FeatureToggleService) CreateFeature(ctx context.Context, ft storage.FeatureToggle) (storage.FeatureToggle, error) {
	result, err := fts.DB.Add(ctx, ft)
	if err != nil {
		return storage.FeatureToggle{}, err
	}
	return result, nil
}

func (fts FeatureToggleService) GetAllByCustomerID(ctx context.Context, ft storage.FeatureByCustomerIdRequestAPI) ([]storage.FeatureByCustomerIdResponseAPI, error) {

	found := false
	var reply []storage.FeatureByCustomerIdResponseAPI

	FeatureToggleList, err := fts.GetAllFeatures(ctx)
	if err != nil {
		return []storage.FeatureByCustomerIdResponseAPI{}, err
	}

	for _, name := range ft.Features {
		for _, featureToggle := range FeatureToggleList {
			found = false
			//add to results only if feature found

			if featureToggle.TechnicalName == name {
				//check for customer subscription
				for _, customer := range featureToggle.CustomerIds {
					if ft.CustomerId == customer {
						found = true
					}
					//log.Printf("customer %s  ft.CustomerId %s  \n", customer, ft.CustomerId)
				}

				expiresOnInt64, err := strconv.ParseInt(featureToggle.ExpiresOn, 10, 64)
				if err != nil {
					return []storage.FeatureByCustomerIdResponseAPI{}, err
				}

				expired := (time.Now().Unix() > expiresOnInt64)
				active := (!featureToggle.Inverted && found)

				reply = append(reply, storage.FeatureByCustomerIdResponseAPI{Name: featureToggle.TechnicalName,
					Active: active, Inverted: featureToggle.Inverted, Expired: expired})

			}
		}

	}

	return reply, nil

}
