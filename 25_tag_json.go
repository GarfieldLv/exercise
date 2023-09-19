package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name  string   `json:"title"`
	Year  int      `json:"year"`
	Price float32  `json:"price"`
	Actor []string `json:"actor"`
}

func main() {
	movie := Movie{
		Name:  "King of Comedy",
		Year:  2010,
		Price: 23.5,
		Actor: []string{"Zhou Xingchi", "Zhang Bozhi"},
	}
	// 序列化，从 struct -> json
	movieJson, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal failed:", err)
	}

	fmt.Printf("movieJson: %s\n", movieJson)

	// 反序列化，从 json -> struct
	var myMovie Movie
	err = json.Unmarshal(movieJson, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal failed: ", err)
	}

	fmt.Printf("myMovie: %v\n", myMovie)
}
