package controllers

import (
	"github.com/revel/revel"
	"encoding/csv"
	"fmt"
	"strings"
	"solarpower/app/models"
	"time"
	"strconv"
	"solarpower/app/db"
)

type LogData struct {
	*revel.Controller
}


func (c LogData) Show() revel.Result {
	var fetchedLogs []models.SolarPowerLog
	db.Gdb.Limit(30).Order("log_date desc").Find(&fetchedLogs)

	var logs []models.SolarPowerLog

	for i := len(fetchedLogs) - 1; i >= 0; i-- {
		revel.INFO.Println("%d", i)
		log := fetchedLogs[i]
		logs = append(logs, log)
	}
	return c.Render(logs)
}

func (c LogData) Add(logData string) revel.Result {

	c.Validation.Required(logData).Message("log data is required!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(LogData.Show)
	}

	revel.INFO.Println(logData)

	reader := csv.NewReader(strings.NewReader(logData))
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()

	if err != nil {
		msg := fmt.Sprintf("Error parsing CSV. Error : %s", err)
		revel.ERROR.Println(msg)
		c.Flash.Error(msg)
		c.FlashParams()
		return c.Redirect(LogData.Show)
	}

	for _, record := range csvData {
		var log models.SolarPowerLog

		yearStr := record[0]
		year, err := strconv.ParseInt(yearStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("year format invalid. year : %s, error : %s", year, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		monthStr := record[1]
		month, err := strconv.ParseInt(monthStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("month format invalid. month : %s, error : %s", month, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		dayStr := record[2]
		day, err := strconv.ParseInt(dayStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("day format invalid. day : %s, error : %s", day, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		hourStr := record[3]
		hour, err := strconv.ParseInt(hourStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("hour format invalid. hour : %s, error : %s", day, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}

		log.LogDate = time.Date(int(year), time.Month(month), int(day), int(hour), 0, 0, 0, time.UTC)



		generatedStr := "0"
		if record[4] != "" {
			generatedStr = record[4]
		}
		generatedFloat, err := strconv.ParseFloat(generatedStr, 32)
		if err != nil {
			msg := fmt.Sprintf("generated not in float format. generated : %s, error : %s", generatedStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Generated = float32(generatedFloat)

		soldStr := "0"
		if record[5] != "" {
			soldStr = record[5]
		}
		soldFloat, err := strconv.ParseFloat(soldStr, 32)
		if err != nil {
			msg := fmt.Sprintf("sold not in float format. sold : %s, error : %s", soldStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Sold = float32(soldFloat)

		boughtStr := "0"
		if record[6] != "" {
			boughtStr = record[6]
		}
		boughtFloat, err := strconv.ParseFloat(boughtStr, 32)
		if err != nil {
			msg := fmt.Sprintf("bought not in float format. bought : %s, error : %s", boughtStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Bought = float32(boughtFloat)

		usedStr := "0"
		if record[7] != "" {
			usedStr = record[7]
		}
		usedFloat, err := strconv.ParseFloat(usedStr, 32)
		if err != nil {
			msg := fmt.Sprintf("used not in float format. used : %s, error : %s", usedStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Used = float32(usedFloat)

		p1GeneratedStr := "0"
		if record[8] != "" {
			p1GeneratedStr = record[8]
		}
		p1GeneratedFloat, err := strconv.ParseFloat(p1GeneratedStr, 32)
		if err != nil {
			msg := fmt.Sprintf("p1Generated not in float format. p1Generated : %s, error : %s", p1GeneratedStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.P1Generated = float32(p1GeneratedFloat)

		externalGeneratedStr := "0"
		if record[11] != "" {
			externalGeneratedStr = record[11]
		}
		externalGeneratedFloat, err := strconv.ParseFloat(externalGeneratedStr, 32)
		if err != nil {
			msg := fmt.Sprintf("externalGenerated not in float format. externalGenerated : %s, error : %s", externalGeneratedStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.ExternalGenerated = float32(externalGeneratedFloat)

		incomeStr := "0"
		if record[17] != "" {
			incomeStr = record[17]
		}
		incomeInt, err := strconv.ParseInt(incomeStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("income not in int format. income : %s, error : %s", incomeStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Income = int32(incomeInt)

		outgoStr := "0"
		if record[18] != "" {
			outgoStr = record[18]
		}
		outgoInt, err := strconv.ParseInt(outgoStr, 10, 32)
		if err != nil {
			msg := fmt.Sprintf("outgo not in int format. outgo : %s, error : %s", outgoStr, err)
			revel.ERROR.Println(msg)
			c.Flash.Error(msg)
			c.FlashParams()
			return c.Redirect(LogData.Show)
		}
		log.Outgo = int32(outgoInt)

		var logQuery models.SolarPowerLog
		logQuery.LogDate = log.LogDate
		var storedLog models.SolarPowerLog

		db.Gdb.Where(logQuery).First(&storedLog)
		if storedLog.Id == 0 {
			db.Gdb.NewRecord(log)
			db.Gdb.Create(&log)
		}
	}


	return c.Redirect(LogData.Show)
}

