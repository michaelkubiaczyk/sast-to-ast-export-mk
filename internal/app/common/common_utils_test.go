package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AstQueryIDTest struct {
	Language, Group, Name, Expected string
}

func TestGetAstQueryID(t *testing.T) {
	astQueryIDTests := []AstQueryIDTest{
		{"Kotlin", "Kotlin_High_Risk", "Code_Injection", "15158446363146771540"},
		{"CSharp", "General", "Find_SQL_Injection_Evasion_Attack", "8984835614866342550"},
		{"Go", "General", "Find_Command_Injection_Sanitize", "9498204717545098527"},
	}

	for _, test := range astQueryIDTests {
		result, err := GetAstQueryID(test.Language, test.Name, test.Group)
		assert.NoError(t, err)
		assert.Equal(t, test.Expected, result)
	}
}
