package network

import (
	"mud/parser"
	"strings"
)


type Location struct {
	Id	string
	Name	string
	Short	string
	Long	string
	Details map[string]string
	//Mobs	[]Mobile
	//Inventory []Thing
	Verbs	map[string]LocationVerb
}

type LocationVerb func(location *Location, client *Client, command *parser.Command) (errorMessage string)

func (loc *Location) ProcessCommand(client *Client, command *parser.Command) (errorMessage string) {
	if command.Verb == "" {
		errorMessage = "No command entered."
	}

	// test against the client player obj

	// is there a dobj in the command?
	// test against the matching dobj

	// is there an iobj in the command?
	// test againt the matching iobj

	// test against the location
	for verb, fVerb := range loc.Verbs {
		if verbMatch(verb, command.Verb) {
			errorMessage = fVerb(loc, client, command)
			break
		}
	}

	// test against each mob

	// test against each thing

	return
}

func verbMatch(verb string, spec string) bool {
	// TODO resolve abbreviated verbs when more than one abbreviation may match
	// will need to pull all the verbs off the location and Things and Mobs and
	// look for least-common matches

	star := strings.Index(verb, "*")
	if star > 1 {
		// build a list of possible matches
		for idx := star; idx > len(verb); idx++ {
			v := verb[0:idx]
			if v == spec {
				return true
			}
		}
	}

	return false
}