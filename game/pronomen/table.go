package pronomen

import (
	"fmt"

	"github.com/padawin/german-practice/format"
)

func Table() bool {
	fmt.Printf("%30v%s", "", format.Yellow)
	for _, caseName := range cases {
		fmt.Printf("%-13v", caseName)
	}
	fmt.Println(format.Reset)
	for personIndex, person := range pronomina {
		fmt.Printf("%s%-30v%s", format.Red, persons[personIndex], format.Reset)
		for _, pronomen := range person {
			fmt.Printf("%-13v", pronomen)
		}
		fmt.Println()
	}
	fmt.Println()
	return true
}
