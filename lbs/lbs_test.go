package lbs

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("执行TestMain")
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Println(err)
		return
	}
	m.Run()
}

func TestSave(t *testing.T) {
	saveUserPoint("1", &point{
		Longitude: 120.1111111,
		Latitude:  30.468746465,
	})

	saveUserPoint("2", &point{
		Longitude: 120.1111115,
		Latitude:  30.468746460,
	})

	saveUserPoint("3", &point{
		Longitude: 120.1111161,
		Latitude:  30.468746462,
	})

	saveUserPoint("4", &point{
		Longitude: 119.1111171,
		Latitude:  32.468746468,
	})

	saveUserPoint("5", &point{
		Longitude: 119.1112222,
		Latitude:  32.468745555,
	})
}

func TestNear(t *testing.T) {
	results, err := near(&point{
		Longitude: 120.1111111,
		Latitude:  30.468746460,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", results)
}

func TestNear2(t *testing.T) {
	results, err := near(&point{
		Longitude: 119.1111171,
		Latitude:  32.468746468,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", results)
}
