package stringornumber

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

type CustomInt int64

func (ci *CustomInt) UnmarshalJSON(data []byte) error {
	var raw interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return errors.New("CustomInt: UnmarshalJSON: " + err.Error())
	}
	switch v := raw.(type) {
	case int:
		*ci = CustomInt(v)
		return nil
	case int32:
		*ci = CustomInt(v)
		return nil
	case int64:
		*ci = CustomInt(v)
		return nil
	case float64:
		*ci = CustomInt(v)
		return nil
	case string:
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return errors.New("CustomInt: parsing \"" + v + "\": not an int")
		}
		*ci = CustomInt(parsed)
		return nil
	default:
		return errors.New("CustomInt: parsing \"" + string(data) + "\": unknown value")
	}
}

func (ci *CustomInt) UnmarshalXML(d *xml.Decoder, se xml.StartElement) error {
	var raw string
	err := d.DecodeElement(&raw, &se)
	if err != nil {
		return errors.New("CustomInt: UnmarshalXML: " + err.Error())
	}
	raw = strings.ReplaceAll(raw, `"`, "")
	res, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return errors.New("CustomInt: parsing \"" + raw + "\": not an int")
	}
	*ci = CustomInt(res)
	return nil
}

type User struct {
	ID      CustomInt `json:"id" xml:"id"`
	Address Address   `json:"address" xml:"address"`
	Age     CustomInt `json:"age" xml:"age"`
}

type Users struct {
	Users []User `xml:"user"`
}

type Address struct {
	CityID CustomInt `json:"city_id" xml:"city_id"`
	Street string    `json:"street" xml:"street"`
}
