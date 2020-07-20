package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct{
	Page int
	Fruits []string
}
type response2 struct{
	Page int	`json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bloB, _ := json.Marshal(true)
	fmt.Println(string(bloB))
	
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	
	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))
		
	stringB, _ := json.Marshal("hello world")
	fmt.Println(string(stringB))
		
	slcB, _ := json.Marshal([]string{"apple", "pineapple", "kivi"})
	fmt.Println(string(slcB))
		
	mapB, _ := json.Marshal(map[string] int{"apple": 5, "pineapple": 37, "kivi": 3})
	fmt.Println(string(mapB))
		
	resD1 :=  &response1 {
		Page: 1,
		Fruits: []string{"apple", "kivi", "grapes"},
	}
	resB1, _ := json.Marshal(resD1)
	fmt.Println(string(resB1))
	resD2 :=  &response2 {
		Page: 1,
		Fruits: []string{"apple", "kivi", "grapes"},
	}
	resB2, _ := json.Marshal(resD2)
	fmt.Println(string(resB2))

	byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"].(float64))
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}