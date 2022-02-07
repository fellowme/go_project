package product_model

import gin_model "github.com/fellowme/gin_common_library/model"

type ProductMain struct {
	gin_model.BaseMysqlStruct
	BrandId       int
	ShopId        int
	ShortTitle    string
	Title         string
	ImageId       int
	Description   string
	Weight        int
	ProductStatus int
}

type Product struct {
	gin_model.BaseMysqlStruct
	ProductMainId int
	ShortTitle    string
	Title         string
}
