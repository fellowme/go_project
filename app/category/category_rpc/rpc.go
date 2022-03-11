package category_rpc

import (
	"context"
	"go_project/app/category/category_dao"
	"go_project/rpc_service"
	"strconv"
	"strings"
)

type CategoryRpcService struct {
	dao category_dao.CategoryDaoInterface
}

func GetCategoryRpcService() *CategoryRpcService {
	return &CategoryRpcService{
		dao: category_dao.GetCategoryDao(),
	}
}

func (s CategoryRpcService) GetCategoryByIds(ctx context.Context, req *rpc_service.CategoryRequest) (*rpc_service.CategoryListResponse, error) {
	idList := make([]int, 0)
	categoryIdListString := strings.Split(req.CategoryIds, ",")
	for _, idString := range categoryIdListString {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		idList = append(idList, idInt)
	}
	categoryList := make([]*rpc_service.CategoryResponse, 0)
	data, err := s.dao.QueryCategoryByIdsDao(idList)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		categoryList = append(categoryList, &rpc_service.CategoryResponse{
			Id:               int32(item.Id),
			CategoryName:     item.CategoryName,
			CategorySort:     int32(item.CategorySort),
			CategoryParentId: int32(item.CategoryParentId),
		})
	}
	return &rpc_service.CategoryListResponse{
		CategoryList: categoryList,
	}, nil
}
