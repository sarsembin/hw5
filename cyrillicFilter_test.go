package main

import (
	"reflect"
	"testing"
)

func TestCyrillicFilter(t *testing.T) {

	type nested struct {
		EngWord string
		RusWord string
		MixedWord *string
	}

	type sample struct {
		EngWord string
		Num int
		RusWord string
		RusWord2 string
		MixedWord *string
		Nest *nested
	}

	str := "This word is cyrillic: борщ, end of word"
	nestedStr := "This word is cyrillic 2: борщ, end of word"
	nest := nested{
		EngWord: "sus",
		RusWord: "ёлка is a tree",
		MixedWord: &nestedStr,
	}

	val := sample{
		EngWord: "amogus",
		Num: 666,
		RusWord: "amoguмогус",
		RusWord2: "Ёжик",
		MixedWord: &str,
		Nest: &nest,
	}

	wantVal := sample{
		EngWord: "amogus",
		Num: 666,
		RusWord: "amogu",
		RusWord2: "",
		MixedWord: &str,
		Nest: &nest,
	}
	wantNest := nested{
		EngWord: "sus",
		RusWord: " is a tree",
		MixedWord: &nestedStr,
	}
	wantStr := "This word is cyrillic: , end of word"
	wantNestedStr := "This word is cyrillic 2: , end of word"

	CyrillicFilter(&val)

	if !reflect.DeepEqual(val, wantVal) {
		t.Errorf("Formatted struct was wrong, got: %v, want %v\n", val, wantVal)
	}
	if !reflect.DeepEqual(nest, wantNest) {
		t.Errorf("Formatted nested struct was wrong, got: %v, want %v\n", nest, wantNest)
	}
	if wantStr != str {
		t.Errorf("Formatted pointer on string value was wrong, got: %v, want %v\n", str, wantStr)
	}
	if wantNestedStr != nestedStr {
		t.Errorf("Formatted pointer on string value was wrong, got: %v, want %v\n", nestedStr, wantNestedStr)
	}
}
