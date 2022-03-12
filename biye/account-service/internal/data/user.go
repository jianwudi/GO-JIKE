package data

import (
	"account-service/internal/biz"
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
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

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *userRepo) sendToKafKa(k *KUser) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "account"
	msg.Key = sarama.StringEncoder(k.Email)
	datas, _ := json.Marshal(k)
	msg.Value = sarama.StringEncoder(datas)
	//发往kafka
	pid, offset, err := r.data.kafkaconn.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}
	kuser := KUser{}
	kuser.User = user
	kuser.NeedDB = 1
	err := r.sendToKafKa(&kuser)
	return err
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	//先去访问缓存
	//email作为key
	u := User{}
	rebytes, err := redis.Bytes(r.data.conn.Do("get", email))
	if err != nil {
		//访问db
		rv := r.data.db.Where("Email = ?", "123").First(&u)
		if rv.Error != nil {
			panic("sss")
		}
		kuser := KUser{}
		kuser.User = u
		kuser.NeedDB = 0
		r.sendToKafKa(&kuser)
	} else {
		json.Unmarshal(rebytes, &u)
	}

	user := &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}
	fmt.Printf("%+v \n", user)
	return user, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
