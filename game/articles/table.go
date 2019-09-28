package articles

import (
	"fmt"

	"github.com/padawin/german-practice/format"
)

func Table() bool {
	fmt.Println("Definite")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative d%ser%s      d%sie%s     das      d%sie%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Accusative d%sen%s      d%sie%s     das      d%sie%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Dative     d%sem%s      d%ser%s     dem      d%sen ...n%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Genitive   d%ses ...s%s d%ser%s     des ...s d%ser%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Println("")
	fmt.Println("Indefinite")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative ein        ein%se%s      ein        %s-%s\n", format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Accusative ein%sen%s      ein%se%s      ein        %s-%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Dative     ein%sem%s      ein%ser%s     einem      %s...n%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Genitive   ein%ses ...s%s ein%ser%s     eines ...s %s-%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Println("")
	fmt.Println("Indefinite none")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative kein        kein%se%s      kein        kein%se%s\n", format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Accusative kein%sen%s      kein%se%s      kein        kein%se%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Dative     kein%sem%s      kein%ser%s     keinem      kein%sen ...n%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	fmt.Printf("Genitive   kein%ses ...s%s kein%ser%s     keines ...s kein%ser%s\n", format.Green, format.Reset, format.Blue, format.Reset, format.Yellow, format.Reset)
	return true
}
