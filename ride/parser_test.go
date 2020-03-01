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
	require.Equal(t, FuncE{
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
	require.Equal(t, FuncE{
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

func TestFuncPlus2(t *testing.T) {
	rs, err := Parse(`address + "a"`)
	require.NoError(t, err)
	require.Equal(t, NewFuncCall("+", "address", "a"), rs)

}

//func TestFuncHard(t *testing.T) {
//	rs, err := Parse("func some(index: Int,  answersCount: Int) = index + answersCount")
//	require.NoError(t, err)
//	require.Equal(t, NewFunc(
//		"some",
//		NewArgs(
//			NewArg("index", "Int"),
//			NewArg("answersCount", "Int")),
//		NewFuncCall("+", "index", "answersCount")), rs)
//}

func TestMatch(t *testing.T) {
	rs, err := Parse("match x {case y:Int => 555 }")
	require.NoError(t, err)
	require.Equal(t, NewMatch("x", TypedCase("y", "Int", NewNumber("555"))), rs)
}

var matchCase = `
  match getString(this) {
    case a: String => a
    case _ => address
  }`

func TestMatch2(t *testing.T) {
	rs, err := Parse(matchCase)

	require.NoError(t, err)
	require.Equal(t, NewMatch(
		NewFuncCall("getString", "this"),
		TypedCase("a", "String", "a"),
		UntypedCase("_", "address")), rs)
}

var fn = `
func getPreviousAnswer(address: String) = {
  match getString(this, address + "_a") {
    case a: String => a
    case _ => address
  }
}`

func TestMatch3(t *testing.T) {
	rs, err := Parse(fn)

	require.NoError(t, err)
	require.Equal(t, NewMatch(
		NewFuncCall("getString", "this"),
		TypedCase("a", "String", "a"),
		UntypedCase("_", "address")), rs)
}

func TestParseArray(t *testing.T) {
	p := `let answers = 
    ["It is certain.",
    "It is decidedly so."]`

	rs, err := Parse(p)

	require.NoError(t, err)
	require.Equal(t, 1, rs)
}

func TestParseFunc(t *testing.T) {
	p := `
	func getAnswer(question: String, previousAnswer: String) = {
    	let hash = sha256(toBytes(question + previousAnswer))
    	let index = toInt(hash)
    	answers[index % answersCount]
	}`

	rs, err := Parse(p)
	require.NoError(t, err)
	require.Equal(t,
		NewFunc("getAnswer",
			NewArgs(
				NewArg("question", "String"),
				NewArg("previousAnswer", "String")),
			nil), rs)
}

func TestLetBlock(t *testing.T) {
	p := `
    	let hash = 10
    	let index = "value"
    	5
	`

	rs, err := Parse(p)

	require.NoError(t, err)
	require.Equal(t,
		NewLetBlock("hash", NewNumber("10"),
			NewLetBlock("index", "value",
				NewNumber("5"))), rs)
}
