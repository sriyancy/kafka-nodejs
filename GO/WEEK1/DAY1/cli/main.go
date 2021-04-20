//.go indicates we work with source file
//start with poackage declaration -indicate compiler we creae executeable package
package main

//entry point for app
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	
)

func main() {
	path:=flag.String("path","myapp.log","The path to the log that should be analyzed")
	level:=flag.String("level","ERROR","Log level to search for.Options are DEBUG,INFO,ERROR AND CRITICAL")
	flag.Parse()
	//few lang accepts arg in parameter for this fun and accept cl parameters
	//we are not forced to deal woth every tym
	//accepts no parameter returns no values
	//er need to access log file
	//done by open fun from os system package
	f, err := os.Open(*path) //f file handler,error are returnd
	//no exception in go rather we return error values
	if err != nil {
		//nill similar to null
		//pointer null
		log.Fatal(err)
	}
	defer f.Close() //defer-after main exits or fun called exits itit go ahead and close at end
	//read content of file
	//use buff io package n call read sometype of inpout stream
	//here file but it maybe anything from n/w or string
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n') //read all string till new line
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s) //formatted i/o
		}
	}
}
