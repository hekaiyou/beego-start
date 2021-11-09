package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

var collectionName = "Demo"

// Demo 某对象类型
type Demo struct {
	Entity     `bson:",inline"`
	Score      int64  `json:"score" bson:"score"`
	PlayerName string `json:"player_name" bson:"player_name"`
}

// DemoEditRequest 某对象的编辑请求类型
type DemoEditRequest struct {
	Score      int64  `json:"score"`
	PlayerName string `json:"player_name"`
}

type DemoDocuUpdate struct {
	Score int64 `json:"score"`
}

// AddDemo 创建某对象的文档
func AddDemo(d Demo) (string, error) {
	d.Entity = Entity{bson.NewObjectId(), time.Now().UTC(), time.Now().UTC()}
	err := insertRow(collectionName, d)
	return d.ID.Hex(), err
}

func GetDemo(id string) (*Demo, error) {
	var result *Demo
	err := findRow(collectionName, bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{}, &result)
	return result, err
}

func GetAllDemo() ([]*Demo, error) {
	var result []*Demo
	err := findAllRow(collectionName, bson.M{}, bson.M{}, &result)
	return result, err
}

func UpdateDemo(id string, du DemoDocuUpdate) error {
	err := updateRow(collectionName, bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": du})
	return err
}

func DeleteDemo(id string) error {
	err := removeRow(collectionName, bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
