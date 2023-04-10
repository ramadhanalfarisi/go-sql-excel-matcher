package main

import (

	"github.com/ramadhanalfarisi/excel-sql-matcher/helper"
)

func main() {
	db := helper.Connect()
	defer db.Close()

	dataExcel := helper.ReadExcel()

	getId := helper.GetId(dataExcel)

	readDataSQL := helper.ReadData(db, getId)

	dataExcel2 := helper.ReadExcel()

	helper.Matcher(readDataSQL, dataExcel2)

}
