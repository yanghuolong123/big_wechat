package models

import (
	"github.com/astaxie/beego/orm"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"yhl/help"
)

func StatCountInfo(begin, end time.Time) int {
	i, err := orm.NewOrm().QueryTable("tbl_info").Filter("create_time__gte", begin).Filter("create_time__lt", end).Count()
	help.Error(err)

	return int(i)
}

func StatPv(begin, end time.Time) int {
	condition := bson.M{
		"time": bson.M{"$gte": begin, "$lt": end},
	}

	i, err := help.MongoDb.C("trace_record").Find(condition).Count()
	help.Error(err)

	return i
}

func StatUv(begin, end time.Time) int {
	condition := bson.M{
		"time": bson.M{"$gte": begin, "$lt": end},
	}

	job := &mgo.MapReduce{
		Map:    "function(){ emit(this.ip, 1) }",
		Reduce: "function(key, values) { return Array.sum(values) }",
	}

	type record struct {
		Ip    string "_id"
		Count int    "value"
	}
	var result []record
	_, err := help.MongoDb.C("trace_record").Find(condition).MapReduce(job, &result)
	help.Error(err)

	return len(result)
}
