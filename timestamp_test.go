package libgotimestamp

import (
	"fmt"
	"testing"
)

func TestMakeTimestamp(t *testing.T) {
	what, where := MakeTimeStamp(23.976, 8991)
	fmt.Println(what, where)
	if what == nil {
		t.Errorf("WAT")
	}

}

func TestMakeFrame(t *testing.T) {
	where := MakeFrame(23.976, "0:08:20.80")
	fmt.Println(where)
	why := MakeFrame(23.976, "0:18:03.89")
	fmt.Println(why)
}
