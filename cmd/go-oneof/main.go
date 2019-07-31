package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gertd/go-oneof/pkg/business"
)

type test struct {
	obj business.Info
	str string
}

func testData() []test {

	tests := []test{
		{business.Info{
			Name:        "Homer",
			AddressKind: business.AddressKindHome,
			Address: business.Home{
				HouseNumber: "123",
				Street:      "Home Street",
				City:        "Seattle",
				State:       "WA",
				PostalCode:  "98101",
			},
			Country: "USA",
		},
			`{
  "name": "Homer",
  "addressKind": "Home",
  "address": {
    "houseNumber": "123",
    "street": "Home Street",
    "city": "Seattle",
    "state": "WA",
    "postalCode": "98101"
  },
  "country": "USA"
}`},
		{business.Info{
			Name:        "Buro",
			AddressKind: business.AddressKindOffice,
			Address: business.Office{
				OfficeNumber: "987",
				Street:       "Office Park",
				City:         "Bellevue",
				State:        "WA",
				PostalCode:   "98004",
			},
			Country: "USA",
		},
			`{
  "name": "Buro",
  "addressKind": "Office",
  "address": {
    "officeNumber": "987",
    "street": "Office Park",
    "city": "Bellevue",
    "state": "WA",
    "postalCode": "98004"
  },
  "country": "USA"
}`},
		{business.Info{
			Name:        "Going Postal",
			AddressKind: business.AddressKindPostalBox,
			Address: business.PostalBox{
				POBoxNumber: "30125",
				City:        "Seattle",
				State:       "WA",
				PostalCode:  "98113-0125",
			},
			Country: "USA",
		},
			`{
  "name": "Going Postal",
  "addressKind": "PostalBox",
  "address": {
    "poBoxNumber": "30125",
    "city": "Seattle",
    "state": "WA",
    "postalCode": "98113-0125"
  },
  "country": "USA"
}`},
	}
	return tests
}

func main() {
	infoMarshal()
	infoUnmarshal()
}

func infoMarshal() {
	tests := testData()
	for i, test := range tests {

		b, err := json.MarshalIndent(test.obj, "", "  ")
		if err != nil {
			fmt.Printf("test %d -- err %v", i, err)
		}
		if string(b) != test.str {
			fmt.Printf("Test %d\nExpected:\n%s\nActual:\n%s\n", i, test.str, string(b))
		}
	}
}

func infoUnmarshal() {
	tests := testData()
	for i, test := range tests {
		var info business.Info
		err := json.Unmarshal([]byte(test.str), &info)
		if err != nil {
			fmt.Printf("test %d -- err %v", i, err)
		}
		if !reflect.DeepEqual(info, test.obj) {
			fmt.Printf("Test %d\nExpected:\n%v\nActual:\n%v\n", i, test.obj, info)
		}
	}
}
