package helper

import (
	"fmt"
	"log"
	"strings"

	go_logging "github.com/ramadhanalfarisi/go-logging"
)

func Matcher(arrSQL []DataSql, chanIn <-chan Data) {
	get_config := go_logging.ConfigLogging(true, "logs/")
	countSuccess := 0
	countSame := 0
	countNotSame := 0
	for item := range chanIn {
		for _, sql := range arrSQL {
			if item.Id == sql.Id {
				item.Date = sql.Date
				if item.RealDate == strings.TrimLeft(sql.Date, "0") {
					item.Issame = true
				} else {
					item.Issame = false
				}
				log.Println("Data has been matched")
				countSuccess++
				if !item.Issame {
					get_config.Info(fmt.Sprintf("Id : %s, %s (Fake) => %s (Real)", item.Id, item.Date, item.RealDate))
					countNotSame++
				} else {
					countSame++
				}
				break
			}
		}
	}
	get_config.Info(fmt.Sprintf("%d Data Executed!", countSuccess))
	get_config.Info(fmt.Sprintf("%d Data Correct!", countSame))
	get_config.Info(fmt.Sprintf("%d Data Incorrect!", countNotSame))
}
