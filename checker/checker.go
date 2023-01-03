package checker

import (
	"strings"
)

//var (
//	c1 = database.Collection{"windows", []string{"винда", "виндовс", "win", "вин", "шиндовс", "венда",
//		"шиндоуз"}}
//	c2 = database.Collection{"java", []string{"джава", "ява", "жаба"}}
//	c3 = database.Collection{"ubuntu", []string{"убунта", "бубунта", "хубунту", "убунт"}}
//	с4 = database.Collection{"python", []string{"питон", "пайтон", "петон", "путон", "питоний"}}
//
//	//data := bson.M
//)

func ValidWord(v interface{}) bool {
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

func Lower(s string) string {
	return strings.ToLower(s)
}
