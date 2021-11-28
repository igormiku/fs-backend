package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/tidwall/buntdb"
)

type buntDBImpl struct {
	DB *buntdb.DB
}

func InitializeBuntDB(file string) DataAccessLayer {
	//make a config
	db, err := buntdb.Open("db/data.db")
	if err != nil {
		log.Fatal(err)
	}

	return &buntDBImpl{DB: db}
}

func (b *buntDBImpl) Update(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {
	jsonFT, err := json.Marshal(ft)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return FeatureToggle{}, err
	}

	err2 := b.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err2 := tx.Set(strconv.Itoa(ft.Id), string(jsonFT), nil)
		return err2
	})

	if err2 != nil {
		fmt.Printf("Error: %s", err2)
		return FeatureToggle{}, err2
	}
	return ft, nil
}

func (b *buntDBImpl) Delete(ctx context.Context, ft FeatureToggle) error {

	err2 := b.DB.Update(func(tx *buntdb.Tx) error {
		_, err2 := tx.Delete(strconv.Itoa(ft.Id))
		return err2
	})

	if err2 != nil {
		fmt.Printf("Error: %s", err2)
		return err2
	}
	return nil
}

func (b *buntDBImpl) Add(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {

	jsonFT, err := json.Marshal(ft)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return FeatureToggle{}, err
	}

	err2 := b.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err2 := tx.Set(strconv.Itoa(ft.Id), string(jsonFT), nil)
		return err2
	})

	if err2 != nil {
		fmt.Printf("Error: %s", err2)
		return FeatureToggle{}, err2
	}
	return ft, nil
}

func (b *buntDBImpl) Select(ctx context.Context, ft FeatureToggle) (FeatureToggle, error) {

	var ft2 FeatureToggle

	err2 := b.DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(strconv.Itoa(ft.Id))
		if err != nil {
			return err
		}
		json.Unmarshal([]byte(val), &ft2)
		//add error handling
		return nil
	})

	if err2 != nil {
		fmt.Printf("Error: %s", err2)
		return FeatureToggle{}, err2
	}

	return ft2, nil
}

func (b *buntDBImpl) SelectAll(ctx context.Context) ([]FeatureToggle, error) {
	var res []FeatureToggle
	var ft2 FeatureToggle

	err2 := b.DB.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			json.Unmarshal([]byte(value), &ft2)
			//add unmarshal error handling
			res = append(res, ft2)
			//fmt.Printf("key: %s, value: %s\n", key, value)
			return true
		})
		return err
	})

	if err2 != nil {
		fmt.Printf("Error: %s", err2)
		return []FeatureToggle{}, err2
	}

	return res, nil
}

func (b *buntDBImpl) Close() error {
	b.DB.Close()
	return nil
}
