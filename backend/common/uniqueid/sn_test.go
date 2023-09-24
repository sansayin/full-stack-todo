package uniqueid

import (
	"fmt"
	"testing"
)

func TestGenSn(t *testing.T) {
	s := GenSn(SN_PREFIX_HOMESTAY_ORDER)
	fmt.Print(s)
}
