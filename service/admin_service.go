package service

import (
	"ginDemo1/model"

	"github.com/go-xorm/xorm"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 *
 */
type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)

	//获取管理员总数
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine) AdminService {

	return &adminService{
		engine: db,
	}

}

type adminService struct {
	engine *xorm.Engine
}

/**
 * 通过用户名和密码查询管理员
 */
func (ad adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {

	var admin model.Admin
	ad.engine.Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.AdminId != 0
}

func (ad adminService) GetAdminCount() (int64, error) {
	count, err := ad.engine.Count(new(model.Admin))
	if err != nil {
		panic(err.Error())

	}
	return count, nil
}
