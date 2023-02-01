/*
 * @Author: licat
 * @Date: 2023-01-14 11:12:42
 * @LastEditors: licat
 * @LastEditTime: 2023-01-14 11:29:49
 * @Description: licat233@gmail.com
 */
package tools

func PluralizedName(name string) string {
	chip := name[len(name)-1:]
	switch chip {
	case "s":
		return name
	case "y":
		return name[:len(name)-1] + "ies"
	default:
		return name + "s"
	}
}
