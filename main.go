package main

import (
	// "strings"
	"fmt"
	"time"
)

var INTRO string = "Isch"
var MINS = map[int]string{
	0:  "",
	5:  "foif",
	10: "zah",
	15: "foifzah",
	20: "zwanzg",
	25: "foifezwanzg",
	30: "halbi",
}
var ADJECTIVES = map[int]string{
	30: "nach",
	60: "vor",
}
var HOURS = map[int]string{
	0:  "nul",
	1:  "eis",
	2:  "zwoi",
	3:  "druu",
	4:  "vier",
	5:  "foif",
	6:  "sach",
	7:  "sibe",
	8:  "acht",
	9:  "nuun",
	10: "zah",
	11: "euf",
	12: "zwouf",
	13: "dryzah",
	14: "vierzah",
	15: "foifzah",
	16: "sachzah",
	17: "sibezah",
	18: "achtzah",
	19: "nuunzah",
	20: "zwanzg",
	21: "einezwanzg",
	22: "zwoiezwanzg",
	23: "druuezwanzg",
}

func main() {
	// get and store the current time
	theTime := time.Now()

	// declare all the output vars
	theHour := ""
	theAdjective := ""
	theMinute := ""

	// get the adjective
	if theTime.Minute() < 30 {
		if theTime.Minute() >= 0 && theTime.Minute() < 5 {
			theAdjective = ""
		} else {
			theAdjective = ADJECTIVES[30]
		}
		theHour = HOURS[(theTime.Hour())]
	} else {
		theAdjective = ADJECTIVES[60]
		theHour = HOURS[(theTime.Hour() + 1)]
	}

	// select the correct minute string
	switch thisMinute := theTime.Minute(); {
	case (thisMinute >= 0 && thisMinute < 5):
		theMinute = MINS[0]
	case (thisMinute >= 5 && thisMinute < 10):
		theMinute = MINS[5]
	case (thisMinute >= 10 && thisMinute < 15):
		theMinute = MINS[10]
	case (thisMinute >= 15 && thisMinute < 20):
		theMinute = MINS[15]
	case (thisMinute >= 20 && thisMinute < 25):
		theMinute = MINS[20]
	case (thisMinute >= 25 && thisMinute < 30):
		theMinute = MINS[25]
	case (thisMinute >= 30 && thisMinute < 35):
		theMinute = MINS[30]
	case (thisMinute >= 35 && thisMinute < 40):
		theMinute = MINS[25]
	case (thisMinute >= 40 && thisMinute < 45):
		theMinute = MINS[20]
	case (thisMinute >= 45 && thisMinute < 50):
		theMinute = MINS[15]
	case (thisMinute >= 50 && thisMinute < 55):
		theMinute = MINS[10]
	case (thisMinute >= 55 && thisMinute < 60):
		theMinute = MINS[5]
	}

	// print that out
	fmt.Printf("%s %s %s %s\n", INTRO, theMinute, theAdjective, theHour)
}
