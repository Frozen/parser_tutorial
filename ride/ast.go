package pg

type Ast interface {
}

type LetBlock struct {
	Name  string
	Value Ast
	Body  Ast
}

func NewLetBlock(name string, value Ast, body Ast) LetBlock {
	return LetBlock{
		Name:  name,
		Value: value,
		Body:  body,
	}
}

type NumberAst string

func NewNumber(s string) Ast {
	return NumberAst(s)
}

//type StringAst string

type FuncArg struct {
	Name string
	Type string
}

type FuncE struct {
	Name     string
	Args     []FuncArg
	FuncBody Ast
	Body     Ast
}

func NewFunc(name string, Args []FuncArg, FuncBody Ast, Body Ast) FuncE {
	return FuncE{
		Name:     name,
		Args:     Args,
		FuncBody: FuncBody,
		Body:     Body,
	}
}

func NewArgs(args ...FuncArg) []FuncArg {
	return args
}

func NewArg(name, Type string) FuncArg {
	return FuncArg{
		Name: name,
		Type: Type,
	}
}

type FuncCall struct {
	Name string
	Args []Ast
}

func NewFuncCall(name string, args ...Ast) FuncCall {
	return FuncCall{
		Name: name,
		Args: args,
	}
}

type CaseE struct {
	Typed    bool
	VarName  string
	TypeName string
	Body     Ast
}

func TypedCase(vaName string, Typename string, Body Ast) CaseE {
	return CaseE{
		Typed:    true,
		VarName:  vaName,
		TypeName: Typename,
		Body:     Body,
	}
}

func UntypedCase(name string, body Ast) CaseE {
	return CaseE{
		Typed:    false,
		VarName:  name,
		TypeName: "",
		Body:     body,
	}
}

type MatchE struct {
	Arg   Ast
	Cases []CaseE
}

func NewMatch(Ast Ast, Cases ...CaseE) MatchE {
	return MatchE{
		Arg:   Ast,
		Cases: Cases,
	}
}

//func NewNumber(s string) Ast {
//	return NumberAst(s)
//}
