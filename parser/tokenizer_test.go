package parser

import "testing"

func TestTokenizeMessage(t *testing.T) {
	c := TokenizeMessage("")
	if c == nil {
		t.Error("Command was nil, failed")
	}
	c = TokenizeMessage("@login")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if len(c.args) > 0 {
		t.Errorf("No arguments provided, got %v", c.args)
	}

	c = TokenizeMessage("@login foo")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if len(c.args) != 1 {
		t.Errorf("One argument sent, got %v", c.args)
	}
	if c.argStr != "foo" {
		t.Errorf("Expected arguments 'foo', got '%s'", c.argStr)
	}

	c = TokenizeMessage("@login foo out")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if c.prep != "out" {
		t.Errorf("Preposition should be out, got %s", c.prep)
	}
	if len(c.args) != 2 {
		t.Errorf("Two arguments sent, got %v", c.args)
	}
	if c.argStr != "foo out" {
		t.Errorf("Expected arguments 'foo out', got '%s'", c.argStr)
	}

	c = TokenizeMessage("@login foo out bar")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if c.prep != "out" {
		t.Errorf("Preposition should be out, got %s", c.prep)
	}
	if c.iobj != "bar" {
		t.Errorf("Indirect object should be bar, got %s", c.iobj)
	}
	if len(c.args) != 3 {
		t.Errorf("Two arguments sent, got %v", c.args)
	}
	if c.argStr != "foo out bar" {
		t.Errorf("Expected arguments 'foo out bar', got '%s'", c.argStr)
	}

	c = TokenizeMessage("@login foo out of bar")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if c.prep != "out of" {
		t.Errorf("Preposition should be out of, got %s", c.prep)
	}
	if c.iobj != "bar" {
		t.Errorf("Indirect object should be bar, got %s", c.iobj)
	}
	if len(c.args) != 4 {
		t.Errorf("Two arguments sent, got %v", c.args)
	}
	if c.argStr != "foo out bar" {
		t.Errorf("Expected arguments 'foo out of bar', got '%s'", c.argStr)
	}

	c = TokenizeMessage("@login foo bar baz")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if c.prep != "" {
		t.Errorf("Preposition was not set, got %s", c.prep)
	}
	if c.iobj != "bar baz" {
		t.Errorf("Indirect object should be 'bar baz', got %s", c.iobj)
	}
	if len(c.args) != 3 {
		t.Errorf("Two arguments sent, got %v", c.args)
	}
	if c.argStr != "foo bar baz" {
		t.Errorf("Expected arguments 'foo bar baz', got '%s'", c.argStr)
	}

	c = TokenizeMessage("@login foo bar baz frobble")
	if c.verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.verb)
	}
	if c.dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.dobj)
	}
	if c.prep != "" {
		t.Errorf("Preposition was not set, got %s", c.prep)
	}
	if c.iobj != "bar baz frobble" {
		t.Errorf("Indirect object should be 'bar baz frobble', got %s", c.iobj)
	}
	if len(c.args) != 4 {
		t.Errorf("Two arguments sent, got %v", c.args)
	}
	if c.argStr != "foo bar baz" {
		t.Errorf("Expected arguments 'foo bar baz frobble', got '%s'", c.argStr)
	}
}
