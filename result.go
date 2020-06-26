package platypus

// Result of handler.Process(ctx, cmd)
type Result struct {
	Out  string
	Leaf bool
}

func (res Result) String() string {
	return string(res.Out)
}

// Tail ...
func (res Result) Tail() bool {
	return res.Leaf
}
