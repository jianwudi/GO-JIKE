package data

import (
	"account-service/internal/conf"

	"github.com/Shopify/sarama"
	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewKafka, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db        *gorm.DB
	conn      redis.Conn
	kafkaconn sarama.SyncProducer
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, conn redis.Conn, kafkaconn sarama.SyncProducer) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, conn: conn, kafkaconn: kafkaconn}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {

		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	return db
}

func NewRedis(c *conf.Data) redis.Conn {
	conn, err := redis.Dial("tcp", c.Redis.Addr)
	if err != nil {
		panic("failed to connect redis")
	}

	return conn
}

func NewKafka(c *conf.Data) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{c.Kafka.Addr}, config)
	if err != nil {
		panic("failed to connect kafka")
	}
	return client
}
