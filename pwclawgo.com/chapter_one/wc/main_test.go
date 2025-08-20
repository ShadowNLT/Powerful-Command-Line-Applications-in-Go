package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := Result{Words: 4, Lines: 0, Bytes: 0}
	res, err := count(b, false, false)
	if err != nil {
		t.Error("An unexpected error occured: ", err)
	}
	if res != exp {
		t.Errorf("Expected %v, got %v instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2 Word3\nLine2\nLine3 Word 1")

	expected := Result{Words: 0, Lines: 3, Bytes: 0}

	result, err := count(b, true, false)
	if err != nil {
		t.Error("An unexpected error occured: ", err)
	}

	if result != expected {
		t.Errorf("Expected %v, got %v instead.\n", expected, result)
	}
}

func TestCountWordsAndBytes(t *testing.T) {
	input := "Go is fun"

	b := bytes.NewBufferString(input)

	expected := Result{Words: 3, Lines: 0, Bytes: 9}
	result, err := count(b, false, true)
	if err != nil {
		t.Error("An unexpected error occured: ", err)
	}

	if result != expected {
		t.Errorf("Expected %v, got %v instead \n", expected, result)
	}
}

func TestCountLinesAndBytes(t *testing.T) {
	input := "hello\nworld\nchatgpt\n"

	b := bytes.NewBufferString(input)

	expected := Result{Words: 0, Lines: 3, Bytes: 20}
	result, err := count(b, true, true)
	if err != nil {
		t.Error("An unexpected error occured: ", err)
	}

	if result != expected {
		t.Errorf("Expected %v, got %v instead \n", expected, result)
	}
}
