package main

import (
	"fmt"
	"golang/itassets" 
)

func main() {
	a := itassets.Asset{
		AssetTag:     "LAP001",
		DeviceModel:  "Dell Laptop",
		PurchaseCost: 55000,
	}

	err := itassets.RegisterAsset(a)
	if err != nil {
		fmt.Println("Error registering asset:", err)
		return
	}

	x, err := itassets.LookupAsset("LAP001")
	if err != nil {
		fmt.Println("Error finding asset:", err)
		return
	}
	fmt.Println("Asset Registered Successfully")
	fmt.Println("Asset Tag:", x.AssetTag)
	fmt.Println("Device Model:", x.DeviceModel)
	fmt.Println("Purchase Cost:", x.PurchaseCost)
	fmt.Println("Status:", x.Status)

	err = itassets.AssignAsset("LAP001", "Suresh")
	if err != nil {
		fmt.Println("Error assigning asset:", err)
		return
	}

	x, err = itassets.LookupAsset("LAP001")
	if err != nil {
		fmt.Println("Error finding asset:", err)
		return
	}
	fmt.Println("\nAfter Assignment:")
	fmt.Println("Assigned User:", x.AssignedUser)
	fmt.Println("Status:", x.Status)

	err = itassets.RetireAsset("LAP001")
	if err != nil {
		fmt.Println("Error retiring asset:", err)
		return
	}
	fmt.Println("\nAsset Retired Successfully")
}
