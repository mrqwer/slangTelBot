package bot

import (
	"errors"
	"fmt"
	"github/mrqwer/slangTelBot/checker"
	"github/mrqwer/slangTelBot/database"
	"log"
	"strings"

	"github.com/hbollon/go-edlib"
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

func topSlangs() string {
	result := ""
	for i := range dictionary {
		f := bson.M{"standard": dictionary[i]}
		t, err := database.GetMongoDoc(database.Dictionary, f)
		if err != nil {
			log.Fatal(err)
		}
		resSlang := ""
		for _, v := range t.Slang {
			resSlang += v + ", "
		}
		resSlang = strings.Trim(resSlang, ", ")
		result += "*Международный стандарт:* " + strings.Title(t.Standard) + "\n" + "**Сленги**: " + resSlang
		result += "\n\n"
	}
	return result
}

func Sum(l []float32) float32 {
	var s float32
	for _, v := range l {
		s += v
	}
	return s
}

func Max(l []float32) (float32, int) {
	var (
		m     float32 = l[0]
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

// топ 10 записей
//func top10Docs() {
//	top10 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
//
//}

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
		fmt.Println(v)
		fmt.Println(newS)
		for j, k := range v.Slang {
			temp[j] = edlib.JaroWinklerSimilarity(newS, k)
		}
		fmt.Println(temp)
		temp2 := make([]float32, 2)
		a, b := Max(temp)
		temp2[0], temp2[1] = a, float32(b)
		correspondenceDict[i] = temp2
	}
	fmt.Println(correspondenceDict)
	resMap := make(map[int]listOfMatchings)
	for k, v := range correspondenceDict {
		if v[0] >= 0.85 {
			resMap[k] = v
		}
	}
	fmt.Println("\n\n\n")
	fmt.Println(resMap)

	//	if t, err := searchByMongo(newS); err == nil {
	//		fmt.Println("Search by Mongo db")
	//		fmt.Printf("%v", newS)
	//		fmt.Println()
	//		return t, nil
	//	} else {
	if len(resMap) == 0 {
		return "", nil
	} else {
		r := ""
		for i, v := range *data {
			if _, containsKey := resMap[i]; containsKey {
				fmt.Println(v.Standard)
				r += " " + v.Standard
			}
		}
		return r, nil
	}

	//	listOfSums := make([]float32, len(correspondenceDict))
	//	var (
	//		index int
	//		m     float32
	//	)
	//	for k, v := range correspondenceDict {
	//		listOfSums[k] = Sum(v)
	//		if m, index = Max(v); m >= 8.5 {
	//			break
	//		}
	//	}
	//	fmt.Print(m, index)
	//	if t, err := searchByMongo(newS); err == nil {
	//		fmt.Println("Search by Mongo db")
	//		fmt.Printf("%v", newS)
	//		fmt.Println()
	//		return t, nil
	//	} else {
	//		_, j := Max(listOfSums)
	//
	//		fmt.Println("Search by Jarowinkler")
	//		fmt.Printf("%v\n", listOfSums)
	//		fmt.Printf("%v\n", j)
	//
	//		s, err := Result(data, j)
	//		if err != nil {
	//			return "", err
	//		}
	//
	//		return standardWithDef(s), nil
	//	}
}

//func compareStrings(s1, s2 string) bool {
//	return s1 == s2
//}

// func standardWithDef(result string) string {
// 	filter := bson.M{"standard": result}
// 	data, err := database.GetMongoDoc(database.Dictionary, filter)
// 	if err != nil {
// 		return ""
// 	}
// 	return data.Standard + "\n" + data.Definition
// }

//func searchByMongo(s string) (string, error) {
//	//filterDoc := bson.M{"$text": bson.M{"$search": s}}
//	filterDoc := bson.M{"slangs": s}
//	data, err := database.GetMongoDoc(database.Dictionary, filterDoc)
//	if err != nil {
//		return "", err
//	}
//	return data.Standard, nil
//}
