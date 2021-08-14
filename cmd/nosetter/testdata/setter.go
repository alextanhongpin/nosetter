package testdata

type Foo struct {
	name string
}

func TestSetData() {
	a := 1
	_ = a

	var f Foo
	f.name = "john"
}
