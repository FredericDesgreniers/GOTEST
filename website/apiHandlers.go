package website

import (
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/qiniu/log"
)

//handles the api landing page
func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "THIS IS THE API")
}

//Displays a list of users in json format
func userHandler(w http.ResponseWriter, r *http.Request){

	c := db.Session.DB("codejam").C("users")

	var results []interface{}
	var err error

	err = c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}
	jsonText, err := json.Marshal(results)
	fmt.Fprint(w,string(jsonText ))
}

