package Marisfrolg_utils

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type VisitorLog struct {
	Id         bson.ObjectId `bson:"_id"`
	UserCode   string        `bson:"UserCode"`
	UserAgent string `bson:"UserAgent"`
	RequestUrl string        `bson:"RequestUrl"`
	Version    string        `bson:"Version"`
	Controller string        `bson:"Controller"`
	Func       string        `bson:"Func"`
	VisitDate  time.Time     `bson:"VisitDate"`
}

func (p *VisitorLog) AddPVInfo() (id string, err error) {
	session, err := mgo.Dial(MONGODB)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//这里要从PVData改成VisitorLog
	c := session.DB("ODSAPP").C("VisitorLog")
	p.Id = bson.NewObjectId()
	err = c.Insert(p)
	if err != nil {
		id = p.Id.Hex()
	}
	return
}
