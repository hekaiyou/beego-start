package models

import (
	"log"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoSession 是 Mongo 数据库会话对象
var MongoSession *mgo.Session
// Database 是 Mongo 数据库名称
var Database string

func init() {
	// 获取配置信息中的 Mongo 数据库连接参数
	host, _ := config.String("MongoHost")
	port, _ := config.String("MongoPort")
	Database, _ = config.String("MongoDatabase")
	var servers strings.Builder
	servers.WriteString(host)
	servers.WriteString(":")
	servers.WriteString(port)
	diaInfo := &mgo.DialInfo{
		Addrs:     []string{servers.String()}, // 数据库地址
		Timeout:   60 * time.Second,           // 连接超时时间
		PoolLimit: 4096,                       // 连接池的数量
	}
	var err error
	MongoSession, err = mgo.DialWithInfo(diaInfo)
	if err != nil {
		log.Fatalf("create mongoDB session error: %v (%v)", err, servers.String())
	}
}

// 连接到具体的Mongo集合，返回MongoDB会话对象和Mongo集合对象
func getConnect(collection string) (*mgo.Session, *mgo.Collection) {
	ms := MongoSession.Copy()               // 复制MongoDB会话对象，避免重复创建MongoDB会话，导致连接数超过最大值
	c := ms.DB(Database).C(collection) // 获取集合对象
	ms.SetMode(mgo.Monotonic, true)    // 一致性模式：Monotonic（单调一致性）
	return ms, c
}

// Entity 公共的Mongo文档存储结构字段，包含ID、创建时间和更新时间
type Entity struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	InsertedAt time.Time     `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time     `json:"last_update" bson:"last_update"`
}

/*
在Mongo集合中插入文档
	collection 操作的文档
	doc        要插入的数据
模型设计
	U.Entity = Entity{bson.NewObjectId(), time.Now().UTC(), time.Now().UTC()}
	err := insertRow(userDatabase, U)
	fmt.Println(err)
*/
func insertRow(collection string, doc interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	return c.Insert(doc)
}

/*
在Mongo集合中删除单个匹配文档
	collection 操作的文档
	selector   删除条件
模型设计
	err := removeRow(userDatabase, bson.M{"_id": bson.ObjectIdHex("615431c35bf5245a9c7c7f80")})
	fmt.Println(err)
*/
func removeRow(collection string, selector interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	return c.Remove(selector)
}

/*
在Mongo集合中删除全部匹配文档
	collection 操作的文档
	selector   删除条件
模型设计
	err := removeAllRow(userDatabase, bson.M{})
	fmt.Println(err)
*/
func removeAllRow(collection string, selector interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	_, err := c.RemoveAll(selector)
	return err
}

/*
在Mongo集合中查询单个匹配文档
	collection 操作的文档
	query      查询条件（筛选文档）
	selector   需要过滤的数据（筛选字段）
	result     查询到的结果
模型设计
	var result User
	err := findRow(userDatabase, bson.M{"name": "Name"}, bson.M{}, &result)
	fmt.Println(err)
	fmt.Println(result)
*/
func findRow(collection string, query, selector, result interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	return c.Find(query).Select(selector).One(result)
}

/*
在Mongo集合中查询全部匹配文档
	collection 操作的文档
	query      查询条件（筛选文档）
	selector   需要过滤的数据（筛选字段）
	result     查询到的结果
模型设计
	var result []*User
	err := findAllRow(userDatabase, bson.M{}, bson.M{"_id": 0}, &result)
	fmt.Println(err)
	fmt.Println(result)
*/
func findAllRow(collection string, query, selector, result interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	return c.Find(query).Select(selector).All(result)
}

/*
在Mongo集合中分页查询全部匹配文档
	collection 操作的文档
	offset     当前页面（从0开始）
	limit      每页的数量值（从1开始）
	query      查询条件（筛选文档）
	selector   需要过滤的数据（筛选字段）
	result     查询到的结果
模型设计
	var result []*User
	err := findPage(userDatabase, 0, 10, bson.M{}, bson.M{}, &result)
	fmt.Println(err)
	fmt.Println(result)
*/
func findPage(collection string, offset, limit int, query, selector, result interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Skip(offset * limit).Limit(limit).All(result)
}

/*
在Mongo集合中更新单个匹配文档
	collection 操作的文档
	selector   更新条件
	update     更新的操作
模型设计
	err := updateRow(userDatabase, bson.M{"_id": bson.ObjectIdHex("615431c35bf5245a9c7c7f80")}, bson.M{"$set": bson.M{"name": "Name"}})
	fmt.Println(err)
*/
func updateRow(collection string, selector, update interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	return c.Update(selector, update)
}

/*
在Mongo集合中更新全部匹配文档
	collection 操作的文档
	selector   更新条件
	update     更新的操作
模型设计
	err := updateAllRow(userDatabase, bson.M{"name": "Name"}, bson.M{"$set": bson.M{"name": "批量更新"}})
	fmt.Println(err)
*/
func updateAllRow(collection string, selector, update interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	_, err := c.UpdateAll(selector, update)
	return err
}

/*
在Mongo集合中更新匹配文档，如果不存在，则插入一个新文档
	collection 操作的文档
	selector   更新条件
	update     更新的操作
模型设计
	U.Entity = Entity{bson.NewObjectId(), time.Now().UTC(), time.Now().UTC()}
	U.Name = "更新或插入"
	U.Email = "Email"
	err := upsertRow(userDatabase, bson.M{"name": "Name"}, bson.M{"$set": U})
	fmt.Println(err)
*/
func upsertRow(collection string, selector, update interface{}) error {
	ms, c := getConnect(collection)
	defer ms.Close() // 每次操作后都主动关闭
	_, err := c.Upsert(selector, update)
	return err
}

/*
判断Mongo集合中的文档数量是否为零
	collection 操作的文档
模型设计
	isEmpty := isEmptyCollection(userDatabase)
	fmt.Println(isEmpty)
*/
func isEmptyCollection(collection string) bool {
	ms, c := getConnect(collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

/*
获取Mongo集合中的匹配文档数量
	collection 操作的文档
	query      查询条件（筛选文档）
模型设计
	count, err := rowCount(userDatabase, bson.M{})
	fmt.Println(err)
	fmt.Println(count)
*/
func rowCount(collection string, query interface{}) (int, error) {
	ms, c := getConnect(collection)
	defer ms.Close()
	return c.Find(query).Count()
}
