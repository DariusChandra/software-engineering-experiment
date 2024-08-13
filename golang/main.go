package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the local time
	localTime := time.Now()

	// Get the location of the local time
	location := localTime.Location()

	// Print the timezone in the format "America/Santiago"
	fmt.Println("Timezone:", location.String())

	location, errTimezone := time.LoadLocation(location.String())
	fmt.Println(location, errTimezone)
}
