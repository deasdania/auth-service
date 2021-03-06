package repository

import "auth-service/api/models"

type IRoleMysql interface {
	GetRoleByName(name string) (*models.Role, error)
	GetRoleById(id string) (*models.Role, error)
	GetAllRole(orderby string) ([]*models.Role, error)
	CreateRole(name *models.Role) error
	UpdateRoleName(id string, name string) error
	DeleteRoleById(id string) error
	GetRoleByUserId(id string) (*models.Role, error)
	CheckUserIsAdmin(user_id string) (bool, error)
}
