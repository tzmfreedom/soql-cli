grammar dml;

statement
    : insertStatement
    | updateStatement
    | deleteStatement
    ;

insertStatement
    : INSERT sobject '(' field (',' field)* ')' VALUES '(' literal (',' literal)* ')'
    ;

updateStatement
    : UPDATE sobject SET field '=' literal (',' field '=' literal)* WhereClause?
    ;

deleteStatement
    : DELETE FROM sobject WhereClause?
    ;

sobject
    : Identifier
    ;

field
    : Identifier
    ;

literal
    :   IntegerLiteral
    |   FloatingPointLiteral
    |   StringLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;

QUOTE         : '\'';
FROM          : F R O M;
INSERT        : I N S E R T;
UPDATE        : U P D A T E;
DELETE        : D E L E T E;
WHERE         : W H E R E;
ON            : O N;
VALUES        : V A L U E S;
SET           : S E T;

StringLiteral
    :   QUOTE StringCharacters? QUOTE
    ;

fragment
StringCharacters
    :   StringCharacter+
    ;

fragment
StringCharacter
    :   ~['\\]
    ;

BooleanLiteral
    :   T R U E
    |   F A L S E
    ;

NullLiteral :   N U L L;

Identifier
    :   JavaLetter JavaLetterOrDigit*
    ;

fragment
JavaLetter
    :   [a-zA-Z$_] // these are the "java letters" below 0xFF
    ;

fragment
JavaLetterOrDigit
    :   [a-zA-Z0-9$_] // these are the "java letters or digits" below 0xFF
    ;

IntegerLiteral
    :   DecimalIntegerLiteral
    ;

FloatingPointLiteral
    :   Digits '.' Digits?
    ;

fragment
DecimalIntegerLiteral
    :   DecimalNumeral
    ;

fragment
IntegerTypeSuffix
    :   [lL]
    ;

fragment
DecimalNumeral
    :   Digits
    ;

fragment
Digits
    :   Digit (Digit)?
    ;

fragment
Digit
    :   '0'
    |   NonZeroDigit
    ;

fragment
NonZeroDigit
    :   [1-9]
    ;

WS  :  [ \t\r\n\u000C]+ -> skip
    ;

WhereClause
    : WHERE [ \t\r\n=!<>a-zA-Z0-9_]+
    ;

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
fragment SPACE : ' ';