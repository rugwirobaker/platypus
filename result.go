package platypus

import "fmt"

var _ (Result) = (*noResult)(nil)

// Result of handler.Process(ctx, cmd)
type Result interface {
	//Print a string format of the Handler response
	fmt.Stringer

	// Tail returns true when the node responding
	// to the command is the last in the chain
	Tail() bool
}

type noResult string

func (res noResult) String() string {
	return string(res)
}

func (res noResult) Tail() bool {
	return true
}
