package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

type Composite struct {
	XMLName    xml.Name `xml:"attribute"`
	Attributes []Attribute
}

type Color struct {
	XMLName xml.Name `xml:"color"`
	Val     string   `xml:"val,attr"`
}

type Point struct {
	XMLName xml.Name `xml:"point"`
	Val     string   `xml:"val,attr"`
}
type Real struct {
	XMLName xml.Name `xml:"real"`
	Val     float64  `xml:"val,attr"`
}
type Rectangle struct {
	XMLName xml.Name `xml:"rectangle"`
	Val     string   `xml:"val,attr"`
}

type Attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Name      string   `xml:"name,attr"`
	Color     Color
	Point     Point
	Real      Real
	String    string `xml:"string"`
	Rectangle Rectangle
	Composite Composite
}

type Object struct {
	XMLName   xml.Name    `xml:"object"`
	ID        string      `xml:"id,attr"`
	Attribute []Attribute `xml:"attribute"`
}

type Layer struct {
	XMLName xml.Name `xml:"layer"`
	Object  []Object `xml:"object"`
}

type Dia struct {
	XMLName xml.Name `xml:"diagram"`
	Layer   Layer
}

type Current struct {
	obj_pos     string
	obj_bb      string
	elem_corner string
	elem_width  float64
	elem_height float64
	fill_color  string
	text_color  string
	line_color  string
}

var currentData map[string]*Current

func parseXML(br io.Reader) {
	b, _ := ioutil.ReadAll(br)

	var dia Dia
	err := xml.Unmarshal(b, &dia)
	if err != nil {
		fmt.Println(err)
	}

	// Build storage for existing locations. color, etc.
	currentData = make(map[string]*Current)
	for _, obj := range dia.Layer.Object {
		var name string
		current := &Current{}
		for _, attr := range obj.Attribute {
			if attr.Name == "name" {
				name = attr.String[1 : len(attr.String)-1]
			} else if attr.Name == "stereotype" {
				name = attr.String[1:len(attr.String)-1] + "." + name
			} else if attr.Name == "obj_pos" {
				current.obj_pos = attr.Point.Val
			} else if attr.Name == "obj_bb" {
				current.obj_bb = attr.Rectangle.Val
			} else if attr.Name == "elem_corner" {
				current.elem_corner = attr.Point.Val
			} else if attr.Name == "elem_width" {
				current.elem_width = attr.Real.Val
			} else if attr.Name == "elem_height" {
				current.elem_height = attr.Real.Val
			} else if attr.Name == "line_color" {
				current.line_color = attr.Color.Val
			} else if attr.Name == "text_color" {
				current.text_color = attr.Color.Val
			} else if attr.Name == "fill_color" {
				current.fill_color = attr.Color.Val
			}
		}
		if len(name) == 0 {
			continue
		}
		//fmt.Printf("%s = %v\n", name, current)
		currentData[name] = current
	}
}

func escape(in string) string {
	s := bytes.NewBufferString("")
	xml.EscapeText(s, []byte(in))
	return s.String()
}
