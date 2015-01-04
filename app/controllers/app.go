package controllers

import "github.com/revel/revel"
// import "solarpower/app/models"
// import "time"
// import "solarpower/app/db"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// log := models.SolarPowerLog{
	// 	LogDate: time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local),
	// 	Generated: 0.5,
	// 	Sold: 25.4,
	// 	Bought: 32.5,
	// 	Used: 0.4,
	// 	P1Generated: 0.5,
	// 	ExternalGenerated: 50.3,
	// }

	// db.Gdb.NewRecord(log)
	// db.Gdb.Create(&log)

	return c.Render()
}


