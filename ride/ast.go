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

type StringAst string

type FuncArg struct {
	Name  string
	Value Ast
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

//func NewNumber(s string) Ast {
//	return NumberAst(s)
//}
