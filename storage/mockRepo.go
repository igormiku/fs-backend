package storage

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type MockDBImpl struct {
	FeatureToggles []FeatureToggle
}

func InitializeMockDB() DataAccessLayer {
	cur_unixtimestamp := strconv.FormatInt(time.Now().Unix(), 10)

	return &MockDBImpl{
		FeatureToggles: []FeatureToggle{
			{Id: 1, TechnicalName: "bitbucket", DisplayName: "BitBucket", Description: "Usage of BitBucket feature", Inverted: false, ExpiresOn: cur_unixtimestamp, CustomerIds: []string{"1", "2"}},
			{Id: 2, TechnicalName: "github", DisplayName: "GitHb", Description: "Usage of GitHub feature", Inverted: false, ExpiresOn: cur_unixtimestamp, CustomerIds: []string{"3", "4"}},
			{Id: 3, TechnicalName: "gitlab", DisplayName: "GitLb", Description: "Usage of GitLab feature", Inverted: false, ExpiresOn: cur_unixtimestamp, CustomerIds: []string{"1", "5"}},
		},
	}
}

func (b *MockDBImpl) Update(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {

	for index, ft2 := range b.FeatureToggles {
		if ft2.Id == ft.Id {

			//TODO : update only defined variables or ones which are not equal existing
			b.FeatureToggles[index] = ft
			return ft, nil
		}
	}

	err := errors.New(fmt.Sprintf("ID does not exist %s", fmt.Sprint(ft.Id)))

	return FeatureToggle{}, err
}

func remove(slice []FeatureToggle, s int) []FeatureToggle {
	return append(slice[:s], slice[s+1:]...)
}

func (b *MockDBImpl) Delete(ctx context.Context, ft FeatureToggle) error {

	for index, ft2 := range b.FeatureToggles {
		if ft2.Id == ft.Id {
			b.FeatureToggles = remove(b.FeatureToggles, index)
		}
	}
	return nil
}

func (b *MockDBImpl) Add(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {

	//TODO: take care of ID (unique etc...)
	ft.Id = rand.Intn(1000)
	b.FeatureToggles = append(b.FeatureToggles, ft)

	return ft, nil
}

func (b *MockDBImpl) Select(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {

	for _, ft2 := range b.FeatureToggles {
		if ft2.Id == ft.Id {
			return ft2, nil
		}
	}

	err := errors.New("ID does not exist")
	return FeatureToggle{}, err
}

func (b *MockDBImpl) SelectAll(ctx context.Context) ([]FeatureToggle, error) {
	return b.FeatureToggles, nil
}

func (b *MockDBImpl) Close() error {
	return nil
}
