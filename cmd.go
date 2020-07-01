package platypus

// Command ...
type Command struct {
	Phone   string
	Request string
}

// NewCommand creates a new instance of Command
func NewCommand(phone string, req string) *Command {
	return &Command{Phone: phone, Request: req}
}
