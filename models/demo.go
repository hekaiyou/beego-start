package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

var databaseName = "Demo"

type Demo struct {
	Entity     `bson:",inline"`
	Score      int64  `json:"score" bson:"score"`
	PlayerName string `json:"player_name" bson:"player_name"`
}

type DemoDocuEdit struct {
	Score      int64  `json:"score"`
	PlayerName string `json:"player_name"`
}

type DemoDocuUpdate struct {
	Score int64 `json:"score"`
}

func AddDemo(d Demo) (string, error) {
	d.Entity = Entity{bson.NewObjectId(), time.Now().UTC(), time.Now().UTC()}
	err := insertRow(databaseName, d)
	return d.ID.Hex(), err
}

func GetDemo(id string) (*Demo, error) {
	var result *Demo
	err := findRow(databaseName, bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{}, &result)
	return result, err
}

func GetAllDemo() ([]*Demo, error) {
	var result []*Demo
	err := findAllRow(databaseName, bson.M{}, bson.M{}, &result)
	return result, err
}

func UpdateDemo(id string, du DemoDocuUpdate) error {
	err := updateRow(databaseName, bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": du})
	return err
}

func DeleteDemo(id string) error {
	err := removeRow(databaseName, bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
