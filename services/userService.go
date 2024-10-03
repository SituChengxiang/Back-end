package services

import (
	"Back-end/database"
	"Back-end/utils"
)

type ReturnUserInfo struct {
	UserID   string `json:"user_id"`          // 学号/工号
	Username string `json:"username"`         // 姓名，用户名(显示名称)
	Phone    string `json:"phone,omitempty"`  // 手机号，联系方式2（选填，预留接口）
	Mail     string `json:"mail"`             // 邮箱，联系方式一
	Avatar   string `json:"avatar,omitempty"` // 头像，大概得是Url了
}

func GetProfileByID(userID string, table string) (ReturnUserInfo, error) {
	var result ReturnUserInfo
	err := database.DB.Table(table).Where("user_id = ?", userID).First(&result).Error
	if err != nil {
		utils.LogError(err)
		return result, err
	}
	return result, nil
}
