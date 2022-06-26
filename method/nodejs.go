package method

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("https://d1ny9casiyy5u5.cloudfront.net/tmp/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
}
