package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: godia inputDirectory outputFilename [comma delimited list of directories to ignore]")
		os.Exit(1)
	}
	in := os.Args[1]
	if in[len(in)-1] != '/' {
		in += "/"
	}
	out := os.Args[2]
	ignore := make(map[string]bool)
	if len(os.Args) == 4 {
		c := strings.Split(os.Args[3], ",")
		for _, a := range c {
			ignore[strings.TrimSpace(a)] = true
		}
	}
	//in := "/home/jgettys/Development/goepc_project/goepc/"
	//out := "goepc"
	//in := "./"
	//out := "test"
	g := GoFiles{Files: make(map[string]*GoStructs)}
	g.process(in, ignore)

	// If the ourput.dia file exists, gunzip to xml so parser can read in positions
	if f, err := os.Open(out + ".dia"); err == nil {
		f.Close()
		dia, err := ReadGzipFile(out + ".dia")
		if err == nil {
			f2 := bytes.NewReader(dia)
			parseXML(f2)
		}
	}

	g.CreateDia(out + ".xml")
	bdata, err := ioutil.ReadFile(out + ".xml")
	if err != nil {
		log.Fatal(err)
	}
	CreateGzipFile(out+".dia", bdata)
	fmt.Printf("Successfully created %s.dia\n", out)
	fmt.Printf("Processed %d files, %d structs and %d fields\n", fileCount, structCount, fieldCount)

}
