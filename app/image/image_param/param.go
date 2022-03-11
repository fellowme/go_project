package image_param

import gin_param "github.com/fellowme/gin_common_library/param"

type (
	GetImageRequestParam struct {
		gin_param.PageRequestParam
		ImageId   int    `json:"image_id,omitempty" form:"image_id"`
		ImageName string `json:"image_name,omitempty" form:"image_name"`
		ImageType int    `json:"image_type,omitempty" form:"image_type"`
	}
)

type (
	ImageListResponse struct {
		Total int64           `json:"total,omitempty"`
		Data  []ImageResponse `json:"data,omitempty"`
	}
	ImageResponse struct {
		Id              int    `json:"id"`
		ImageUrl        string `json:"image_url,omitempty" `
		ImageName       string `json:"image_name,omitempty" `
		ImageUniqueName string `json:"image_unique_name,omitempty" `
		ImageSort       int    `json:"image_sort,omitempty"`
		ImageType       int    `json:"image_type,omitempty"`
		ImageHeight     string `json:"image_height,omitempty" `
		ImageWidth      string `json:"image_width,omitempty" `
		ImageSize       string `json:"image_size,omitempty" `
	}
)
