package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (v Kadai1) ID() int {
	return v.id
}

func (v Kadai1) Name() string {
	return v.name
}

func (v *Kadai1) SetID(i int) {
	v.id = i
}

func (v *Kadai1) SetName(s string) {
	v.name = s
}
