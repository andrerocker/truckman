package models

type Route struct {
	Path     string  `json:"path"`
	Distance float64 `json:"distance"`
	Cost     float64 `json:"cost"`
}

type RouteForm struct {
	Origin   string  `form:"origin" binding:"required"`
	Target   string  `form:"target" binding:"required"`
	Autonomy float64 `form:"autonomy" binding:"required"`
	FuelCost float64 `form:"fuelCost" binding:"required"`
}
