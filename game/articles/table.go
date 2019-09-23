package articles

import (
	"fmt"
)

var Red string = "\033[31m"
var Green string = "\033[32m"
var Blue string = "\033[34m"
var Yellow string = "\033[33m"
var Reset string = "\033[0m"

func Table() bool {
	fmt.Println("Definite")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative d%ser%s      d%sie%s     das      d%sie%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Accusative d%sen%s      d%sie%s     das      d%sie%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Dative     d%sem%s      d%ser%s     dem      d%sen ...n%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Genitive   d%ses ...s%s d%ser%s     des ...s d%ser%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Println("")
	fmt.Println("Indefinite")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative ein        ein%se%s      ein        %s-%s\n", Blue, Reset, Yellow, Reset)
	fmt.Printf("Accusative ein%sen%s      ein%se%s      ein        %s-%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Dative     ein%sem%s      ein%ser%s     einem      %s...n%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Genitive   ein%ses ...s%s ein%ser%s     eines ...s %s-%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Println("")
	fmt.Println("Indefinite none")
	fmt.Println("           male     female  neutral  Plural")
	fmt.Printf("Nominative kein        kein%se%s      kein        kein%se%s\n", Blue, Reset, Yellow, Reset)
	fmt.Printf("Accusative kein%sen%s      kein%se%s      kein        kein%se%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Dative     kein%sem%s      kein%ser%s     keinem      kein%sen ...n%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	fmt.Printf("Genitive   kein%ses ...s%s kein%ser%s     keines ...s kein%ser%s\n", Green, Reset, Blue, Reset, Yellow, Reset)
	return true
}
