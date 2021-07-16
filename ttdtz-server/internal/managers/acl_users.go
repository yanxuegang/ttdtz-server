package managers

import (
	"ttdtz-server/global"
	. "ttdtz-server/internal/models"

	"github.com/jinzhu/gorm"
)

func GetAclUserById(userId uint) (*AclUsers, error) {
	//todo errcode
	if userId == 0 {
		return nil, nil
	}
	var (
		AclUser = new(AclUsers)
		err     error
	)
	AclUser.Id = userId
	err = getAclUserFromDB(AclUser)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, nil
	}
	return AclUser, nil
}

func getAclUserFromDB(acluser *AclUsers) error {
	return global.GetDB("app_line").First(&acluser).Error
}
