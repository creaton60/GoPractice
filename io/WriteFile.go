package main

import "os"
import "fmt"

//noinspection GoTypesCompatibility
func main() {
	f, err := os.Create("example")
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil{
		return err
	}
}
