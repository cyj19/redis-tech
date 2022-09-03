package lbs

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

/**
redis 实现查询附近的人功能
思路：使用redis的 地理空间（geospatial）数据结构存放用户位置
*/

const GEO_KEY = "user_geo"

var rdb *redis.Client

type point struct {
	Longitude float64 // 经度
	Latitude  float64 // 纬度
}

// 上传用户位置
func saveUserPoint(userId string, p *point) {
	err := rdb.GeoAdd(context.Background(), GEO_KEY, &redis.GeoLocation{
		Name:      userId,
		Longitude: p.Longitude,
		Latitude:  p.Latitude,
	}).Err()
	if err != nil {
		log.Println(err)
	}
}

// 查询附近的人
func near(p *point) ([]redis.GeoLocation, error) {
	return rdb.GeoRadius(context.Background(), GEO_KEY, p.Longitude, p.Latitude, &redis.GeoRadiusQuery{
		Radius:    200,  // 半径
		Count:     10,   // 查询数量
		WithCoord: true, // 返回结果带经纬度
		WithDist:  true, // 返回结果带距离
	}).Result()
}
