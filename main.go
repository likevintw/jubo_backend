package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ResponseFormat struct {
	Message string `json:"message"`
}

type patient struct {
	Id      int    "bson:`id`"
	Name    string "bson:`name`"
	OrderId string "bson:`ordrId`"
	Gender  string "bson:`gender`"
	Illness string "bson:`illness`"
	History string "bson:`history`"
	Dialog  string "bson:`dialog`"
}

func CreateEchoHandler() (*EchoHandler, error) {
	e := echo.New()
	e.Use(middleware.CORS())
	var mux sync.Mutex
	account := fmt.Sprintf("%s", os.Getenv("MONGO_INITDB_ROOT_USERNAME"))
	password := fmt.Sprintf("%s", os.Getenv("MONGO_INITDB_ROOT_PASSWORD"))
	url := fmt.Sprintf("mongodb://%s:%s@localhost", account, password)
	fmt.Println(url)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	handler := EchoHandler{e, mux, client}
	return &handler, nil
}

type EchoHandler struct {
	E      *echo.Echo
	Mux    sync.Mutex
	client *mongo.Client
}

func (h EchoHandler) RunHTTPServer() {
	h.E.GET("/returnPatients", h.returnPatients)
	h.E.GET("/updateDialog", h.updateDialog)
	h.E.Logger.Fatal(h.E.Start(":80"))
}
func (h *EchoHandler) returnPatients(c echo.Context) error {
	fmt.Println("link to returnPatients()")
	var err error
	collection := h.client.Database("db").Collection("patients")
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []patient
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("error collection.Find", err)
		return c.JSON(400, err)
	}
	for cur.Next(context.TODO()) {
		var elem patient
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("error cur.Decode", err)
			return c.JSON(400, err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("error cur.Err()", err)
		return c.JSON(400, err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents: %+v\n", results)
	return c.JSON(http.StatusOK, results)

}
func (h *EchoHandler) updateDialog(c echo.Context) error {
	var err error
	json_map := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(json_map["id"])
	fmt.Println(json_map["dialog"])

	response := map[string]string{"message": "ok"}

	return c.JSON(http.StatusOK, response)

}

func main() {
	h, err := CreateEchoHandler()
	if err != nil {
		panic(err)
	}
	h.RunHTTPServer()
}
