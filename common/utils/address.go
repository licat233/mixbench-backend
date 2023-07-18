/*
 * @Author: licat
 * @Date: 2022-12-03 16:21:57
 * @LastEditors: licat
 * @LastEditTime: 2023-01-11 14:47:27
 * @Description: licat233@gmail.com
 */
package utils

import (
	"bytes"
	"log"
	"net"
	"path"

	"github.com/oschwald/geoip2-golang"
)

type AddressInfo struct {
	City      string  // 城市名称
	Provinces string  // 省份
	Country   string  // 国家名
	Continent string  // 洲名
	TimeZone  string  // 时区
	Latitude  float64 // 纬度
	Longitude float64 // 经度
}

func (a *AddressInfo) ToString() string {
	buf := bytes.Buffer{}
	buf.WriteString("Continent:" + a.Continent + "\n")
	buf.WriteString("Country:" + a.Country + "\n")
	buf.WriteString("Provinces:" + a.Provinces + "\n")
	buf.WriteString("City:" + a.City + "\n")
	return buf.String()
}

func QueryAddress(ipAddress string) *AddressInfo {
	dbfile := path.Join(ProjectRoot, "deploy/ip-db/GeoLite2-City.mmdb")
	db, err := geoip2.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipAddress)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	var info = new(AddressInfo)

	// 城市名称
	if len(record.City.Names) > 0 {
		info.City = record.City.Names["zh-CN"]
		//fmt.Println(record.City.Names["zh-CN"])
	}

	// 省份
	if len(record.Subdivisions) > 0 {
		if len(record.Subdivisions[0].Names) > 0 {
			info.Provinces = record.Subdivisions[0].Names["zh-CN"]
			//fmt.Println(record.Subdivisions[0].Names["zh-CN"])
		}
	}

	// 国家名
	if len(record.Country.Names) > 0 {
		info.Country = record.Country.Names["zh-CN"]
		//fmt.Println(record.Country.Names["zh-CN"])
	}

	// 洲名
	if len(record.Continent.Names) > 0 {
		info.Continent = record.Continent.Names["zh-CN"]
		//fmt.Println(record.Continent.Names["zh-CN"])
	}

	// 时区
	info.TimeZone = record.Location.TimeZone
	//fmt.Println(record.Location.TimeZone)

	// 纬度
	info.Latitude = record.Location.Latitude

	// 经度
	info.Longitude = record.Location.Longitude

	return info
}

func LocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	res := "0.0.0.0"
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = ipnet.IP.String()
			}
		}
	}
	return res, nil
}
