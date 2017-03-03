package parser

import "testing"

func TestTokenizeMessage(t *testing.T) {
	c := TokenizeMessage("")
	if c == nil {
		t.Error("Command was nil, failed")
	}
	c = TokenizeMessage("@login")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if len(c.Args) > 0 {
		t.Errorf("No arguments provided, got %v", c.Args)
	}

	c = TokenizeMessage("@login foo")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if len(c.Args) != 1 {
		t.Errorf("One argument sent, got %v", c.Args)
	}
	if c.ArgStr != "foo" {
		t.Errorf("Expected arguments 'foo', got '%s'", c.ArgStr)
	}

	c = TokenizeMessage("@login foo on")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if c.Prep != "on" {
		t.Errorf("Preposition should be on, got %s", c.Prep)
	}
	if len(c.Args) != 2 {
		t.Errorf("Two arguments sent, got %v", c.Args)
	}
	if c.ArgStr != "foo on" {
		t.Errorf("Expected arguments 'foo on', got '%s'", c.ArgStr)
	}

	c = TokenizeMessage("@login foo on bar")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if c.Prep != "on" {
		t.Errorf("Preposition should be on, got %s", c.Prep)
	}
	if c.Iobj != "bar" {
		t.Errorf("Indirect object should be bar, got %s", c.Iobj)
	}
	if len(c.Args) != 3 {
		t.Errorf("Two arguments sent, got %v", c.Args)
	}
	if c.ArgStr != "foo on bar" {
		t.Errorf("Expected arguments 'foo on bar', got '%s'", c.ArgStr)
	}

	c = TokenizeMessage("@login foo out of bar")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if c.Prep != "out of" {
		t.Errorf("Preposition should be out of, got %s", c.Prep)
	}
	if c.Iobj != "bar" {
		t.Errorf("Indirect object should be bar, got %s", c.Iobj)
	}
	if len(c.Args) != 4 {
		t.Errorf("Two arguments sent, got %v", c.Args)
	}
	if c.ArgStr != "foo out of bar" {
		t.Errorf("Expected arguments 'foo out of bar', got '%s'", c.ArgStr)
	}

	c = TokenizeMessage("@login foo bar baz")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if c.Prep != "" {
		t.Errorf("Preposition was not set, got %s", c.Prep)
	}
	if c.Iobj != "bar baz" {
		t.Errorf("Indirect object should be 'bar baz', got %s", c.Iobj)
	}
	if len(c.Args) != 3 {
		t.Errorf("Two arguments sent, got %v", c.Args)
	}
	if c.ArgStr != "foo bar baz" {
		t.Errorf("Expected arguments 'foo bar baz', got '%s'", c.ArgStr)
	}

	c = TokenizeMessage("@login foo bar baz frobble")
	if c.Verb != "@login" {
		t.Errorf("Expected @login for verb, got %s", c.Verb)
	}
	if c.Dobj != "foo" {
		t.Error("Direct object should be foo, got %s", c.Dobj)
	}
	if c.Prep != "" {
		t.Errorf("Preposition was not set, got %s", c.Prep)
	}
	if c.Iobj != "bar baz frobble" {
		t.Errorf("Indirect object should be 'bar baz frobble', got %s", c.Iobj)
	}
	if len(c.Args) != 4 {
		t.Errorf("Two arguments sent, got %v", c.Args)
	}
	if c.ArgStr != "foo bar baz frobble" {
		t.Errorf("Expected arguments 'foo bar baz frobble', got '%s'", c.ArgStr)
	}
}
