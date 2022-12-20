grammar RW3C;
prog
    :	expr_list += run* EOF 
    ;
run
    :   (stmt | expr | 'break') ';'*
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
    |   chain_expr   
    ;

chain_expr
    :   atomic_expr (call_expr | access_prop_expr)*
    ;

atomic_expr
    :   fn_expr 
    |   lit_expr 
    |   scope 
    |   switch_expr
    |   if_expr 
    |   while_expr
    ;

stmt
    :   MULTI_COMMENT 
    |   SINGLE_COMMENT
    |   assign_stmt
    |   def_var_stmt
    |   def_type_stmt
    |   struct_stmt
    |   ret_stmt
    |   if_stmt
    |   while_stmt
    |   switch_stmt
    ;
if_stmt
    :   'if ' expr run ('else ' if_stmt)* ('else ' run)?
    ;
if_expr
    :   'if ' expr expr ('else ' if_expr)* ('else ' expr)?
    ;
switch_stmt
    :   'switch ' expr '{' (('case ' expr ':' run) | ('default' ':' run))* '}'
    ;
switch_expr
    :   'switch ' expr '{' (('case ' expr ':' expr) | ('default' ':' expr))* '}'
    ;
while_stmt
    :   'while ' expr run
    ;
while_expr
    :   'while ' expr expr
    ;
ret_stmt
    :   'ret' (' ' expr)?
    ;
def_var_stmt
    :   'let ' name_expr (':' '?'? type_expr) ('=' expr)?
    |   'mut ' name_expr (':' '?'? type_expr) '=' expr
    ;
def_type_stmt
    :   'type ' IDENTIF ('=' type_expr)?
    ;
struct_stmt
    :   'struct ' IDENTIF ('{' (IDENTIF ':' type_arg (',' IDENTIF ':' type_arg)* ','?)? '}')?
    ;
assign_stmt
    :   expr '=' expr 
    ;
call_expr
    :   '(' (expr (',' expr)* ','?)? ')' 
    ;
access_prop_expr
    :   '.' IDENTIF 
    |   '[' expr ']' 
    ;
fn_expr
    :   'fn ' IDENTIF? '(' (arg (',' arg)* ','?)? ')' (':' '?'? type_expr)? (run | scope)
    ;
scope
    :   '{' run* '}'
    ;
access_prop_type_expr
    :   '.' IDENTIF
    |   '[' type_expr ']'
    ;
type_expr
    :   (lit_expr | ('fn' '(' (type_arg (',' type_arg)* ','?)? ')' (':' '?'? type_expr)?)) access_prop_type_expr*
    ;
type_arg
    :   type_expr '?'?
    ;
name_expr
    :   IDENTIF
    ;
arg
    : name_expr (':' '?'? type_arg)? ('=' expr)?
    ;
lit_expr
    :   LONG 
    |   DOUBLE 
    |   STR 
    |   IDENTIF
    |   BOOL
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
MULTI_COMMENT 
    :   '/*' .*? '*/' -> channel(HIDDEN) 
    ;
SINGLE_COMMENT
    :   '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN) 
    ;