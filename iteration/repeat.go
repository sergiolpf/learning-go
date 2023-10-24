package iteration

import (
	"fmt"
	"math/rand"
)

func Repeat(character string, quantity int) string {
	var repeated string
	for i := 0; i < quantity; i++ {
		repeated += character
	}
	return repeated

}

func main() {
	fmt.Println(Repeat("ha", rand.Intn(345)))
}
