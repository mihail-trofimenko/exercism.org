package main

import "regexp"

// Hey should have a comment documenting it.

//"Sure." This is his response if you ask him a question, such as "How are you?" The convention used for questions is that it ends with a question mark.
//"Whoa, chill out!" This is his answer if you YELL AT HIM. The convention used for yelling is ALL CAPITAL LETTERS.
//"Calm down, I know what I'm doing!" This is what he says if you yell a question at him.
//"Fine. Be that way!" This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characters.
//"Whatever." This is what he answers to anything else.
// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	//regexp compilation templates
	sure := `^[0-9a-zA-Z%\^*@#$%&)(-=!., ]+\?\s*$`
	whoa := `^[A-Z0-9%\^*@#$%&)(-=, !]*[A-Z][A-Z0-9%\^*@#$%&)(-=, !]*[.!]?$`
	calm := `^[A-Z\', ]+\?$`
	fine := `^[\s*]+$`

	match, _ := regexp.MatchString(sure, remark)
	if match {
		return "Sure."
	}
	match, _ = regexp.MatchString(whoa, remark)
	if match {
		return "Whoa, chill out!"
	}
	match, _ = regexp.MatchString(calm, remark)
	if match {
		return "Calm down, I know what I'm doing!"
	}
	match, _ = regexp.MatchString(fine, remark)
	if match || remark == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}