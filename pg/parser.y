%{
package pg

type pair struct {
  key string
  val interface{}
}

func setResult(l yyLexer, v string) {
  l.(*lex).result = v
}
%}

%union{
  obj map[string]interface{}
  list []interface{}
  pair pair
  val interface{}
  str string
}

%token LexError
%token <val> Number
%token <str> String Literal

%type <str> object //members
//%type <pair> pair
//%type <val> array
//%type <list> elements
//%type <val> value


%start object

%%

object: "SELECT" String
  {
    //$$ = $1
    setResult(yylex, $1)
  }
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
