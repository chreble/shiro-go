package shiro

import (
	"testing"
	"fmt"
)

var (
	p  Permission = "edit log,viewLog , modify log,archive log:edit,create,list"
	p2 Permission = ""
	r1 Permission = "modify log:list"
)

// This test will check for correct no of tokens and subtokens in the 
// permission string
func TestTokenize(t *testing.T) {
	tokens, err := p.tokenize()
	if err != nil {
		Log.Println(err)
	}
	if len(tokens) == 2 {
		t.Log("Correct no of tokens generated")
	}
	if len(tokens[0]) == 4 {
		t.Log("Correct no of subtoken in 1st token")
	}
	if len(tokens[1]) == 3 {
		t.Log("Correct no of subtoken in 2nd token")
	}
}

// This test will check for empty string
func TestTokenize1(t *testing.T) {
	tokens, err := p2.tokenize()
	if err != nil {
		Log.Println(err)
	}
	fmt.Printf("%v", tokens)
	if len(tokens) == 0 {
		t.Log("No permissions available")
	}
}
func TestImplies(t *testing.T) {
	if p2.Implies(r1) {
		t.Log("User is permitted")
	}
}
