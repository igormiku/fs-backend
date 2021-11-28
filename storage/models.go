package storage

import "context"

type FeatureToggle struct {
	Id            int      `json:"id"`
	DisplayName   string   `json:"displayname"`
	TechnicalName string   `json:"technicalname"`
	ExpiresOn     string   `json:"expireson"`
	Description   string   `json:"description"`
	Inverted      bool     `json:"inverted"`
	CustomerIds   []string `json:"customerids"`
}

type DataAccessLayer interface {
	Add(context.Context, FeatureToggle) (FeatureToggle, error)
	Delete(context.Context, FeatureToggle) error
	Select(context.Context, FeatureToggle) (FeatureToggle, error)
	Update(context.Context, FeatureToggle) (FeatureToggle, error)
	SelectAll(context.Context) ([]FeatureToggle, error)
	Close() error
}

type FeatureByCustomerIdRequestAPI struct {
	CustomerId string   `json:"customerId"`
	Features   []string `json:"features"`
}

type FeatureByCustomerIdResponseAPI struct {
	Name     string `json:"name"`
	Active   bool   `json:"active"`
	Inverted bool   `json:"inverted"`
	Expired  bool   `json:"expired"`
}
