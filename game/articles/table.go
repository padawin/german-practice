package articles

import (
	"fmt"

	"github.com/padawin/german-practice/format"
)

var colors = []string{format.Green, format.Blue, format.Orange, format.Yellow}

func Table() bool {
	for _, response := range responses {
		fmt.Printf("%s%s%s\n", format.Red, response.Name, format.Reset)
		fmt.Printf("%10v", "")
		for _, gender := range genders {
			fmt.Printf("%-13v", gender)
		}
		fmt.Println()
		for caseIndex, caseName := range cases {
			fmt.Printf("%-10v", caseName[1])
			for genderIndex, _ := range genders {
				val := response.getFormatted(caseIndex, genderIndex, -13, colors[genderIndex])
				fmt.Print(val)
			}
			fmt.Println()
		}
		fmt.Println()
	}
	return true
}
