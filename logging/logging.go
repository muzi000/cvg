package logging

import (
	"fmt"
	"time"
)

func FormatLog(scale, txt string) string {
	switch scale {
	case "info":
		return "[info] " + time.Now().Format(time.Layout) + " " + txt
	case "wanning":
		return "[wanning] " + time.Now().Format(time.Layout) + " " + txt
	case "err":
		return "[error] " + time.Now().Format(time.Layout) + " " + txt
	}
	return "[error] " + time.Now().Format(time.Layout) + " " + txt
}

func LogPrint(scale, txt string) {
	fmt.Println(FormatLog(scale, txt))
}
