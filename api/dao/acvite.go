/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-21 23:06:49
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-21 23:21:46
 * @FilePath: /voting-ranking/api/dao/acvite.go
 * @Description:
 *
 */
package dao

import (
	"time"
	"voting-ranking/api/dto"
	"voting-ranking/api/model"
	"voting-ranking/common/util"
	"voting-ranking/pkg/db"
)

func AddActivity(dto dto.ActivityDto) error {
	activity := model.Activity{
		Name:       dto.Name,
		CreateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&activity).Error
	return err
}

func GetActivityByName(name string) (activity model.Activity) {
	db.Db.Where("name = ?", name).First(&activity)
	return activity
}
