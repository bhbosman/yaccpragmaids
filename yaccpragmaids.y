%{
package yaccpragmaids

import "strconv"

%}
%token RWIdl
%token RWLocal
%token RWRmi
%token RWDce
%token Value

%union {
	Value string
	Version IVersion
}


%type <Version> version
%type <Value> Value


%start id_spec
%%

id_spec:
	RWIdl ':' Value ':' version
	{
		if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
			v := NewIdlVersion($3, $5)
			setTypeVersion.SetTypeValue(v)
		}
	}
	|RWLocal ':' Value
	{
		if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
			v := NewLocalVersion($3)
			setTypeVersion.SetTypeValue(v)
		}
	}
	|RWDce ':' Value ':' version
	{
		if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
			v, e := NewDceVersion($3, $5)
			if e != nil{
				YaccPragmaIdslex.Error(e.Error())
				return 1
			}
			setTypeVersion.SetTypeValue(v)
		}
	}
	|RWRmi ':' Value ':' Value
	{
		if len($5) != 16 {
			YaccPragmaIdslex.Error("hash value incorrect length")
			return 1
		}
		intValue05, err := strconv.ParseUint($5, 16, 64)
		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
			v := NewRmiVersion($3, uint64(intValue05), 0)
			setTypeVersion.SetTypeValue(v)
		}
	}
	|RWRmi ':' Value ':' Value  ':' Value
	{
		if len($5) != 16 {
			YaccPragmaIdslex.Error("hash value incorrect length")
			return 1
		}
		intValue05, err := strconv.ParseUint($5, 16, 64)
		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		intValue07, err := strconv.ParseUint($7, 16, 64)
		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}

		if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
			v := NewRmiVersion($3, uint64(intValue05), intValue07)
			setTypeVersion.SetTypeValue(v)
		}
	}
version:
	Value
	{
    		intValue01, err := strconv.ParseUint($1, 10, 64)
    		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		$$ = NewVersion(intValue01, 0)
	}
	|Value '.'
	{
    		intValue01, err := strconv.ParseUint($1, 10, 64)
    		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		$$ = NewVersion(intValue01, 0)
	}
	|Value '.' Value
	{
    		intValue01, err := strconv.ParseUint($1, 10, 64)
    		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		intValue03, err := strconv.ParseUint($3, 10, 64)
    		if err != nil {
			YaccPragmaIdslex.Error(err.Error())
			return 1
		}
		$$ = NewVersion(intValue01, intValue03)
	}
%%



