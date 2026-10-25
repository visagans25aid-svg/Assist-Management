package itassets

import "errors"
type Asset struct {
	AssetTag     string
	DeviceModel  string
	AssignedUser string
	PurchaseCost float64
	Status       string
}
var assets = make(map[string]Asset)

func RegisterAsset(a Asset) error {
	if _, ok := assets[a.AssetTag]; ok {
		return errors.New("asset already exists")
	}
	a.Status = "In-Stock"
	assets[a.AssetTag] = a
	return nil
}

func LookupAsset(tag string) (*Asset, error) {
	a, ok := assets[tag]
	if !ok {
		return nil, errors.New("asset not found")
	}
	return &a, nil
}

func AssignAsset(tag, username string) error {
	a, ok := assets[tag]
	if !ok {
		return errors.New("asset not found")
	}
	a.AssignedUser = username
	a.Status = "Deployed"
	assets[tag] = a
	return nil
}

func RetireAsset(tag string) error {
	if _, ok := assets[tag]; !ok {
		return errors.New("asset not found")
	}
	delete(assets, tag)
	return nil
}
