package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/********************

Should be missing
*****************/

type Field struct {
	ID      int
	Name    string // Name of field with comment space
	Type    string //Type of field without comment space
	Comment string
	Line    string
}

type Strct struct {
	ID     int
	Parent *GoStructs
	Name   string
	Line   string
	Fields map[string]*Field
}

type GoStructs struct {
	Package string
	Strcts  map[string]*Strct
}

type GoFiles struct {
	Files map[string]*GoStructs
}

func (g *GoFiles) CreateDia(fn string) {
	// Need to build a package.fieldName map for connections
	pfn := make(map[string]*Strct)
	for _, a := range g.Files {
		for _, b := range a.Strcts {
			name := fmt.Sprintf("%s.%s", a.Package, b.Name)
			pfn[name] = b
		}
	}

	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	sendHeader(f)
	for _, a := range g.Files {
		for _, b := range a.Strcts {
			sendObject(f, b.ID, b)
			// Any connections?
			for _, c := range b.Fields {
				typ := c.Type
				if len(typ) > 5 {
					if typ[:4] == "map[" {
						typ = typ[strings.Index(typ, "]")+1:]
					} else if typ[:5] == "chan " {
						typ = typ[5:]
					}
				}
				if len(typ) > 3 {
					if typ[0] == '[' {
						pos := strings.Index(typ, "]")
						typ = typ[pos+1:]
					}
				}
				// Pointer? Dereference
				if typ[0] == '*' {
					typ = typ[1:]
				}
				// Is this type from another package?
				pkg := a.Package
				if strings.Index(typ, ".") > 0 {
					pkg = typ[:strings.Index(typ, ".")]
					typ = typ[strings.Index(typ, ".")+1:]
				}
				name := fmt.Sprintf("%s.%s", pkg, typ)
				if x, ok := pfn[name]; ok {
					sendImplements(f, gid, b.ID, c.ID*2+9, x.ID, 3)
				}
			}
		}
	}
	sendFooter(f)
}
