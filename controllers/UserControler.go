package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"gonum.org/v1/gonum/floats"

	"github.com/AthanatiusC/FaceRecogApi/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Recognize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var userEmbedding models.UserEmbeddings
	var userEmbeddings []models.UserEmbeddings
	var recognition models.UserRecognition
	var recognitionList []models.UserRecognition
	var log models.Log

	json.NewDecoder(r.Body).Decode(&recognition)

	if len(recognition.Embedding) == 0 {
		respondErrorValidationJSON(w, 422, "Input Embedding Null", map[string]interface{}{})
		return
	}

	cursor, err := models.GetDB("main").Collection("users").Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&userEmbedding)
		userEmbeddings = append(userEmbeddings, userEmbedding)
		userEmbedding = models.UserEmbeddings{}
	}

	for _, UserEmbeddingList := range userEmbeddings {
		// Index := index
		if len(UserEmbeddingList.Embeddings) == 0 {
			continue
		} else {
			var val []float64
			for _, embeddingList := range UserEmbeddingList.Embeddings {
				val = append(val, euclideanDistance(embeddingList, recognition.Embedding))
			}
			recognition.UserID = UserEmbeddingList.UserID
			recognition.Name = UserEmbeddingList.Name
			recognition.Accuracy = floats.Min(val)
			recognition.Elapsed = time.Since(start).String()
			recognitionList = append(recognitionList, recognition)
		}
		// log.Println(maximum)
	}
	if len(recognitionList) == 0 {
		respondErrorValidationJSON(w, 422, "Cannot Recognize Face!", map[string]interface{}{})
		return
	} else {
		var acculist []float64
		for _, value := range recognitionList {
			acculist = append(acculist, value.Accuracy)
		}
		res := recognitionList[floats.MinIdx(acculist)]
		fmt.Println(res.UserID)

		if res.Accuracy <= 0.2 {
			attendanceBody := models.AttendanceBody{
				UserID:        res.UserID,
				CameraID:      res.CameraID,
				PhotoEncoding: res.PhotoEncoding,
			}

			log = models.Log{
				UserID:   res.UserID,
				CameraID: res.CameraID,
			}
			newAttendance(w, attendanceBody)
			insertLog := logStore(log)
			fmt.Println("insert log :", insertLog)
		}

		respondJSON(w, 200, "Returned Matching Identities", map[string]interface{}{
			"user_id":  res.UserID,
			"name":     res.Name,
			"accuracy": res.Accuracy,
			"elapsed":  res.Elapsed,
		})
		return
	}
}

func euclideanDistance(emb1, emb2 []float64) float64 {
	val := 0.0
	for i := range emb1 {
		val += math.Pow(emb1[i]-emb2[i], 2)
	}
	return val
}
