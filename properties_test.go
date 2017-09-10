package prop

import (
	"testing"
	"fmt"
)

func Test( t *testing.T) {
	p:=NewProperties("./test.properties")
	fmt.Println(p.String())
	if p.GetString("name")!="maodou"{
		t.Error("The value is wrong!")
	}
}
