package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"unicode"
)

var fileCount, structCount, fieldCount int

func (g *GoFiles) process(path string, ignore map[string]bool) {
	cwd := path
	c, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("ReadDir: %v", err)
		return
	}
	//fmt.Println("Listing ", cwd)
	for _, entry := range c {
		name := entry.Name()
		if _, ok := ignore[name]; ok {
			continue
		}
		if entry.IsDir() {
			path := fmt.Sprintf("%s%s/", cwd, name)
			g.process(path, ignore)
			continue
		}

		fpath := fmt.Sprintf("%s%s", cwd, name)
		if len(fpath) > 3 {
			if fpath[len(fpath)-3:] == ".go" {
				gs := &GoStructs{Strcts: make(map[string]*Strct)}
				if gs.parseGo(fpath) {
					fileCount++
					structCount += len(gs.Strcts)
					g.Files[fpath] = gs
				}
			}
		}
	}
}

var gid int = 0

func (g *GoStructs) parseGo(fn string) bool {
	bdata, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Println(err)
		return false
	}

	re := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	newBytes := re.ReplaceAll(bdata, nil)

	rdr := bytes.NewReader(newBytes)

	scanner := bufio.NewScanner(rdr)
	var strct *Strct
	linenum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			linenum++
			continue
		}
		if len(line) >= 2 && line[:2] == "//" {
			continue
		}
		words := parseLine(line)
		if strct == nil {
			if len(words) >= 2 && words[0] == "package" {
				g.Package = words[1]
				continue
			}
			if len(words) >= 3 {
				if words[0] == "type" && words[2] == "struct" {
					strct = &Strct{ID: gid, Parent: g, Name: words[1], Fields: make(map[string]*Field), Line: fmt.Sprintf("%s:%d", fn, linenum)}
					gid++
				}
			}
		} else {
			if words[0] == "}" {
				g.Strcts[strct.Name] = strct
				strct = nil
			} else {
				strct.parseField(words, len(strct.Fields), fmt.Sprintf("%s:%d", fn, linenum))
			}
		}
		linenum++
	}
	return len(g.Strcts) > 0
}

func (s *Strct) parseField(l []string, fid int, line string) {
	var f *Field
	// Composite?
	if len(l) == 1 {
		f = &Field{ID: fid, Name: l[0], Type: l[0], Line: line}
	} else {
		f = &Field{ID: fid, Name: l[0], Type: l[1], Line: line}
	}
	if len(l) > 2 {
		if l[1] == "chan" {
			f.Type = fmt.Sprintf("%s %s", l[1], l[2])
		} else if l[2][0] == '/' {
			f.Comment = strings.TrimSpace(l[2][2:])
			for a := 3; a < len(l); a++ {
				f.Comment += l[a] + " "
			}
			f.Comment = strings.TrimSpace(f.Comment)
		}
	}
	fieldCount++
	s.Fields[f.Name] = f
}

// parseLine returns the space delimited words and flag indicating whether it encountered
// 0 = no comment flag
// 1 = opening /*
// 2 = closing */
func parseLine(s string) []string {
	var str strings.Builder

	// If there are any comments, space them so they can be separated
	z := make([]string, 0)
	for _, c := range s {
		if unicode.IsSpace(c) {
			if str.Len() > 0 {
				z = append(z, str.String())
			}
			str.Reset()
		} else {
			str.WriteRune(c)
		}
	}
	if str.Len() > 0 {
		z = append(z, str.String())
	}
	return z
}
