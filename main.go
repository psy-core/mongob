package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main() {
	if len(os.Args) <= 5 {
		fmt.Println("Usage: mongodel <ip:port> <dbname> <collectionname> <objectKey> <objectval>")
		return
	}

	mongoAddr := os.Args[1]
	mongodb := os.Args[2]
	mongocol := os.Args[3]
	objectKey := os.Args[4]
	objectValue := os.Args[5]

	msession, err := mgo.Dial(mongoAddr)
	if err != nil {
		panic(err)
	}
	defer msession.Close()

	mcoll := msession.DB(mongodb).C(mongocol)
	err = mcoll.Remove(bson.M{objectKey: objectValue})
	if err != nil {
		fmt.Printf("[ERROR] error when remove mongo collection:%v, key:%v, value:%v. err:%v\n",
			mongocol, objectKey, objectValue, err)
	}
}
