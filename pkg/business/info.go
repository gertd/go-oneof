package business

import (
	"encoding/json"
)

// Info -- common information of business object
type Info struct {
	Name        string      `json:"name"`
	AddressKind AddressKind `json:"addressKind,omitempty"`
	Address     Address     `json:"address"`
	Country     string      `json:"country"`
}

// UnmarshalJSON - -Info structure
func (i *Info) UnmarshalJSON(b []byte) error {

	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["name"], &i.Name)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["addressKind"], &i.AddressKind)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["country"], &i.Country)
	if err != nil {
		return err
	}

	switch i.AddressKind {
	case AddressKindHome:

		var home Home
		err = json.Unmarshal(*objMap["address"], &home)
		if err != nil {
			return err
		}
		i.Address = home

	case AddressKindOffice:

		var office Office
		err = json.Unmarshal(*objMap["address"], &office)
		if err != nil {
			return err
		}
		i.Address = office

	case AddressKindPostalBox:

		var postalBox PostalBox
		err = json.Unmarshal(*objMap["address"], &postalBox)
		if err != nil {
			return err
		}
		i.Address = postalBox

	}

	return nil
}

// Address -- Oneof (Home|Office|PostalBox)
type Address interface {
	isAddressType()
}

// Home -- Private home location address
type Home struct {
	HouseNumber string `json:"houseNumber"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
}

func (h Home) isAddressType() {}

// Office -- Legal office location address
type Office struct {
	OfficeNumber string `json:"officeNumber"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postalCode"`
}

func (o Office) isAddressType() {}

// PostalBox -- PO Box Address
type PostalBox struct {
	POBoxNumber string `json:"poBoxNumber"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
}

func (p PostalBox) isAddressType() {}
