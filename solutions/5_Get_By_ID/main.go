/*
Have 2 endpoints?
- all squirrels
- squirrels by id

- read file once

*/
//func handleSquirrelByID(w http.ResponseWriter, r *http.Request) {
//	fileContents, err := os.ReadFile("../../static/data.json")
//	if err != nil {
//		panic(err)
//	}
//	list := []SquirrelSighting{}
//	err = json.Unmarshal(fileContents, list)
//	if err != nil {
//		panic(err)
//	}
//
//	first := list[0]
//	resp, err := json.Marshal(first)
//	if err != nil {
//		panic(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(resp)
//}
