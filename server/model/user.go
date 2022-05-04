/**
 * @Author: Yuxue.Fan<yuxue.fan@yidatec.com>
 * @Description:
 * @File: user.go
 * @Version: 1.0.0
 * @Date: 2022/2/18 10:34
 */

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2Ffc%2F64%2Fab%2Ffc64ab13576aeb6522356f19c4426aee.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1647743817&t=1c9ace914b246cc1f17bbd043a54f5bf;comment:用户头像"`
	//Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	UserProfile string 		 `json:"user_profile" gorm:"default:''; comment:个人描述"`
}

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	//DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	//Children        []SysAuthority `json:"children" gorm:"-"`
	//SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
