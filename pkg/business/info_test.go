package business

import (
	"encoding/json"
	"reflect"
	"testing"
)

type test struct {
	obj Info
	str string
}

func testData() []test {

	tests := []test{
		{Info{
			Name:        "Homer",
			AddressKind: AddressKindHome,
			Address: Home{
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
		{Info{
			Name:        "Buro",
			AddressKind: AddressKindOffice,
			Address: Office{
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
		{Info{
			Name:        "Going Postal",
			AddressKind: AddressKindPostalBox,
			Address: PostalBox{
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

func TestInfoMarshal(t *testing.T) {
	tests := testData()
	for i, test := range tests {

		b, err := json.MarshalIndent(test.obj, "", "  ")
		if err != nil {
			t.Errorf("test %d -- err %v", i, err)
		}
		if string(b) != test.str {
			t.Errorf("Test %d\nExpected:\n%s\nActual:\n%s\n", i, test.str, string(b))
		}
	}
}
func TestInfoUnmarshal(t *testing.T) {
	tests := testData()
	for i, test := range tests {
		var info Info
		err := json.Unmarshal([]byte(test.str), &info)
		if err != nil {
			t.Errorf("test %d -- err %v", i, err)
		}
		if !reflect.DeepEqual(info, test.obj) {
			t.Errorf("Test %d\nExpected:\n%v\nActual:\n%v\n", i, test.obj, info)
		}
	}
}
