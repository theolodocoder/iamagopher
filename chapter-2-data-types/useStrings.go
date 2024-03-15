package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))
	// index considers the case of the words
	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	// hasprefix also considers the case
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	// Fields
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	fmt.Println(t)
	t = s.Fields("")
	fmt.Println(t)
	f("Fields: %v\n", len(t))
	// Split string by delimeter
	f("%s\n", s.Split("ab+cd+efg", "+"))
	// Replace
	/*
		The strings.Replace() function takes four parameters. The first parameter is the
		string that you want to process. The second parameter contains the string that,
		if found, will be replaced by the third parameter of strings.Replace(). The last
		parameter is the maximum number of replacements that are allowed to happen.
		If that parameter has a negative value, then there is no limit to the number of
		replacements that can take place.

	*/
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	// Spjit after
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	// TrimFunc takes in a callback function of what you want to trin by
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
