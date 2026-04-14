package domain

type Car struct {
	ID           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Engine       string  `json:"engine"`
	CapacityCC   int     `json:"capacity_cc"`
	HorsePower   int     `json:"horse_power"`
	TopSpeedKMH  int     `json:"top_speed_kmh"`
	Acceleration float64 `json:"acceleration"`
	Price        float64 `json:"price"`
	FuelType     string  `json:"fuel_type"`
	Seats        int     `json:"seats"`
	TorqueNM     int     `json:"torque_nm"`
}
