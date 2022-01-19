package menu_service

import (
	"go_project/app/menu/menu_dao"
	"go_project/app/menu/menu_param"
)

type MenuServiceInterface interface {
	PostMenuService(param menu_param.PostMenuRequestParam) error
	GetMenuListService(param menu_param.GetMenuRequestParam) (menu_param.MenuListResponse, error)
	PatchMenuService(param menu_param.PatchMenuRequestParam) error
	DeleteMenuService(id int) error
}

type MenuService struct {
	dao menu_dao.MenuDaoInterface
}

func GetMenuService() MenuService {
	return MenuService{
		dao: menu_dao.GetMenuDao(),
	}
}

func (s MenuService) PostMenuService(param menu_param.PostMenuRequestParam) error {
	return s.dao.PostMenuByParamDao(param)
}

func (s MenuService) GetMenuListService(param menu_param.GetMenuRequestParam) (menu_param.MenuListResponse, error) {
	menuList, total, err := s.dao.GetMenuListByParamDao(param)
	return menu_param.MenuListResponse{
		Total: total,
		List:  menuList,
	}, err
}

func (s MenuService) PatchMenuService(param menu_param.PatchMenuRequestParam) error {
	return s.dao.PatchMenuByParamDao(param)
}

func (s MenuService) DeleteMenuService(id int) error {
	return s.dao.DeleteMenuByIdDao(id)

}
