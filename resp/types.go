package resp

type Value struct {
	Type  string
	Str   string
	Num   int
	Bulk  string
	Array []Value
}
