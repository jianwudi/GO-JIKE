package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var rdb redis.UniversalClient

const (
	ip   string = "127.0.0.1"
	port uint16 = 6381
)

var ctx = context.Background()

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6381",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")

	write(1000, "len10", generateValue(10))
	write(500, "len20", generateValue(20))
	write(200, "len50", generateValue(50))
	write(100, "len100", generateValue(100))
	write(50, "len200", generateValue(200))
	write(10, "len1k", generateValue(1000))
	write(2, "len5k", generateValue(5000))
	analysis()
}

func write(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		err := rdb.Set(ctx, k, value, 100*time.Second).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}
