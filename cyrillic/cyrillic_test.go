package cyrillic

import (
	"testing"
)

func TestFilter(t *testing.T) {
	s2Str := "strрусскийing for s2"
	s1Str := "strрусииущлалщвгыаing for s1"

	s2ExpStr := "string for s2"
	s1ExpStr := "string for s1"

	testTable := []struct {
		input    S1
		expected S1
	}{
		{input: S1{
			F1: "strinвцйгрвыграg for f1",
			F2: &s1Str,
			F3: 158,
			F4: 3.14,
			F5: S2{
				F2: &s2Str,
				F3: S3{
					Name: "NameOfS3",
				},
				F4: &S3{
					Name: "NameOfS3Ptrсруссимв",
				},
			},
		},
			expected: S1{
				F1: "string for f1",
				F2: &s1ExpStr,
				F3: 158,
				F4: 3.14,
				F5: S2{
					F2: &s2ExpStr,
					F3: S3{
						Name: "NameOfS3",
					},
					F4: &S3{
						Name: "NameOfS3Ptr",
					},
				},
			},
		},
	}

	for _, v := range testTable {
		Filter(&v.input)
		input := v.input
		expect := v.expected
		if input.F1 != expect.F1 {
			t.Errorf("Incorrect reuslt in F1")
		}
		if *input.F2 != *expect.F2 {
			t.Errorf("Incorrect reuslt in F2")
		}
		if *input.F5.F2 != *expect.F5.F2 {
			t.Errorf("Incorrect reuslt in F5, F2")
		}
		if input.F5.F3.Name != expect.F5.F3.Name {
			t.Errorf("Incorrect reuslt in F5, F3, Name")
		}
		if input.F5.F4.Name != expect.F5.F4.Name {
			t.Errorf("Incorrect reuslt in F5, F4, Name")
		}
	}

}
