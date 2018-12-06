package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main() {
	if len(os.Args) <= 5 {
		fmt.Println("Usage: mongob <ip:port> <dbname> <collectionname> <[find|del]> <objectKey> [objectval]")
		return
	}

	mongoAddr := os.Args[1]
	mongodb := os.Args[2]
	mongocol := os.Args[3]
	action := os.Args[4]
	objectKey := os.Args[5]

	msession, err := mgo.Dial(mongoAddr)
	if err != nil {
		panic(err)
	}
	defer msession.Close()

	mcoll := msession.DB(mongodb).C(mongocol)
	switch action {
	case "find":
		result := make(map[string]interface{})
		iter := mcoll.Find(nil).Iter()
		for iter.Next(&result) {
			fmt.Println(result[objectKey].(string))
		}
		if iter.Err() != nil {
			fmt.Printf("[ERROR] error iter find mongo collection:%v, key:%v. err:%v\n",
				mongocol, objectKey, iter.Err())
		}
	case "del":
		if len(os.Args) <= 6 {
			fmt.Printf("[ERROR] invalid objectValue.\n")
			return
		}
		err = mcoll.Remove(bson.M{objectKey: os.Args[6]})
		if err != nil {
			fmt.Printf("[ERROR] error when remove mongo collection:%v, key:%v, value:%v. err:%v\n",
				mongocol, objectKey, os.Args[6], err)
		}
	default:

	}
}
