package pg

type Ast interface {
}

type LetBlock struct {
	Name string
	Body Ast
}

func NewLetBlock(name string, body Ast) LetBlock {
	return LetBlock{
		Name: name,
		Body: body,
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

type FuncDeclaration struct {
	Name string
	Args []FuncArg
	Body Ast
}

func NewFuncDeclaration(name string, Args []FuncArg, Body Ast) FuncDeclaration {
	return FuncDeclaration{
		Name: name,
		Args: Args,
		Body: Body,
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
	VarName  string
	TypeName string
	Body     Ast
}

func NewCase(vaName string, Typename string, Body Ast) CaseE {
	return CaseE{
		VarName:  vaName,
		TypeName: Typename,
		Body:     Body,
	}
}

type MatchE struct {
	Arg   Ast
	Cases []CaseE
}

func NewMatch(Ast Ast, Cases []CaseE) MatchE {
	return MatchE{
		Arg:   Ast,
		Cases: Cases,
	}
}

//func NewNumber(s string) Ast {
//	return NumberAst(s)
//}
