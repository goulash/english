// Copyright (c) 2015, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package english makes it easy to write correct english.
package english

import (
	"bytes"
	"fmt"
	"reflect"
)

// Interface converts an array or slice of anything to a slice of interface,
// so it can be passed to List:
//
//  List("and", Interface(data)...)
//
func Interface(v interface{}) []interface{} {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		// OK
	default:
		panic("value is not an array or slice")
	}

	vv := reflect.ValueOf(v)
	n := vv.Len()
	out := make([]interface{}, n)
	for i := range out {
		out[i] = vv.Index(i).Interface()
	}
	return out
}

// List returns as a string the proper english for expressing a list of something.
// The conjunction is typically "or" or "and". Oxford comma is used.
func List(conjunction string, args ...interface{}) string {
	switch n := len(args); n {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(args[0])
	case 2:
		return fmt.Sprint(args[0], conjunction, args[1])
	default:
		var buf bytes.Buffer
		for _, v := range args[:n-1] {
			buf.WriteString(fmt.Sprint(v))
			buf.WriteString(", ")
		}
		buf.WriteString(conjunction)
		buf.WriteRune(' ')
		buf.WriteString(fmt.Sprint(args[n-1]))
		return buf.String()
	}
}

// S returns "s" if n is not 1.
func S(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
