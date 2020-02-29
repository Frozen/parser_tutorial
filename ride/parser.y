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
}

%token LexError
%token Let
%token Eq
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

let_definition: Let Literal Eq expression
{
	__yyfmt__.Println($4)
	$$ = NewLetBlock($2, $4)
}


//definition:

expression: Number
{
  __yyfmt__.Println("number", $1)
  $$ = NewNumber($1)
}
| String
  {
  __yyfmt__.Println("string", $1)
    $$ = StringAst($1)
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
