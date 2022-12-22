grammar RW3C;
prog
    :	exprList += run* EOF 
    ;
run
    :   (stmt | expr) ';'*
    ;
whileRun
    :   (run | 'break')*
    ;
whileExprBody
    :   (expr | 'break')*
    ;
expr
    :   '(' expr ')'
    |   expr ('*'|'/') expr
    |   expr ('+'|'-') expr
    |   expr ('>'|'<'|'<='|'>='|'=='|'!=') expr
    |   expr ('&&'|'||') expr
    |   ('++'|'--') expr
    |   ('+'|'-') expr
    |   '!' expr
    |   expr ('++'|'--')
    |   chainExpr
    |   atomicExpr
    ;

atomicExpr
    :   fnExpr 
    |   litExpr 
    |   scope 
    |   switchExpr
    |   ifExpr 
    |   whileExpr
    ;

chainExpr
    :   atomicExpr (callExpr | accessPropExpr)
    ;
callExpr
    :   '(' (expr (',' expr)* ','?)? ')' 
    ;
accessPropExpr
    :   '.' IDENTIF 
    |   '[' expr ']' 
    ;

stmt
    :   MULTI_COMMENT 
    |   SINGLE_COMMENT
    |   assignStmt
    |   defVarStmt
    |   defTypeStmt
    |   structStmt
    |   retStmt
    |   ifStmt
    |   whileStmt
    |   switchStmt
    ;
ifStmt
    :   'if ' expr run ('else ' ifStmt)* ('else ' run)?
    ;
ifExpr
    :   'if ' expr expr ('else ' ifExpr)* ('else ' expr)?
    ;
switchStmt
    :   'switch ' expr '{' (('case ' expr ':' run) | ('default' ':' run))* '}'
    ;
switchExpr
    :   'switch ' expr '{' (('case ' expr ':' expr) | ('default' ':' expr))* '}'
    ;
whileStmt
    :   'while ' expr whileRun
    ;
whileExpr
    :   'while ' expr whileExprBody
    ;
retStmt
    :   'ret' (' ' expr)?
    ;
defVarStmt
    :   'let ' nameExpr (':' '?'? typeExpr) ('=' expr)?
    |   'mut ' nameExpr (':' '?'? typeExpr) '=' expr
    ;
defTypeStmt
    :   'type ' IDENTIF ('=' typeExpr)?
    ;
structStmt
    :   'struct ' IDENTIF ('{' (IDENTIF ':' typeArg (',' IDENTIF ':' typeArg)* ','?)? '}')?
    ;
assignStmt
    :   expr '=' expr 
    ;
fnExpr
    :   'fn ' IDENTIF? '(' (arg (',' arg)* ','?)? ')' (':' '?'? typeExpr)? (run | scope)
    ;
scope
    :   '{' run* '}'
    ;
accessPropTypeExpr
    :   '.' IDENTIF
    |   '[' typeExpr ']'
    ;
typeExpr
    :   (litExpr | ('fn' '(' (typeArg (',' typeArg)* ','?)? ')' (':' '?'? typeExpr)?)) accessPropTypeExpr*
    ;
typeArg
    :   typeExpr '?'?
    ;
nameExpr
    :   IDENTIF
    ;
arg
    : nameExpr (':' '?'? typeArg)? ('=' expr)?
    ;
litExpr
    :   LONG 
    |   DOUBLE 
    |   STR 
    |   IDENTIF
    |   BOOL
    |   NULL
    ;

RESERVED
    : 'type'
    | 'let'
    | 'mut'
    | 'struct'
    | 'fn'
    | 'if'
    | 'switch'
    | 'ret'
    ;
IDENTIF_PART
    :   IDENTIF_START
    |   [0-9]
    ;
IDENTIF_START
    :   [A-Z]
    |   [a-z]
    |   '_'
    |   '$'
    ;
IDENTIF 
    :   IDENTIF_START IDENTIF_PART* 
    ;
NEWLINE 
    :   [\r\n]+ -> skip 
    ;
fragment NUM
    :   NUM_START NUM_BODY*
    ;
fragment NUM_START
    :   [0-9]
    ;
fragment NUM_BODY
    :   NUM_START
    |   '_'
    ;
LONG
    :   NUM+ 
    ;
DOUBLE   
    :   (NUM*.NUM+) | (NUM+.NUM*) 
    ;
STR     
    :   '"' ( ESC | ~[\\"] )* '"' 
    ;
ESC     
    :   '\\"' 
    |   '\\\\' 
    ;
BOOL
    :   'true'
    |   'false'
    ;
NULL
    :   'null'
    ;
MULTI_COMMENT 
    :   '/*' .*? '*/' -> channel(HIDDEN) 
    ;
SINGLE_COMMENT
    :   '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN) 
    ;