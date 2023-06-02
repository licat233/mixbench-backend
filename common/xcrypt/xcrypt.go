package xcrypt

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// GenerateFromPassword returns a byte slice so we need to
	// convert it to a string
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/** 验证密码强度
 * 长度 {8,64}
 * 大写字母 [A-Z]，必有
 * 小写字母 [a-z]，必有
 * 数字 \d，必有
 * 特殊符号 [~!@#$%^&*()_+-.|?]，可有可不有
 */
func PasswordValidator(password string) error {
	// 密码长度至少为8个字符，最多为64个字符
	if len(password) < 8 || len(password) > 64 {
		return errors.New("password length should be between 8 and 64 characters")
	}

	// 密码必须包含至少一个大写字母、一个小写字母和一个数字，以及一个特殊符号

	hasUpper := false
	hasLower := false
	hasDigit := false
	// hasSpecial := false
	// specialChars := "~!@#$%^&*()_+-.|?"
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
		// if strings.ContainsRune(specialChars, char) {
		// 	hasSpecial = true
		// }
	}
	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	// if !hasSpecial {
	// 	return errors.New("password must contain at least one special character")
	// }

	return nil
}
