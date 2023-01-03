package bot

import (
	"errors"
	"github/mrqwer/slangTelBot/checker"
	"github/mrqwer/slangTelBot/database"
	"log"

	edlib "github.com/hbollon/go-edlib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func formatPrint(arr []string) string {
//	s := ""
//	for _, v := range arr {
//		s += (v + " ")
//	}
//	return s
//}

func Sum(l []float32) float32 {
	var s float32
	for _, v := range l {
		s += v
	}
	return s
}

func Max(l []float32) (float32, int) {
	var (
		m     float32
		index int
	)
	for i, v := range l {
		if m < v {
			m = v
			index = i
		}
	}
	return m, index
}

func Result(data *[]database.Collection, index int) (string, error) {
	for i, v := range *data {
		if i == index {
			return v.Standard, nil
		}
	}
	return "", nil
}

func findStandard(s interface{}) (string, error) {
	if !checker.ValidWord(s) {
		return "", errors.New("invalid type in search algorithm")
	}
	newS := checker.Lower(s.(string))
	filterDoc := bson.M{}
	findOptions := options.Find()
	data, err := database.GetMongoDocs(database.Dictionary, filterDoc, findOptions)
	if err != nil {
		log.Printf("When extracting data and putting into slice of struct\n%v", err)
	}
	type listOfMatchings []float32
	correspondenceDict := make(map[int]listOfMatchings)

	for i, v := range *data {
		temp := make([]float32, len(v.Slang))
		for j, k := range v.Slang {
			temp[j] = edlib.JaroWinklerSimilarity(newS, k)
		}
		correspondenceDict[i] = temp
	}

	listOfSums := make([]float32, len(correspondenceDict))
	var (
		index int
		m     float32
	)
	for k, v := range correspondenceDict {
		listOfSums[k] = Sum(v)
		if m, index = Max(v); m >= 9.5 {
			break
		}
	}
	if m >= 9.5 {
		return Result(data, index)
	} else {
		_, j := Max(listOfSums)
		return Result(data, j)
	}
}
