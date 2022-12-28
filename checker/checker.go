package checker

//var (
//	c1 = database.Collection{"windows", []string{"винда", "виндовс", "win", "вин", "шиндовс", "венда",
//		"шиндоуз"}}
//	c2 = database.Collection{"java", []string{"джава", "ява", "жаба"}}
//	c3 = database.Collection{"ubuntu", []string{"убунта", "бубунта", "хубунту", "убунт"}}
//	с4 = database.Collection{"python", []string{"питон", "пайтон", "петон", "путон", "питоний"}}
//
//	//data := bson.M
//)

func isValidWord(v interface{}) (bool, error) {
	switch v.(type) {
	case string:
		return true, nil
	default:
		return false, nil
	}
}

func createData() {

}
