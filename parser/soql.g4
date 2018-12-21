grammar soql;

query
    : selectClause
      fromClause
      whereClause?
      withClause?
      groupClause?
      orderClause?
      limitClause?
      offsetClause?
      viewClause?
    ;

selectClause
    : SELECT fieldList
    ;

fieldList
    : selectField (COMMA selectField)*
    ;

selectField
    : soqlField
    | subquery
    | TYPEOF soqlField
      (WHEN apexIdentifier THEN fieldList)+
      ELSE fieldList
      END
    ;

fromClause
    : FROM apexIdentifier (USING SCOPE filterScope)?
    ;

filterScope
    :
    ;

soqlField
    : (apexIdentifier DOT)* apexIdentifier  # SoqlFieldReference
    | apexIdentifier LPAREN (soqlField (COMMA soqlField)*)? RPAREN # SoqlFunctionCall
    ;

subquery
    : query
    ;

whereClause
    : WHERE whereFields
    ;

whereFields
    : whereField
    | whereFields and_or=(SOQL_AND|SOQL_OR) whereFields
    ;

whereField
    :
       SOQL_NOT?
       soqlField
       op=(
         '='
         | '<'
         | '>'
         | '<='
         | '>='
         | '!='
         | '<>'
         | LIKE
         | IN
       )
       soqlValue
    |  '(' whereFields ')'
    ;

limitClause
    :  LIMIT IntegerLiteral
    ;

orderClause
    :  ORDER BY soqlField (',' soqlField)* asc_desc=(ASC | DESC)? (NULLS nulls=(LAST | FIRST))?
    ;

soqlValue
    :  literal
    |  apexIdentifier COLON literal
    ;

withClause
    :  WITH DATA CATEGORY soqlFilteringExpression
    ;

soqlFilteringExpression
    :
    ;

groupClause
    :  GROUP BY soqlField (',' soqlField)* (HAVING havingConditionExpression)?
    ;

havingConditionExpression
    :
    ;

offsetClause
    :  OFFSET IntegerLiteral
    ;

viewClause
    : FOR (VIEW | REFERENCE) (UPDATE (TRACKING | VIEWSTAT))?
    ;

// Apex - SOSL literal

soslLiteral
    : '[' soslQuery ']'
	;

soslQuery
    : FIND literal IN ALL FIELDS RETURNING soslReturningObject (',' soslReturningObject)*
    ;

soslReturningObject
    : Identifier ('(' Identifier (',' Identifier)* ')')?
    ;
apexIdentifier
    :  Identifier
    |  DATA
    |  GROUP
    |  SCOPE
    |  CATEGORY
    |  REFERENCE
    |  OFFSET
    |  THEN
    |  FIND
    |  RETURNING
    |  ALL
    ;

literal
    :   IntegerLiteral
    |   FloatingPointLiteral
    |   StringLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;

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

QUOTE         : '\'';
LPAREN        : '(';
RPAREN        : ')';
WHEN          : W H E N;
ELSE          : E L S E;
COMMA         : C O M M A;
DOT           : D O T;
FOR           : F O R;
UPDATE        : U P D A T E;
COLON         : C O L O N;
SELECT        : S E L E C T;
FROM          : F R O M;
WHERE         : W H E R E;
LIMIT         : L I M I T;
ORDER         : O R D E R;
BY            : B Y;
ASC           : A S C;
DESC          : D E S C;
WITH          : W I T H;
TYPEOF        : T Y P E O F;
REFERENCE     : R E F E R E N C E;
VIEW          : V I E W;
VIEWSTAT      : V I E W S T A T;
TRACKING      : T R A C K I N G;
OFFSET        : O F F S E T;
IN            : I N;
END           : E N D;
USING         : U S I N G;
DATA          : D A T A;
CATEGORY      : C A T E G O R Y;
GROUP         : G R O U P;
HAVING        : H A V I N G;
NULLS         : N U L L S;
FIRST         : F I R S T;
LAST          : L A S T;
SCOPE         : S C O P E;
ROLLUP        : R O L L U P;
CUBE          : C U B E;
LIKE          : L I K E;
THEN          : T H E N;
SOQL_AND   : A N D;
SOQL_OR    : O R;
SOQL_NOT   : N O T;
FIND       : F I N D;
FIELDS     : F I E L D S;
RETURNING  : R E T U R N I N G;
ALL        : A L L;

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