package lexer_test

import (
	"github.com/FabianAlmos/almiconfig/lexer"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	empty    = ""
	testLine = "test"
	zeroRune = rune(0)

	requiredLine           = "required"
	reqEnvAccLine          = "required,env=ACCESS_SECRET"
	reqEnvAccSliceTypeLine = "required,env=ACCESS_SECRET,type=[,]string"
)

var (
	required           = []string{"required"}
	reqEnvAcc          = []string{"required", "env=ACCESS_SECRET"}
	reqEnvAccSliceType = []string{"required", "env=ACCESS_SECRET", "type=[,]string"}
)

func TestNewLexer(t *testing.T) {
	l := lexer.NewLexer(empty)
	assert.NotNil(t, l)
}

func TestLexer_HasNext(t *testing.T) {
	l := lexer.NewLexer(testLine)
	assert.True(t, l.HasNext())
}

func TestLexer_Next_Successful(t *testing.T) {
	l := lexer.NewLexer(testLine)
	assert.NotEqual(t, zeroRune, l.Next())
}

func TestLexer_Next_Fail(t *testing.T) {
	l := lexer.NewLexer(empty)
	assert.Equal(t, zeroRune, l.Next())
}

func TestLexer_Tokenize_SuccessfulLexConstraint(t *testing.T) {
	l := lexer.NewLexer(requiredLine)
	assert.Equal(t, required, l.Tokenize())
}

func TestLexer_Tokenize_SuccessfulLexMultipleConstraints(t *testing.T) {
	l := lexer.NewLexer(reqEnvAccLine)
	assert.Equal(t, reqEnvAcc, l.Tokenize())
}

func TestLexer_Tokenize_SuccessfulLexMultipleConstraints_NoSplitOnBracketComma(t *testing.T) {
	l := lexer.NewLexer(reqEnvAccSliceTypeLine)
	assert.Equal(t, reqEnvAccSliceType, l.Tokenize())
}
