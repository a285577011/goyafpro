package business

import (
	"strconv"
)

type Base struct {
}

//检测page和count参数，如果没有值赋予默认值1和10
//因为page和count参数都是在列表的时候才需要传递，所以这里直接计算出offset
func (this *Base) CheckPageAndCount(params *map[string]string) {
	if _, ok := (*params)["page"]; !ok || len((*params)["page"]) == 0 {
		(*params)["page"] = "1"
	}
	if _, ok := (*params)["count"]; !ok || len((*params)["count"]) == 0 {
		(*params)["count"] = "10"
	}

	page, _ := strconv.Atoi((*params)["page"])
	count, _ := strconv.Atoi((*params)["count"])
	offset := (page - 1) * count
	if offset < 0 {
		offset = 0
	}
	(*params)["offset"] = strconv.Itoa(offset)
}
