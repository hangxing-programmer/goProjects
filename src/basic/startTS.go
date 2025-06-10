package basic

import (
	"fmt"
	"time"
)

func test() {
	startTS := uint64(440709554502107146)
	physicalMs := startTS >> 18

	pdEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	unixMs := int64(physicalMs) + pdEpoch
	actualTime := time.UnixMilli(unixMs)

	fmt.Println("StartTS:", startTS)
	fmt.Println("物理时间（毫秒）:", physicalMs)
	fmt.Println("实际时间:", actualTime.UTC())
}
