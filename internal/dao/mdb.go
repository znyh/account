package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	mg "github.com/znyh/middle-end/library/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (

	//database name
	AccountDatabase = "account"

	//Collection name
	GuestAccount    = "GuestAccount"
	FacebookAccount = "FacebookAccount"
	LineAccount     = "LineAccount"
	OPPOAccount     = "OPPOAccount"
	Binding         = "Binding"
	Blank           = "AccountBlank"
	Freeze          = "AccountFreeze"
	Silent          = "AccountSilent"
	RegInfo         = "RegRecord"
)

var (
	CollectionMap map[string]*mongo.Collection
)

func NewMongoDB() (db *mg.DB, err error) {
	var cfg struct {
		Client *mg.Config
	}

	if err = paladin.Get("mdb.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}

	db = mg.NewMongoDB(cfg.Client)
	InitCollectionMap(db)
	return
}

func InitCollectionMap(db *mg.DB) {
	CollectionMap = make(map[string]*mongo.Collection)
	database := db.Database(AccountDatabase)
	CollectionMap[GuestAccount] = database.Collection(GuestAccount)
	CollectionMap[FacebookAccount] = database.Collection(FacebookAccount)
	CollectionMap[LineAccount] = database.Collection(LineAccount)
	CollectionMap[OPPOAccount] = database.Collection(OPPOAccount)
	CollectionMap[Binding] = database.Collection(Binding)
	CollectionMap[Blank] = database.Collection(Blank)
	CollectionMap[Freeze] = database.Collection(Freeze)
	CollectionMap[Silent] = database.Collection(Silent)
	CollectionMap[RegInfo] = database.Collection(RegInfo)
}

func (d *dao) MongoUpdate(ctx context.Context, table string, filter, update interface{}) (err error) {

	if v, ok := CollectionMap[table]; ok {
		opts := new(options.UpdateOptions)
		opts.SetUpsert(true)
		_, err = d.mdb.UpdateOne(ctx, v, filter, update, opts)
	}
	return
}
