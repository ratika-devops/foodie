package model

type FoodTruckLocation struct {
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	HumanAddress string `json:"human_address"`
}
type FoodTruck struct {
	Objectid               string            `json:"objectid"`
	Applicant              string            `json:"applicant"`
	Facilitytype           string            `json:"facilitytype"`
	Cnn                    string            `json:"cnn"`
	Locationdescription    string            `json:"locationdescription"`
	Address                string            `json:"address"`
	Blocklot               string            `json:"blocklot"`
	Block                  string            `json:"block"`
	Lot                    string            `json:"lot"`
	Permit                 string            `json:"permit"`
	Status                 string            `json:"status"`
	Fooditems              string            `json:"fooditems"`
	X                      string            `json:"x"`
	Y                      string            `json:"y"`
	Latitude               string            `json:"latitude"`
	Longitude              string            `json:"longitude"`
	Schedule               string            `json:"schedule"`
	Approved               string            `json:"approved"`
	Received               string            `json:"received"`
	Priorpermit            string            `json:"priorpermit"`
	Expirationdate         string            `json:"expirationdate"`
	Location               FoodTruckLocation `json:"location"`
}
