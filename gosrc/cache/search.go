package cache

import (
	"foodies/gosrc/model"
	"strings"
)

func Search(searchBy, searchWhat string) (rtrn []model.FoodTruck) {
	switch searchBy {
	case "truck":
		for _, truck := range store.data {
			if ContainsI(truck.Applicant, searchWhat) {
				rtrn = append(rtrn, truck)
			}
		}
	case "cousine":
		for _, truck := range store.data {
			if ContainsI(truck.Fooditems, searchWhat) {
				rtrn = append(rtrn, truck)
			}
		}
	case "address":
		for _, truck := range store.data {
			if ContainsI(truck.Address, searchWhat) {
				rtrn = append(rtrn, truck)
			}
		}
	}
	return
}

func ContainsI(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}
