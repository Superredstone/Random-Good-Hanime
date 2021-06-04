package RandomGoodHanime

import (
	"fmt"
	"time"
)

func Cron(commandToPrint, msgUsername, msgFirstName string) {
	dt := time.Now()
	dtFormatted := dt.Format("01-02-2006 15:04:05")

	FormattedNice := "[" + dtFormatted + "] @" + msgUsername + " " + msgFirstName + " ==> " + commandToPrint + "\n"

	Log = Log + FormattedNice + "\n"

	fmt.Println(FormattedNice)
}
