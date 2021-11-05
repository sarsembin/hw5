package main

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestStringOrNumber(t *testing.T) {
	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		t.Errorf("Error while json.Unmarshal: %v", err)
	}

	want := []User{
		{
			ID:      CustomFloat64{1},
			Address: Address{
				CityID: CustomFloat64{5} ,
				Street: "Satbayev",
			},
			Age:     CustomFloat64{20},
		},
		{
			ID:      CustomFloat64{1},
			Address: Address{
				CityID: CustomFloat64{6},
				Street: "Al-Farabi",
			},
			Age:     CustomFloat64{32},
		},
	}

	if !reflect.DeepEqual(want, users) {
		t.Errorf("Result of json.Marshall was wrong, got: %v, want: %v\n", users, want)
	}

	var xmlUsers Users
	if err := xml.Unmarshal(rawXml, &xmlUsers); err != nil {
		t.Errorf("Error while xml.Unmarshal: %v", err)
	}

	if !reflect.DeepEqual(want, xmlUsers.Users) {
		t.Errorf("Result of xml.Marshall was wrong, got: %v, want: %v\n", xmlUsers.Users, want)
	}
}
