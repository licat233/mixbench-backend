package xcrypt

import (
	"errors"
	"regexp"

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

func PasswordValidator(password string) error {
	// 密码长度至少为8个字符，最多为64个字符
	if len(password) < 8 || len(password) > 64 {
		return errors.New("password length should be between 8 and 64 characters")
	}

	// 密码必须包含至少一个大写字母、一个小写字母和一个数字
	re := regexp.MustCompile(`^[a-zA-Z0-9]*[a-z][a-zA-Z0-9]*[A-Z][a-zA-Z0-9]*\d[a-zA-Z0-9]*$`)
	if !re.MatchString(password) {
		return errors.New("the password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return nil
}
