package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aaomidi/monitor-control-go/api"
)

func main() {
	brightness, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Setting brightness to: %d\n", brightness)
	api.GetMonitors(brightness)
}
