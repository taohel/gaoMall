package utils

import (
	"regexp"
)

// ValidatePassword 验证密码是否符合复杂度要求
func ValidatePassword(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	// 正则表达式：至少一个小写字母、一个大写字母、一个数字和一个特殊字符
	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,30}$`)
	return regex.MatchString(password)
}
