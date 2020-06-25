package platypus

// Command ...
type Command struct {
	Pattern string
}

// // Len of tokens
// func (cmd *Command) Len() int {
// 	return len(cmd.Tokens)
// }

// // Next returns the next token in the chain of command and it's index
// func (cmd *Command) Next(index int) (*Token, int) {
// 	if cmd.Len() < index {
// 		return cmd.Tokens[index], index
// 	}
// 	return nil, -1
// }

// // Tokenize ...
// func parse(raw string) *Command {
// 	var tokens []*Token

// 	values := strings.Split(raw, "*")

// 	for _, value := range values {
// 		var token *Token

// 		if value == "" {
// 			continue
// 		}

// 		if len(value) <= 3 {
// 			token = newToken(value)
// 			tokens = append(tokens, token)
// 		} else {
// 			token = newToken(value)
// 			tokens = append(tokens, token)
// 		}
// 	}
// 	return &Command{Tokens: tokens}
// }
