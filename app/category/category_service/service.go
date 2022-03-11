package category_service

import (
	"errors"
	"go_project/app/category/category_dao"
	"go_project/app/category/category_param"
)

type CategoryServiceInterface interface {
	GetCategoryListByParamService(param category_param.GetCategoryListRequestParam) (category_param.CategoryListResponse, error)
	CreateCategoryByParamService(param category_param.CategoryRequestParam) error
	UpdateCategoryByParamService(param category_param.CategoryRequestParam) error
	DeleteCategoryByIdService(id int) error
	RebuildCategoryService() error
}

type CategoryService struct {
	dao category_dao.CategoryDaoInterface
}

func GetCategoryService() CategoryService {
	return CategoryService{
		dao: category_dao.GetCategoryDao(),
	}
}

func (s CategoryService) GetCategoryListByParamService(param category_param.GetCategoryListRequestParam) (category_param.CategoryListResponse, error) {
	total, data, err := s.dao.GetCategoryListByParamDao(param)
	if err != nil {
		return category_param.CategoryListResponse{}, err
	}
	return category_param.CategoryListResponse{
		Total: total,
		List:  data,
	}, nil
}

func (s CategoryService) CreateCategoryByParamService(param category_param.CategoryRequestParam) error {
	return s.dao.CreateCategoryByParamDao(param)
}

func (s CategoryService) RebuildCategoryService() error {
	rootIdList := []int{0}
	dataList, err := s.dao.QueryCategoryListByParentId(rootIdList)
	if err != nil {
		return err
	}
	if len(dataList) == 0 {
		return errors.New("")
	}
	dataMap := make(map[int][]category_param.CategoryParam, 0)
	rootIdList = nil
	for _, item := range dataList {
		rootIdList = append(rootIdList, item.Id)
	}
	index := 1
	dataMap[index] = dataList
	for {
		categoryList, _ := s.dao.QueryCategoryListByParentId(rootIdList)
		if len(categoryList) == 0 {
			break
		}
		index++
		rootIdList = nil
		dataMap[index] = categoryList
		for _, item := range categoryList {
			rootIdList = append(rootIdList, item.Id)
		}
	}
	return nil
}

func (s CategoryService) UpdateCategoryByParamService(param category_param.CategoryRequestParam) error {
	return s.dao.UpdateCategoryByParamDao(param)
}

func (s CategoryService) DeleteCategoryByIdService(id int) error {
	return s.dao.DeleteCategoryByIdDao(id)
}
