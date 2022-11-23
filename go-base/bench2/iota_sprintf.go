package bench2

import (
	"fmt"
	"strconv"
)

func UseStringsItoa() {
	strconv.Itoa(12345)
}

func UseSprintf() {
	fmt.Sprintf("%d", 12345)
}
