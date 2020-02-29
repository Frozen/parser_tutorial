package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse1(t *testing.T) {
	rs, err := Parse(`let x=  "fsas" `)
	//require.No
	t.Logf("%+v", rs)
	t.Log(err)
}

func TestParse2(t *testing.T) {
	rs, err := Parse(`let answersCount = 201`)
	t.Logf("%+v", rs)
	t.Log(err)
}

func TestFuncNoParams(t *testing.T) {
	rs, err := Parse(`func getAnswer()  = { 5 }`)
	t.Logf("%+v", rs)
	t.Log(err)
}

func TestFunc1Param(t *testing.T) {
	rs, err := Parse(`func getAnswer(x: Int)  = 5 `)
	require.NoError(t, err)
	require.Equal(t, FuncDeclaration{
		Name: "getAnswer",
		Args: NewArgs(FuncArg{
			Name: "x",
			Type: "Int",
		}),
		Body: NewNumber("5"),
	}, rs)
}

func TestFunc2Params(t *testing.T) {
	rs, err := Parse(`func getAnswer(x: Int, str: String)  = { 5 }`)
	require.NoError(t, err)
	require.Equal(t, FuncDeclaration{
		Name: "getAnswer",
		Args: NewArgs(NewArg("x", "Int"), NewArg("str", "String")),
		Body: NewNumber("5"),
	}, rs)
}

var example = `
func getAnswer(question: String, previousAnswer: String) = {
    let hash = sha256(toBytes(question + previousAnswer))
    let index = toInt(hash)
    answers[index % answersCount]
}`

func TestParseFuncCallNoArgs(t *testing.T) {
	rs, err := Parse("toBytes()")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("toBytes"), rs)
}

func TestParseFuncCall1Args(t *testing.T) {
	rs, err := Parse("toBytes(bla)")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("toBytes", "bla"), rs)
}

func TestParseFuncCall2Args(t *testing.T) {
	rs, err := Parse("toBytes(10, xxx)")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("toBytes", NewNumber("10"), "xxx"), rs)
}

func TestParseFuncCallCall(t *testing.T) {
	rs, err := Parse("sha256(toBytes(question))")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("sha256", NewFuncCall("toBytes", "question")), rs)
}

func TestFuncGetByIndex(t *testing.T) {
	rs, err := Parse("array[10]")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("getByIndex", "array", NewNumber("10")), rs)
}

func TestFuncMod(t *testing.T) {
	rs, err := Parse("index % answersCount")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("%", "index", "answersCount"), rs)
}

func TestFuncPlus(t *testing.T) {
	rs, err := Parse("index + answersCount")
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("+", "index", "answersCount"), rs)
}

func TestFuncHard(t *testing.T) {
	rs, err := Parse("func some(index: Int,  answersCount: Int) = index + answersCount")
	require.NoError(t, err)
	require.Equal(t, NewFuncDeclaration(
		"some",
		NewArgs(
			NewArg("index", "Int"),
			NewArg("answersCount", "Int")),
		NewFuncCall("+", "index", "answersCount")), rs)
}

var matchCase = `
  match getString(this, address + "_a") {
    case a: String => a
    case _ => address
  }`

func TestMatch(t *testing.T) {
	rs, err := Parse("match x {}")
	require.NoError(t, err)
	require.Equal(t, NewMatch("x", nil), rs)
}
