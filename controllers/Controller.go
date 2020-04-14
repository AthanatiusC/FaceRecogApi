package controllers

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zuramai/smartschool_api/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Sales struct {
	Region        string
	Country       string
	ItemType      string
	SalesChannel  string
	OrderPriority string
	OrderDate     string
	OrderID       string
	ShipDate      string
	UnitSold      string
	UnitPrice     string
	UnitCost      string
	TotalRevenue  string
	TotalCost     string
	TotalProfit   string
}

func ImportCsv(w http.ResponseWriter, r *http.Request) {
	// var sales Sales
	// var saless []Sales
	var saless []interface{}

	// models.GetDB("main").Collection("sales").InsertMany(context.TODO(), saless)
	timeNow := time.Now()
	csvFile, _ := os.Open("sales2.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		saless = append(saless, bson.M{
			"Region":        line[0],
			"Country":       line[1],
			"ItemType":      line[2],
			"SalesChannel":  line[3],
			"OrderPriority": line[4],
			"OrderDate":     line[5],
			"OrderID":       line[6],
			"ShipDate":      line[7],
			"UnitSold":      line[8],
			"UnitPrice":     line[9],
			"UnitCost":      line[10],
			"TotalRevenue":  line[11],
			"TotalCost":     line[12],
			"TotalProfit":   line[13],
		})
	}
	// cur, _ := models.GetDB("main").Collection("sales").Find(context.TODO(), bson.M{})

	// for cur.Next(context.TODO()) {
	// 	cur.Decode(&sales)
	// 	saless = append(saless, sales)
	// 	sales = Sales{}
	// }
	timeThen := time.Since(timeNow)
	fmt.Println("time elapsed:", timeThen)

	// salesJSON, _ := json.Marshal(sales)
	respondJSON(w, 200, "Success get all data sales", saless)
}

func respondJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	var payload models.Payload
	if status == 200 {
		payload.Status = true
	} else {
		payload.Status = false
	}
	payload.Message = message
	payload.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondErrorValidationJSON(w http.ResponseWriter, status int, message string, data map[string]interface{}) {
	var payload models.ErrorValidation
	payload.Message = "Error"
	payload.Errors = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func makeTimestampMilli() int64 {
	return unixMilli(time.Now())
}
