package parser

import "strings"

func TokenizeMessage(msg string) (command *Command) {
	command = &Command{}

	msg = strings.TrimSpace(msg)

	// TODO decide if we need to treat quoted sections as single words

	if msg != "" {
		// initially tokenize by spaces
		msgParts := strings.Split(msg, " ")

		// identify the verb (always the first word)
		command.Verb = msgParts[0]

		// populate the raw args
		if len(msgParts) > 1 {
			command.Args = msgParts[1:]
			command.ArgStr = strings.Join(command.Args, " ")

			// populate the parameters for the verb
			// the pattern is always dobj prep iobj
			command.Dobj = msgParts[1]

			if len(msgParts) > 2 {
				// is the next part a preposition?
				prepStr := msgParts[2]
				iobjIndex := 2
				if len(msgParts) > 3 {
					doublePrepStr := strings.Join(msgParts[2:4], " ")

					if wordInList(DOUBLE_WORD_PREPOSITIONS, doublePrepStr) {
						command.Prep = doublePrepStr
						iobjIndex = 4
					}
				}
				if command.Prep == "" {
					if wordInList(SINGLE_WORD_PREPOSITIONS, prepStr) {
						command.Prep = prepStr
						iobjIndex = 3
					}
				}

				if len(msgParts) > iobjIndex {
					command.Iobj = strings.Join(msgParts[iobjIndex:], " ")
				}
			}
		}
	}

	return
}

func wordInList(list []string, target string) bool {
	for _, word := range list {
		if target == word {
			return true
		}
	}

	return false
}