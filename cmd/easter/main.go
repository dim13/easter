package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dim13/easter"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "year")
	dfmt := flag.String("format", "Mon _2 Jan 2006", "date format")
	flag.Parse()

	fmt.Println("Easter", easter.Easter(*year).Format(*dfmt))
	fmt.Println("Paskha", easter.Paskha(*year).Format(*dfmt))
	fmt.Println("Pesach", easter.Pesach(*year).Format(*dfmt))
}
