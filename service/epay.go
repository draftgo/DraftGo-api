package service

import (
	"github.com/draftgo/DraftGo-api/setting/operation_setting"
	"github.com/draftgo/DraftGo-api/setting/system_setting"
)

func GetCallbackAddress() string {
	if operation_setting.CustomCallbackAddress == "" {
		return system_setting.ServerAddress
	}
	return operation_setting.CustomCallbackAddress
}
