package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-file/common"
	"strconv"
	"strings"
)

type Option struct {
	Key   string `json:"key" gorm:"primaryKey;type:string"`
	Value string `json:"value" gorm:"type:string;"`
}

func AllOption() ([]*Option, error) {
	var options []*Option
	var err error
	err = DB.Find(&options).Error
	return options, err
}

func InitOptionMap() {
	common.OptionMap = make(map[string]interface{})
	common.OptionMap["FileUploadPermission"] = common.FileUploadPermission
	common.OptionMap["FileDownloadPermission"] = common.FileDownloadPermission
	common.OptionMap["ImageUploadPermission"] = common.ImageUploadPermission
	common.OptionMap["ImageDownloadPermission"] = common.ImageDownloadPermission
	options, _ := AllOption()
	for _, option := range options {
		updateOptionMap(option.Key, option.Value)
	}
}

func UpdateOption(key string, value string) {
	// Save to database first
	option := Option{
		Key:   key,
		Value: value,
	}
	if DB.Model(&option).Where("key = ?", key).Updates(&option).RowsAffected == 0 {
		DB.Create(&option)
	}
	// Update OptionMap
	updateOptionMap(key, value)
}

func updateOptionMap(key string, value string) {
	if strings.HasSuffix(key, "Permission") {
		common.OptionMap[key], _ = strconv.Atoi(value)
	} else {
		common.OptionMap[key] = value
	}
}
