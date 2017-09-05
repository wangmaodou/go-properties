package property

import (
	"testing"
	"fmt"
	"./"
)

func (p Properties) test( t *testing.T) {
	p:=NewProperties("./property/test.properties")
	fmt.Println(p.String())
}
