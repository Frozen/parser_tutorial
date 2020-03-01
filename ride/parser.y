%{
package pg

type pair struct {
  key string
  val interface{}
}

func setResult(l yyLexer, v Ast) {
  l.(*lex).result = v
}
%}

%union{
  obj map[string]interface{}
  list []interface{}
  pair pair
  val interface{}
  str string
  ast Ast
  defArgs []FuncArg
  callArgs []Ast
  cases []CaseE
  case_ CaseE
}

%token LexError
%token Let
%token Eq
%token Func
%token OpenB // (
%token CloseB // )
%token OpenF // {
%token CloseF // }
%token OpenS // [
%token CloseS // ]
%token Mod // %
%token Comma
%token Colon // :
%token Plus // +
%token Match // match
%token Case // case
%token RightArrow // =>
%token Compare // ==
%token <str> Number
%token <str> String Literal

//%type <str> object //members
//%type <pair> pair
//%type <val> array
//%type <list> elements
//%type <val> value

%type <ast> expression
%type <ast> definition_or_expression
%type <ast> definition
%type <ast> let_definition
%type <ast> func_definition
%type <ast> func_body
//%type <ast> func_optional_params
%type <defArgs> func_optional_params
%type <ast> func_call
%type <callArgs> func_call_args
%type <ast> match_expr
%type <case_> required_case
%type <cases> optional_cases
%type <case_> typed_case
%type <case_> untyped_case
%type <ast> simple_type
%type <cases> match_cases



%start start

%%

start: definition_or_expression
{
	setResult(yylex, $1)
}


definition_or_expression: definition
  {
    $$ = $1
  }
  | expression
  {
    $$ = $1
  }


definition: let_definition
{
//	setResult(yylex, $1)
	$$ = $1
}
| func_definition
{
	$$ = $1
}

let_definition: Let Literal Eq expression
{
	__yyfmt__.Println($4)
	$$ = NewLetBlock($2, $4)
}
//func_definition: Func Literal OpenB func_opt_params CloseB Eq func_body
func_definition: Func Literal OpenB func_optional_params CloseB Eq func_body
{
	$$ = NewFuncDeclaration($2, $4, $7)
}

func_optional_params: // empty
{
	$$ = []FuncArg{}
}
| Comma Literal Colon Literal func_optional_params
{
	$$ = append([]FuncArg{FuncArg {Name: $2, Type: $4}},   $5...)
}
| Literal Colon Literal func_optional_params
{
	$$ = append([]FuncArg{FuncArg {Name: $1, Type: $3}}, $4...)
}

func_body: OpenF definition_or_expression CloseF
{
	$$ = $2
} | expression
{
	$$ = $1
}


//definition:

expression:
func_call
{
	$$ = $1
}
| match_expr
{
  $$ = $1
}

| Number
{
  __yyfmt__.Println("number", $1)
  $$ = NewNumber($1)
}
| String
  {
  __yyfmt__.Println("string", $1)
    $$ = $1
  }
  | Literal
  {
  __yyfmt__.Println("literal", $1)
    $$ = $1
  }

match_expr: Match expression match_cases
{
    $$ = NewMatch($2, $3...)
}

match_cases: OpenF required_case optional_cases CloseF
{
	/// bla bla
    $$ = append([]CaseE{$2}, $3...)
}


required_case: typed_case
{
  $$ = $1
}
| untyped_case
{
  $$ = $1
}

optional_cases: /* empty */
{
  // optional_case
  $$ = []CaseE{}
}
| typed_case optional_cases
{
$$ = append([]CaseE{$1}, $2...)
}
| untyped_case optional_cases
{
$$ = append([]CaseE{$1}, $2...)
}

typed_case: Case Literal Colon Literal RightArrow expression
{
  $$ = TypedCase($2, $4, $6)
}

untyped_case: Case Literal RightArrow expression
{
  $$ = UntypedCase($2, $4)
}


func_call: Literal OpenB func_call_args CloseB
{
	$$ = NewFuncCall($1, $3...)
}
// index
  | Literal OpenS expression CloseS
  {
  	$$ = NewFuncCall("getByIndex", $1, $3)
  }
  // mod (x % y)
  | simple_type Mod simple_type
  {
    	$$ = NewFuncCall("%", $1, $3)
  }
  |simple_type Plus simple_type
  {
        $$ = NewFuncCall("+", $1, $3)
  }

simple_type: Number
{
  $$ = $1
}
| String
{
  $$ = $1
}
| Literal
{
  $$ = $1
}


func_call_args: /* empty */
{
	$$ = nil
}
| Comma expression func_call_args
{
	$$ = append([]Ast{$2}, $3...)
}
| expression func_call_args
{
	$$ = append([]Ast{$1}, $2...)
}

//| String

//
//members:
//  {
//    $$ = map[string]interface{}{}
//  }
//| pair
//  {
//    $$ = map[string]interface{}{
//      $1.key: $1.val,
//    }
//  }
//| members ',' pair
//  {
//    $1[$3.key] = $3.val
//    $$ = $1
//  }
//
//pair: String ':' value
//  {
//    $$ = pair{key: $1.(string), val: $3}
//  }
//
//array: '[' elements ']'
//  {
//    $$ = $2
//  }
//
//elements:
//  {
//    $$ = []interface{}{}
//  }
//| value
//  {
//    $$ = []interface{}{$1}
//  }
//| elements ',' value
//  {
//    $$ = append($1, $3)
//  }

//value:
//  String
//| Number
//| Literal
//| object
//  {
//    $$ = $1
//  }
//| array
