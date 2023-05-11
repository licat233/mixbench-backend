/*
 * @Author: licat
 * @Date: 2023-02-18 14:18:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:18:07
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"errors"
	"mixbench/service/v1/admin/model"
)

func verifyAdminerStatus(adminer *model.Adminer) error {
	if adminer.Status == 0 {
		return errors.New("this account has been disabled")
	}
	return nil
}
