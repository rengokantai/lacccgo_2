package main
import (
	"fmt"
	"log"
	"os"
	"image/jpeg"
)

func main() {
	f, err := os.Open("./1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//fmt.Println(f) //this will print out address of image
	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IMAGE TYPE: %T \n", img)
}