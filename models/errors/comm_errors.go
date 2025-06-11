package errs

import "errors"

var ErrInvalidRecipient = errors.New("Invalid recipient")
var ErrConversationNotFound = errors.New("Conversation not found")
