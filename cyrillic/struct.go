package cyrillic

type S1 struct {
	F1 string
	F2 *string
	F3 int
	F4 float64
	F5 S2
}

type S2 struct {
	F2 *string
	F3 S3
	F4 *S3
}

type S3 struct {
	Name string
}