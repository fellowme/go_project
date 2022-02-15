package product_model

import (
	gin_model "github.com/fellowme/gin_common_library/model"
	"go_project/app/product/product_const"
)

type ProductMain struct {
	gin_model.BaseMysqlStruct
	BrandId           int                 `json:"brand_id,omitempty" gorm:"index:brand_id;comment:品牌id"`
	ShopId            int                 `json:"shop_id,omitempty" gorm:"index:shop_id;comment:专柜id"`
	ShortTitle        string              `json:"short_title,omitempty" gorm:"type:varchar(100);comment:短标题"`
	Title             string              `json:"title,omitempty" gorm:"type:varchar(200);comment:标题"`
	Weight            int                 `json:"weight,omitempty" gorm:"DEFAULT:100;comment:权重"`
	ProductMainStatus int                 `json:"product_main_status,omitempty" gorm:"DEFAULT:-1;comment:上下线状态 默认 下线"`
	CategoryId        int                 `json:"category_id,omitempty" gorm:"comment:类别id"`
	ProductCode       string              `json:"product_code,omitempty" gorm:"type:varchar(100);comment:商品码"`
	ProductMainType   int                 `json:"product_main_type,omitempty" gorm:"DEFAULT:1;comment:商品类型 1商品"`
	SaleTime          gin_model.LocalTime `json:"sale_time" gorm:"comment:开售时间"`
}

func (receiver ProductMain) TableName() string {
	return product_const.ProductMainTableName
}

type Product struct {
	gin_model.BaseMysqlStruct
	ProductMainId int    `json:"product_main_id,omitempty" gorm:"index:product_main_id;comment:product_main_id"`
	ShortTitle    string `json:"short_title,omitempty" gorm:"type:varchar(100);comment:短标题"`
	Title         string `json:"title,omitempty" gorm:"type:varchar(200);comment:标题"`
	ProductStatus int    `json:"product_status,omitempty" gorm:"DEFAULT:-1;comment:上下线状态 默认 下线"`
	IsMainProduct bool   `json:"is_main_product,omitempty" gorm:"DEFAULT:0;comment:是否是主商品 默认false"`
	Price         string `json:"price,omitempty" gorm:"comment:价格"`
	RealPrice     string `json:"real_price,omitempty" gorm:"comment:真实价格"`
	Description   string `json:"description,omitempty" gorm:"comment:描述"`
	Weight        int    `json:"weight,omitempty" gorm:"comment:权重"`
}

func (receiver Product) TableName() string {
	return product_const.ProductTableName
}

type Stock struct {
	gin_model.BaseMysqlStruct
	ProductId   int `json:"product_id,omitempty" gorm:"comment:商品id"`
	StockNumber int `json:"stock_number,omitempty" gorm:"comment:库存"`
}

func (receiver Stock) TableName() string {
	return product_const.StockTableName
}

type ProductImage struct {
	gin_model.BaseMysqlStruct
	ProductId        int `json:"product_id,omitempty" gorm:"comment:商品id "`
	ProductImageType int `json:"product_image_type" gorm:"comment:商品图片类型 1product_main 2 product"`
	ImageId          int `json:"image_id" gorm:"comment:图片id"`
}

func (receiver ProductImage) TableName() string {
	return product_const.ProductImageTableName
}
