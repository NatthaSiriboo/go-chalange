package main
import (
	"go-challenge/cipher"
	"errors"
	"flag"
	"os"
	"io/ioutil"
	//"io"
	"fmt"
	"log"
)
func getFile() (string, error) {
	// We need to validate that we're getting the correct number of arguments
	if len(os.Args) < 2 {
		return "no file detect", errors.New("A filepath argument is required")
	}
	flag.Parse() // This will parse all the arguments from the terminal

	fileLocation := flag.Arg(0)
	return fileLocation, nil
}


func main() {
	fileLocation, err:= getFile()
	if err != nil {
		log.Fatal(err)
	}
	//content := make([]byte, 499)
	//fmt.Print(content[8])
	content, err := ioutil.ReadFile(fileLocation)
	//fmt.Print(content[8])
	//fmt.Print("*******")
	if err != nil {
		log.Fatal(err)
	}
	//content := make(byte[], 4096)
	reader, err := os.Open(fileLocation) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	//var reader io.Reader 
	rot128reader, err := cipher.NewRot128Reader(reader) 
	if err != nil {
		log.Fatal(err)
	}
	final, err := rot128reader.Read(content)
	fmt.Print(string(content))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print("***&&&****")
	fmt.Print(final)
} 