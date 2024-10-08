package services

import (
	"Back-end/database"
	"Back-end/models"
	"Back-end/utils"
	"fmt"

	"gorm.io/gorm"
)

type ReturnUserInfo struct {
	UserID   string `json:"user_id"`          // 学号/工号
	Username string `json:"username"`         // 姓名，用户名(显示名称)
	Phone    string `json:"phone,omitempty"`  // 手机号，联系方式2（选填，预留接口）
	Mail     string `json:"mail"`             // 邮箱，联系方式一
	Avatar   string `json:"avatar,omitempty"` // 头像，大概得是Url了
}

// 获取用户信息
func GetProfileByID(userID string, table string) (ReturnUserInfo, error) {
	var result ReturnUserInfo
	err := database.DB.Table(table).Where("user_id = ?", userID).First(&result).Error
	if err != nil {
		utils.LogError(err)
		return result, err
	}
	return result, nil
}

// 更新用户信息
func UpdateProfile(userID, table, object, value string) error {
	if table != "students" && table != "admins" {
		return fmt.Errorf("invalid table name: %s", table)
	}

	result := database.DB.Table(table).Where("user_id = ?", userID).UpdateColumn(object, value)
	return result.Error
}

// 管理员的评分和已经接单的数量更新
func UpdateAdminHaddone(UndertakerID string, add int) error {
	var targetAdmin models.Admin
	result := database.DB.Model(&targetAdmin).Where("id = ?", UndertakerID).Update("had_done", gorm.Expr("had_done + ?", add))
	return result.Error
}

func UpdateAdminEvaluation(UndertakerID string, Grade int) error {
	var targetAdmin models.Admin
	result := database.DB.Model(&targetAdmin).Where("id = ?", UndertakerID).Update("evalutaion", gorm.Expr("evalutaion + ?", Grade))
	return result.Error
}
