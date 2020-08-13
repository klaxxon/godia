package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	xmlonly := flag.Bool("xml", false, "Uncompressed xml file only")
	in := flag.String("i", "", "Directory to scan")
	out := flag.String("o", "", "Output file (will add extension)")
	ignoreDirs := flag.String("ignore", "", "Comma delimited directories to ignore")
	flag.Parse()

	if *in == "" {
		log.Fatal("Must include directory to scan.")
	}
	if *out == "" {
		log.Fatal("Must include output filename.")
	}
	if (*in)[len(*in)-1] != '/' {
		*in += "/"
	}
	ignore := make(map[string]bool)
	if *ignoreDirs != "" {
		c := strings.Split(*ignoreDirs, ",")
		for _, a := range c {
			ignore[strings.TrimSpace(a)] = true
		}
	}
	//in := "/home/jgettys/Development/goepc_project/goepc/"
	//out := "goepc"
	//in := "./"
	//out := "test"
	g := GoFiles{Files: make(map[string]*GoStructs)}
	g.process(*in, ignore)

	// If the ourput.dia file exists, gunzip to xml so parser can read in positions
	if f, err := os.Open(*out + ".dia"); err == nil {
		f.Close()
		dia, err := ReadGzipFile(*out + ".dia")
		if err == nil {
			f2 := bytes.NewReader(dia)
			parseXML(f2)
		}
	}

	g.CreateDia(*out + ".xml")
	if !*xmlonly {
		bdata, err := ioutil.ReadFile(*out + ".xml")
		if err != nil {
			log.Fatal(err)
		}
		CreateGzipFile(*out+".dia", bdata)
		fmt.Printf("Successfully created %s.dia\n", *out)
	} else {
		fmt.Printf("Successfully created %s.xml\n", *out)
	}
	fmt.Printf("Processed %d files, %d structs and %d fields\n", fileCount, structCount, fieldCount)

}
