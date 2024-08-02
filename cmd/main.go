package main

import (
	"fmt"

	"github.com/wolfedale/halyna/internal/mrad"
)

func main() {
	rad := mrad.New(450, 0, 0.1)
	fmt.Println(rad.Clicks(15), "clicks")
}
