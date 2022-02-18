package product_service

func getProductMainTypeNameByCode(code int) string {
	productMainTypeNameMap := map[int]string{
		1: "正常商品",
		2: "预售商品",
	}
	productMainTypeName, ok := productMainTypeNameMap[code]
	if ok {
		return productMainTypeName
	}
	return ""
}

func getProductMainStatusNameByCode(code int) string {
	productMainStatusNameMap := map[int]string{
		-1: "下线",
		1:  "上线",
	}
	productMainStatusName, ok := productMainStatusNameMap[code]
	if ok {
		return productMainStatusName
	}
	return ""
}
