package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page  int
	Fruit []string
}

type response2 struct {
	Page  int      `json:"page"`
	Fruit []string `json:"fruit"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))
	fmt.Println(fltB)

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	fmt.Println(strB)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB)
	fmt.Println(string(slcB))

	res1D := &response1{
		Page:  1,
		Fruit: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:  1,
		Fruit: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	fmt.Println(strs)

	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruit)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
