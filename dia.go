package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sendHeader(f *os.File) {
	f.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
	<dia:diagram xmlns:dia="http://www.lysator.liu.se/~alla/dia/">
		<dia:diagramdata>
			<dia:attribute name="background">
				<dia:color val="#ffffffff"/>
			</dia:attribute>
			<dia:attribute name="pagebreak">
				<dia:color val="#000099ff"/>
			</dia:attribute>
			<dia:attribute name="paper">
				<dia:composite type="paper">
					<dia:attribute name="name">
						<dia:string>#Letter#</dia:string>
					</dia:attribute>
					<dia:attribute name="tmargin">
						<dia:real val="2.5399999618530273"/>
					</dia:attribute>
					<dia:attribute name="bmargin">
						<dia:real val="2.5399999618530273"/>
					</dia:attribute>
					<dia:attribute name="lmargin">
						<dia:real val="2.5399999618530273"/>
					</dia:attribute>
					<dia:attribute name="rmargin">
						<dia:real val="2.5399999618530273"/>
					</dia:attribute>
					<dia:attribute name="is_portrait">
						<dia:boolean val="true"/>
					</dia:attribute>
					<dia:attribute name="scaling">
						<dia:real val="1"/>
					</dia:attribute>
					<dia:attribute name="fitto">
						<dia:boolean val="false"/>
					</dia:attribute>
				</dia:composite>
			</dia:attribute>
			<dia:attribute name="grid">
				<dia:composite type="grid">
					<dia:attribute name="dynamic">
						<dia:boolean val="true"/>
					</dia:attribute>
					<dia:attribute name="width_x">
						<dia:real val="1"/>
					</dia:attribute>
					<dia:attribute name="width_y">
						<dia:real val="1"/>
					</dia:attribute>
					<dia:attribute name="visible_x">
						<dia:int val="1"/>
					</dia:attribute>
					<dia:attribute name="visible_y">
						<dia:int val="1"/>
					</dia:attribute>
					<dia:composite type="color"/>
				</dia:composite>
			</dia:attribute>
			<dia:attribute name="color">
				<dia:color val="#d8e5e5ff"/>
			</dia:attribute>
			<dia:attribute name="guides">
				<dia:composite type="guides">
					<dia:attribute name="hguides"/>
					<dia:attribute name="vguides"/>
				</dia:composite>
			</dia:attribute>
			<dia:attribute name="display">
				<dia:composite type="display">
					<dia:attribute name="antialiased">
						<dia:boolean val="false"/>
					</dia:attribute>
					<dia:attribute name="snap-to-grid">
						<dia:boolean val="false"/>
					</dia:attribute>
					<dia:attribute name="snap-to-object">
						<dia:boolean val="true"/>
					</dia:attribute>
					<dia:attribute name="show-grid">
						<dia:boolean val="true"/>
					</dia:attribute>
					<dia:attribute name="show-connection-points">
						<dia:boolean val="true"/>
					</dia:attribute>
				</dia:composite>
			</dia:attribute>
		</dia:diagramdata>
		<dia:layer name="Background" visible="true" connectable="true" active="true">` + "\n"))
}

func getCurrentData(name string) *Current {
	current := currentData[name]
	if current == nil {
		current = &Current{}
		current.obj_pos = "3.75,9.75"
		current.obj_bb = "3.7,9.7;11.23,14.2"
		current.elem_corner = "3.75,9.75"
		current.elem_height = 4.4
		current.elem_width = 7.43
		current.fill_color = "#ffffffff"
		current.line_color = "#000000ff"
		current.text_color = "#000000ff"
	}
	return current
}

func sendObject(f *os.File, id int, s *Strct) {
	name := s.Parent.Package + "." + s.Name
	current := getCurrentData(name)
	f.Write([]byte(fmt.Sprintf(`<dia:object type="UML - Class" version="0" id="O%d">
	<dia:attribute name="obj_pos">
		<dia:point val="%s"/>
	</dia:attribute>
	<dia:attribute name="obj_bb">
		<dia:rectangle val="%s"/>
	</dia:attribute>
	<dia:attribute name="elem_corner">
		<dia:point val="%s"/>
	</dia:attribute>
	<dia:attribute name="elem_width">
		<dia:real val="%f"/>
	</dia:attribute>
	<dia:attribute name="elem_height">
		<dia:real val="%f"/>
	</dia:attribute>
	<dia:attribute name="name">
		<dia:string>#%s#</dia:string>
	</dia:attribute>
	<dia:attribute name="stereotype">
		<dia:string>#%s#</dia:string>
	</dia:attribute>
	<dia:attribute name="comment">
		<dia:string>#%s#</dia:string>
	</dia:attribute>
	<dia:attribute name="abstract">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="suppress_attributes">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="suppress_operations">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="visible_attributes">
		<dia:boolean val="true"/>
	</dia:attribute>
	<dia:attribute name="visible_operations">
		<dia:boolean val="true"/>
	</dia:attribute>
	<dia:attribute name="visible_comments">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="wrap_operations">
		<dia:boolean val="true"/>
	</dia:attribute>
	<dia:attribute name="wrap_after_char">
		<dia:int val="40"/>
	</dia:attribute>
	<dia:attribute name="comment_line_length">
		<dia:int val="17"/>
	</dia:attribute>
	<dia:attribute name="comment_tagging">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="allow_resizing">
		<dia:boolean val="false"/>
	</dia:attribute>
	<dia:attribute name="line_width">
		<dia:real val="0.10000000000000001"/>
	</dia:attribute>
	<dia:attribute name="line_color">
		<dia:color val="%s"/>
	</dia:attribute>
	<dia:attribute name="fill_color">
		<dia:color val="%s"/>
	</dia:attribute>
	<dia:attribute name="text_color">
		<dia:color val="%s"/>
	</dia:attribute>
	<dia:attribute name="normal_font">
		<dia:font family="monospace" style="0" name="Courier"/>
	</dia:attribute>
	<dia:attribute name="abstract_font">
		<dia:font family="monospace" style="88" name="Courier-BoldOblique"/>
	</dia:attribute>
	<dia:attribute name="polymorphic_font">
		<dia:font family="monospace" style="8" name="Courier-Oblique"/>
	</dia:attribute>
	<dia:attribute name="classname_font">
		<dia:font family="sans" style="80" name="Helvetica-Bold"/>
	</dia:attribute>
	<dia:attribute name="abstract_classname_font">
		<dia:font family="sans" style="88" name="Helvetica-BoldOblique"/>
	</dia:attribute>
	<dia:attribute name="comment_font">
		<dia:font family="sans" style="8" name="Helvetica-Oblique"/>
	</dia:attribute>
	<dia:attribute name="normal_font_height">
		<dia:real val="0.80000000000000004"/>
	</dia:attribute>
	<dia:attribute name="polymorphic_font_height">
		<dia:real val="0.80000000000000004"/>
	</dia:attribute>
	<dia:attribute name="abstract_font_height">
		<dia:real val="0.80000000000000004"/>
	</dia:attribute>
	<dia:attribute name="classname_font_height">
		<dia:real val="1"/>
	</dia:attribute>
	<dia:attribute name="abstract_classname_font_height">
		<dia:real val="1"/>
	</dia:attribute>
	<dia:attribute name="comment_font_height">
		<dia:real val="0.69999999999999996"/>
	</dia:attribute>
	<dia:attribute name="attributes">`+"\n",
		id,
		current.obj_pos,
		current.obj_bb,
		current.elem_corner,
		current.elem_width,
		current.elem_height,
		s.Name,
		s.Parent.Package,
		s.Line,
		current.line_color,
		current.fill_color,
		current.text_color,
	)))
	// Bad sorting
	for x := 0; x < len(s.Fields); x++ {
		for _, a := range s.Fields {
			if a.ID != x {
				continue
			}
			f.Write([]byte(fmt.Sprintf(`<dia:composite type="umlattribute">
			<dia:attribute name="name">
				<dia:string>#%s-%d#</dia:string>
			</dia:attribute>
			<dia:attribute name="type">
				<dia:string>#%s#</dia:string>
			</dia:attribute>
			<dia:attribute name="value">
				<dia:string>##</dia:string>
			</dia:attribute>
			<dia:attribute name="comment">
				<dia:string>#%s#</dia:string>
			</dia:attribute>
			<dia:attribute name="visibility">
				<dia:enum val="0"/>
			</dia:attribute>
			<dia:attribute name="abstract">
				<dia:boolean val="false"/>
			</dia:attribute>
			<dia:attribute name="class_scope">
				<dia:boolean val="false"/>
			</dia:attribute>
		</dia:composite>`+"\n", a.Name, a.ID, a.Type, a.Comment)))
		}
	}
	f.Write([]byte(`</dia:attribute>
		<dia:attribute name="operations"/>
		<dia:attribute name="template">
			<dia:boolean val="false"/>
		</dia:attribute>
		<dia:attribute name="templates"/>
	</dia:object>` + "\n"))
}

func sendFooter(f *os.File) {
	f.Write([]byte(`</dia:layer>
	</dia:diagram>` + "\n"))
}

type Rect struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func parsePair(s string) (float64, float64) {
	comma := strings.Index(s, ",")
	x, err := strconv.ParseFloat(s[:comma], 64)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	y, err := strconv.ParseFloat(s[comma+1:], 64)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	return x, y
}

func getBoundingBoxFloat(s string) Rect {
	semi := strings.Index(s, ";")
	x1, y1 := parsePair(s[:semi])
	x2, y2 := parsePair(s[semi+1:])
	return Rect{x1: x1, y1: y1, x2: x2, y2: y2}
}

func sendImplements(f *os.File, id int, from *Strct, frompos int, to *Strct) {
	var fx, fy, tx, ty float64

	fpos := getCurrentData(fmt.Sprintf("%s.%s", from.Parent.Package, from.Name))
	tpos := getCurrentData(fmt.Sprintf("%s.%s", to.Parent.Package, to.Name))
	frect := getBoundingBoxFloat(fpos.obj_bb)
	trect := getBoundingBoxFloat(tpos.obj_bb)
	fx = frect.x2
	fy = frect.y1 + 2.7 + float64(frompos)*0.8
	tx = trect.x1
	ty = trect.y1 + 1.2
	frompos = frompos*2 + 9

	minx := math.Min(fx, tx)
	maxx := math.Max(fx, tx)
	miny := math.Min(fy, ty)
	maxy := math.Max(fy, ty)

	f.Write([]byte(fmt.Sprintf(`    <dia:object type="UML - Implements" version="0" id="O%d">
	<dia:attribute name="obj_pos">
		<dia:point val="%f,%f"/>
	</dia:attribute>
	<dia:attribute name="obj_bb">
		<dia:rectangle val="%f,%f;%f,%f"/>
	</dia:attribute>
	<dia:attribute name="meta">
		<dia:composite type="dict"/>
	</dia:attribute>
	<dia:attribute name="conn_endpoints">
		<dia:point val="%f,%f"/>
		<dia:point val="%f,%f"/>
	</dia:attribute>
	<dia:attribute name="text">
		<dia:string>##</dia:string>
	</dia:attribute>
	<dia:attribute name="text_font">
		<dia:font family="monospace" style="0" name="Courier"/>
	</dia:attribute>
	<dia:attribute name="text_height">
		<dia:real val="0.80000000000000004"/>
	</dia:attribute>
	<dia:attribute name="text_colour">
		<dia:color val="#000000ff"/>
	</dia:attribute>
	<dia:attribute name="text_pos">
		<dia:point val="%f,%f"/>
	</dia:attribute>
	<dia:attribute name="line_width">
		<dia:real val="0.10000000000000001"/>
	</dia:attribute>
	<dia:attribute name="line_colour">
		<dia:color val="#000000ff"/>
	</dia:attribute>
	<dia:attribute name="diameter">
		<dia:real val="0.69999999999999996"/>
	</dia:attribute>
	<dia:connections>
		<dia:connection handle="0" to="O%d" connection="%d"/>
		<dia:connection handle="1" to="O%d" connection="3"/>
	</dia:connections>
</dia:object>`, id,
		fx, fy,
		minx-0.2, miny-0.2, maxx+0.2, maxy+0.2,
		fx, fy, tx, ty,
		(minx+maxx)/2, (miny+maxy)/2,
		from.ID, frompos, to.ID)))
	fmt.Printf("Connection #%d %s --> %s\n", id, fmt.Sprintf("%s.%s", from.Parent.Package, from.Name), fmt.Sprintf("%s.%s", to.Parent.Package, to.Name))
	fmt.Printf("     O%d.%d ---> %s O%d.3\n", from.ID, frompos, to.Name, to.ID)
	fmt.Printf("     %0.1f,%0.1f -> %0.1f,%0.1f  (%0.1f,%0.1f)-(%0.1f,%0.1f)\n", fx, fy, tx, ty, minx, miny, maxx, maxy)
}
