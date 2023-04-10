package helper

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Data struct {
	Id       string
	Title    string
	Date     string
	RealDate string
	Issame   bool
}

func ReadExcel() <-chan Data {
	chanOut := make(chan Data)
	xlsx, err := excelize.OpenFile("files/articles.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Sheet 1"
	rows := 100

	go func() {
		for i := 2; i < rows; i++ {
			row := Data{
				Id:       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
				Title:    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
				RealDate: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			}
			chanOut <- row
		}
		close(chanOut)
	}()

	return chanOut
}

func GetId(chanIn <-chan Data) <-chan string {
	chanOut := make(chan string)
	go func() {
		for item := range chanIn {
			chanOut <- item.Id
		}
		close(chanOut)
	}()

	return chanOut
}
