package shiro

import (
	"testing"
)

var (
	p  Permission = "edit log,viewLog , modify log,archive log:edit,create,list"
	r1 Permission = "modify log:list"
)

func TestTokenize(t *testing.T) {
	tokens := p.tokenize()
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

func TestImplies(t *testing.T) {
	if p.Implies(r1) {
		t.Log("User is permitted")
	}
}
