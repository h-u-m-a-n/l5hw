package main

import (
	"fmt"
	"github.com/h-u-m-a-n/l5hw/cyrillic"
)

func main() {
	s3 := cyrillic.S3{Name: "NameOfS3"}
	s3Ptr := &cyrillic.S3{Name: "NameOfS3Pаувшыоавыиаргыфиtr"}
	s2Str := "strрусскийing for s2"
	s2 := cyrillic.S2{
		F2: &s2Str,
		F3: s3,
		F4: s3Ptr,
	}
	s1Str := "string for s1"
	s1 := cyrillic.S1{
		F1: "сообще на ру сstring for f1",
		F2: &s1Str,
		F3: 158,
		F4: 3.14,
		F5: s2,
	}
	fmt.Println(s1)
	fmt.Println(s3Ptr)
	cyrillic.Filter(&s1)
	fmt.Println(s3Ptr)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s2Str)

}
