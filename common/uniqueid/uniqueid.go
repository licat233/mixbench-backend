/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 17:25:49
 * @Description: licat233@gmail.com
 */
package uniqueid

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

func NewUUID() string {
	return uuid.New().String()
}
