package models

import (
	"github.com/astaxie/beego/orm"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"yhl/help"
)

func StatCountInfo(begin, end time.Time) int {
	//i, err := orm.NewOrm().QueryTable("tbl_info").Filter("create_time__gte", begin).Filter("create_time__lt", end).Count()
	i, err := orm.NewOrm().QueryTable("tbl_info").Filter("create_time__between", begin, end).Count()
	help.Error(err)

	return int(i)
}

func StatInfoViews(begin, end time.Time) int {
	var sum int
	err := orm.NewOrm().Raw("select sum(views) from tbl_info where create_time between ? and ?", begin, end).QueryRow(&sum)
	help.Error(err)

	return sum
}

func StatPv(begin, end time.Time) int {
	condition := bson.M{
		"time": bson.M{"$gte": begin, "$lt": end},
		"uri":  bson.M{"$ne": "/login"},
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
