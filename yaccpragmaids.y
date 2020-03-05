%{
package yaccpragmaids


%}
%token RWIdl
%token RWLocal
%token RWRmi
%token RWDce
%token Name
%token Version
%token ClassName
%token HashCode
%token SerializationVersionUID
%token Uid

%union {
	Value string
}
%token <Value> Version Name
%start id_spec
%%

//(1)
id_spec:
	RWIdl ':' Name ':' Version
	{
	}
	|RWLocal ':' Name ':' Version
	{
	}
	|RWDce ':' Uid ':' Version
	{
	}
	|RWRmi ':' ClassName ':' HashCode
	{
	}
	|RWRmi ':' ClassName ':' HashCode  ':' SerializationVersionUID
	{
	}
%%



