package libgotimestamp
import (
	"testing"
	"fmt"
)

func TestMakeTimestamp(t *testing.T) {
	what, where := MakeTimeStamp(23.976, 6963)
	fmt.Println(what, where)
	if what == nil {
		t.Errorf("WAT")
	}

}
