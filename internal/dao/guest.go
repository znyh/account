package dao

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/account/internal/model"
	pb "github.com/znyh/proto/account"
)

func (d *dao) GuestRegister(ctx context.Context, u *pb.UserInfo, m map[string]interface{}) (err error) {
	if err = d.guestRegisterRedis(ctx, u, m); err != nil {
		return
	}
	if err = d.guestRegisterMysql(ctx, u); err != nil {
		return
	}
	if err = d.guestRegisterMongo(ctx, m); err != nil {
		return
	}
	return
}

func (d *dao) guestRegisterRedis(ctx context.Context, u *pb.UserInfo, m map[string]interface{}) (err error) {
	err = d.RedisHMSet(ctx, model.GetHAccountGuestKey(u.UserID), m)
	if err != nil {
		log.Error("redisGuestRegister error, bson:%+v", m)
		return
	}
	return
}

func (d *dao) guestRegisterMysql(ctx context.Context, u *pb.UserInfo) (err error) {
	query := fmt.Sprintf("INSERT account SET Nick=?,Password=?,Telephone=?,OS=?,Sex=?,UserID=?,HealUrl=?,City=?,Birth=?,Age=?,Online=?,CreateTime=?,LastLoginTime=?,IsRobot=?")
	log.Info("query:%+v v", query, )

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Error("insert prepare, error:%+v ", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(ctx, u.Nick, u.Password, u.Telephone, u.OS, u.Sex, u.UserID, u.HealUrl, u.City, u.Birth, u.Age, u.Online, u.CreateTime, u.LastLoginTime, u.IsRobot)
	if err != nil {
		log.Error("insert exec, error:%+v", err.Error())
		return
	}
	return
}

func (d *dao) guestRegisterMongo(ctx context.Context, m map[string]interface{}) (err error) {
	c := d.mdb.Database("db").Collection("coll")
	res, err := c.InsertOne(ctx, m)
	if err != nil {
		log.Error("MongoGuestRegister error, bson:%+v", m)
		return
	}
	log.Info("bson: %+v ret:%+v", m, res)
	return
}
