// The MIT License (MIT)
//
// Copyright (c) 2017 AT&T
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html"
	"net/http"
	"strings"
)

var portNumber = flag.Int("port", 8085, "Port number listening on")
var databaseIP = flag.String("database-ip", "127.0.0.1", "Ip of MongoDB: ")

type Mark struct {
	Id      string          `json:"Id" bson:"_id,omitempty"`
	Label   string          `json:"Label"`
	Type    string          `json:"Type"`
	Content json.RawMessage `json:"Content"`
}

func resetHandler(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	session, err := mgo.Dial(*databaseIP)
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer session.Close()

	session.DB("dali").C("marks").RemoveAll(nil)
	writeResponse(responseWriter, http.StatusAccepted, "done")
}

func pingHandler(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	writeResponse(responseWriter, http.StatusAccepted, "pong")
}

func getIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	objectId := bson.NewObjectId()
	writeResponse(responseWriter, http.StatusAccepted, objectId.Hex())
}

func markHandler(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	session, err := mgo.Dial(*databaseIP)
	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer session.Close()

	c := session.DB("dali").C("marks")

	path := html.EscapeString(request.URL.Path)

	var results []Mark

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	if request.Method == "POST" {
		markStruct := &Mark{}

		if error := json.NewDecoder(request.Body).Decode(&markStruct); error != nil {
			writeResponse(responseWriter, http.StatusBadRequest, error.Error())
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")

		currentId := markStruct.Id
		currentLabel := markStruct.Label
		currentType := markStruct.Type
		currentContent := markStruct.Content

		err = c.Insert(&Mark{Id: currentId, Label: currentLabel, Type: currentType, Content: currentContent})
		if err != nil {
			writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		} else {
			writeResponse(responseWriter, http.StatusAccepted, currentId)
		}
	} else if request.Method == "GET" {
		splitPath := strings.Split(path, "/")
		if (len(splitPath) == 2) || (splitPath[2] == "") {
			err = c.Find(bson.M{}).All(&results)
			if err != nil {
				writeResponse(responseWriter, http.StatusBadRequest, err.Error())
			} else {
				returnList(responseWriter, http.StatusAccepted, results)
			}
		} else if (len(splitPath) == 3) && (splitPath[2] != "") {
			splitPath := strings.Split(path, "/mark/")
			id := splitPath[1]
			result := Mark{}
			err = c.Find(bson.M{"_id": id}).One(&result)

			if err != nil {
				writeResponse(responseWriter, http.StatusBadRequest, err.Error())
			} else {
				returnMark(responseWriter, http.StatusAccepted, result)
			}
		}
	} else if request.Method == "PUT" {
		fmt.Println("PUT")
		markStruct := &Mark{}
		fmt.Println(markStruct)

		if error := json.NewDecoder(request.Body).Decode(&markStruct); error != nil {
			writeResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		markId := markStruct.Id
		markLabel := markStruct.Label
		markType := markStruct.Type
		markContent := markStruct.Content
		fmt.Println(markType)

		idOfEntry := bson.M{"_id": markId}

		change := bson.M{"$set": bson.M{"_id": markId, "label": markLabel, "type": markType, "content": markContent}}
		err = c.Update(idOfEntry, change)
		if err != nil {
			writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		} else {
			writeResponse(responseWriter, http.StatusAccepted, "done")
		}
	} else if request.Method == "DELETE" {
		splitPath := strings.Split(path, "/mark/")
		id := splitPath[1]

		err = c.Remove(bson.M{"_id": id})
		if err != nil {
			writeResponse(responseWriter, http.StatusBadRequest, err.Error())
		} else {
			writeResponse(responseWriter, http.StatusAccepted, "done")
		}
	}
}

func returnList(responseWriter http.ResponseWriter, code int, marks []Mark) {
	responseWriter.WriteHeader(code)
	data := struct {
		Marks []Mark `json:"marks"`
	}{
		marks,
	}
	json.NewEncoder(responseWriter).Encode(&data)
}

func returnMark(responseWriter http.ResponseWriter, code int, requestedMark Mark) {
	responseWriter.WriteHeader(code)
	json.NewEncoder(responseWriter).Encode(&requestedMark)
}

func writeResponse(responseWriter http.ResponseWriter, code int, message string) {
	responseWriter.WriteHeader(code)
	data := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		code,
		message,
	}
	json.NewEncoder(responseWriter).Encode(&data)
}

func main() {
	flag.Parse()
	listenOn := fmt.Sprintf(":%v", *portNumber)

	http.HandleFunc("/id", getIdHandler)
	http.HandleFunc("/mark/", markHandler)
	http.HandleFunc("/mark", markHandler)
	http.HandleFunc("/admin/reset", resetHandler)
	http.HandleFunc("/admin/ping", pingHandler)

	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./templates"))))

	http.ListenAndServe(listenOn, nil)
}
