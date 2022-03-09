package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/garyburd/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	conn redis.Conn
)

func init() {
	db = NewDB()
	conn = NewRedis()
}

type BizUser struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
}
type KUser struct {
	NeedDB int32
	User
}

func CreateUser(u *User) error {

	rv := db.Create(u)
	return rv.Error

}

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:dangerous@tcp(127.0.0.1:3307)/realworld?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {

		panic("failed to connect database")
	}

	return db
}

func NewRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic("failed to connect redis")
	}

	return conn
}

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("account") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("account", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		for msg := range pc.Messages() {
			//存到数据库
			var kuser KUser
			//		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			err := json.Unmarshal(msg.Value, &kuser)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%+v\n", kuser)
			//先存入数据库
			if kuser.NeedDB == 1 {
				err = CreateUser(&kuser.User)
				if err != nil {
					panic(err)
				}
			} else {
				//从db获取
				rv := db.Where("Email = ?", kuser.Email).First(&kuser.User)
				if rv.Error != nil {
					panic("sss")
				}
			}
			//存入redis
			fmt.Println(kuser.User.Email)
			datas, _ := json.Marshal(&kuser.User)
			_, err = conn.Do("set", kuser.User.Email, datas)
			if err != nil {
				panic("sss")
			}
			/* 			rebytes, err := redis.Bytes(conn.Do("get", "123"))
			   			if err != nil {
			   				panic("sss1")
			   			}
			   			object := &User{}
			   			json.Unmarshal(rebytes, object)
			   			fmt.Printf("%+v\n", object) */
		}
	}
}
