package stringornumber

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func TestCustomInt_UnmarshalJSON(t *testing.T) {
	var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 2,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

	expected := []User{
		{
			ID:      1,
			Address: Address{
				CityID: 5,
				Street: "Satbayev",
			},
			Age:     20,
		},
		{
			ID:      2,
			Address: Address{
				CityID: 6,
				Street: "Al-Farabi",
			},
			Age:     32,
		},
	}

	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		t.Fatal(err)
	}

	for i, user := range users {
		exp := expected[i]
		if user.ID != exp.ID {
			t.Fatal("Id doesn't match")
		}
		if user.Age != exp.Age {
			t.Fatal("Age doesn't match")
		}
		if user.Address.Street != exp.Address.Street {
			t.Fatal("Address.Street doesn't match")
		}
		if user.Address.CityID != exp.Address.CityID {
			t.Fatal("Address.CityID doesn't match")
		}
	}
}

func TestCustomInt_UnmarshalXML(t *testing.T) {
	var rawXML = []byte(`
<users>
	<user>
		<id>3</id>
		<address>
			<city_id>"7"</city_id>
			<street>Satbayev</street>
		</address>
		<age>45</age>
	</user>
	<user>
		<id>"4"</id>
		<address>
			<city_id>8</city_id>
			<street>Satbayev</street>
		</address>
		<age>"55"</age>
	</user>
</users>
`)

	expected := []User{
		{
			ID:      3,
			Address: Address{
				CityID: 7,
				Street: "Satbayev",
			},
			Age:     45,
		},
		{
			ID:      4,
			Address: Address{
				CityID: 8,
				Street: "Satbayev",
			},
			Age:     55,
		},
	}

	var users Users
	if err := xml.Unmarshal(rawXML, &users); err != nil {
		panic(err)
	}

	for i, user := range users.Users {
		exp := expected[i]
		if user.ID != exp.ID {
			t.Fatal("Id doesn't match")
		}
		if user.Age != exp.Age {
			t.Fatal("Age doesn't match")
		}
		if user.Address.Street != exp.Address.Street {
			t.Fatal("Address.Street doesn't match")
		}
		if user.Address.CityID != exp.Address.CityID {
			t.Fatal("Address.CityID doesn't match")
		}
	}


}
