// Code generated by gocc; DO NOT EDIT.

package parser



type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Proto	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Proto : Syntax	<< X[0], nil >>`,
		Id:         "Proto",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Proto : Syntax ProtoContents	<< X[0], nil >>`,
		Id:         "Proto",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContents : ProtoContent	<< X[0], nil >>`,
		Id:         "ProtoContents",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContents : ProtoContents ProtoContent	<< X[0], nil >>`,
		Id:         "ProtoContents",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContent : Import	<< X[0], nil >>`,
		Id:         "ProtoContent",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContent : Package	<< X[0], nil >>`,
		Id:         "ProtoContent",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContent : Option	<< X[0], nil >>`,
		Id:         "ProtoContent",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContent : TopLevelDef	<< X[0], nil >>`,
		Id:         "ProtoContent",
		NTType:     3,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ProtoContent : emptyStatement	<< X[0], nil >>`,
		Id:         "ProtoContent",
		NTType:     3,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Syntax : "syntax" "=" quote "proto3" quote ";"	<< X[0], nil >>`,
		Id:         "Syntax",
		NTType:     4,
		Index:      10,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Import : "import" ImportState strLit ";"	<< X[0], nil >>`,
		Id:         "Import",
		NTType:     5,
		Index:      11,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ImportState : "weak"	<< X[0], nil >>`,
		Id:         "ImportState",
		NTType:     6,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ImportState : "public"	<< X[0], nil >>`,
		Id:         "ImportState",
		NTType:     6,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Package : "package" FullIdent ";"	<< X[0], nil >>`,
		Id:         "Package",
		NTType:     7,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Option : "option" OptionName "=" Constant ";"	<< X[0], nil >>`,
		Id:         "Option",
		NTType:     8,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OptionName : ident	<< X[0], nil >>`,
		Id:         "OptionName",
		NTType:     9,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OptionName : ident DotRepeatedIdents	<< X[0], nil >>`,
		Id:         "OptionName",
		NTType:     9,
		Index:      17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OptionName : "(" FullIdent ")"	<< X[0], nil >>`,
		Id:         "OptionName",
		NTType:     9,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OptionName : "(" FullIdent ")" DotRepeatedIdents	<< X[0], nil >>`,
		Id:         "OptionName",
		NTType:     9,
		Index:      19,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DotRepeatedIdents : "." ident	<< X[0], nil >>`,
		Id:         "DotRepeatedIdents",
		NTType:     10,
		Index:      20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DotRepeatedIdents : DotRepeatedIdents "." ident	<< X[0], nil >>`,
		Id:         "DotRepeatedIdents",
		NTType:     10,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDef : Message	<< X[0], nil >>`,
		Id:         "TopLevelDef",
		NTType:     11,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDef : Enum	<< X[0], nil >>`,
		Id:         "TopLevelDef",
		NTType:     11,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TopLevelDef : Service	<< X[0], nil >>`,
		Id:         "TopLevelDef",
		NTType:     11,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "double"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "float"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "int32"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "int64"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "uint32"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "uint64"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "sint32"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "sint64"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      32,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "fixed32"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "fixed64"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "sfixed32"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "sfixed64"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "bool"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "string"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : "bytes"	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : MessageType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Type : EnumType	<< X[0], nil >>`,
		Id:         "Type",
		NTType:     12,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldNumber : intLit	<< X[0], nil >>`,
		Id:         "FieldNumber",
		NTType:     13,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Field : Type FieldName "=" FieldNumber ";"	<< X[0], nil >>`,
		Id:         "Field",
		NTType:     14,
		Index:      43,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Field : Type FieldName "=" FieldNumber "[" FieldOptions "]" ";"	<< X[0], nil >>`,
		Id:         "Field",
		NTType:     14,
		Index:      44,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Field : "repeated" Type FieldName "=" FieldNumber ";"	<< X[0], nil >>`,
		Id:         "Field",
		NTType:     14,
		Index:      45,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Field : "repeated" Type FieldName "=" FieldNumber "[" FieldOptions "]" ";"	<< X[0], nil >>`,
		Id:         "Field",
		NTType:     14,
		Index:      46,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldOptions : FieldOption	<< X[0], nil >>`,
		Id:         "FieldOptions",
		NTType:     15,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldOptions : FieldOptions "," FieldOption	<< X[0], nil >>`,
		Id:         "FieldOptions",
		NTType:     15,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldOption : OptionName "=" Constant	<< X[0], nil >>`,
		Id:         "FieldOption",
		NTType:     16,
		Index:      49,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Oneof : "oneof" OneofName "{" "}"	<< X[0], nil >>`,
		Id:         "Oneof",
		NTType:     17,
		Index:      50,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Oneof : "oneof" OneofName "{" RepeatedOneOfFieldOrEmpty "}"	<< X[0], nil >>`,
		Id:         "Oneof",
		NTType:     17,
		Index:      51,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RepeatedOneOfFieldOrEmpty : OneofField	<< X[0], nil >>`,
		Id:         "RepeatedOneOfFieldOrEmpty",
		NTType:     18,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RepeatedOneOfFieldOrEmpty : RepeatedOneOfFieldOrEmpty OneofField	<< X[0], nil >>`,
		Id:         "RepeatedOneOfFieldOrEmpty",
		NTType:     18,
		Index:      53,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RepeatedOneOfFieldOrEmpty : emptyStatement	<< X[0], nil >>`,
		Id:         "RepeatedOneOfFieldOrEmpty",
		NTType:     18,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RepeatedOneOfFieldOrEmpty : RepeatedOneOfFieldOrEmpty emptyStatement	<< X[0], nil >>`,
		Id:         "RepeatedOneOfFieldOrEmpty",
		NTType:     18,
		Index:      55,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OneofField : Type FieldName "=" FieldNumber ";"	<< X[0], nil >>`,
		Id:         "OneofField",
		NTType:     19,
		Index:      56,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OneofField : Type FieldName "=" FieldNumber "[" FieldOptions "]" ";"	<< X[0], nil >>`,
		Id:         "OneofField",
		NTType:     19,
		Index:      57,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MapField : "map" "<" KeyType "," Type ">" MapName "=" FieldNumber ";"	<< X[0], nil >>`,
		Id:         "MapField",
		NTType:     20,
		Index:      58,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MapField : "map" "<" KeyType "," Type ">" MapName "=" FieldNumber "[" FieldOptions "]" ";"	<< X[0], nil >>`,
		Id:         "MapField",
		NTType:     20,
		Index:      59,
		NumSymbols: 13,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "int32"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "int64"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "uint32"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      62,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "uint64"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "sint32"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "sint64"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "fixed32"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "fixed64"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      67,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "sfixed32"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "sfixed64"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "bool"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `KeyType : "string"	<< X[0], nil >>`,
		Id:         "KeyType",
		NTType:     21,
		Index:      71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Reserved : "reserved" Ranges ";"	<< X[0], nil >>`,
		Id:         "Reserved",
		NTType:     22,
		Index:      72,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Reserved : "reserved" FieldNames ";"	<< X[0], nil >>`,
		Id:         "Reserved",
		NTType:     22,
		Index:      73,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Ranges : Range	<< X[0], nil >>`,
		Id:         "Ranges",
		NTType:     23,
		Index:      74,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Ranges : Ranges "," Range	<< X[0], nil >>`,
		Id:         "Ranges",
		NTType:     23,
		Index:      75,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Range : intLit	<< X[0], nil >>`,
		Id:         "Range",
		NTType:     24,
		Index:      76,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Range : intLit "to" intLit	<< X[0], nil >>`,
		Id:         "Range",
		NTType:     24,
		Index:      77,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Range : intLit "to" "max"	<< X[0], nil >>`,
		Id:         "Range",
		NTType:     24,
		Index:      78,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldNames : FieldName	<< X[0], nil >>`,
		Id:         "FieldNames",
		NTType:     25,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldNames : FieldNames "," FieldName	<< X[0], nil >>`,
		Id:         "FieldNames",
		NTType:     25,
		Index:      80,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Enum : "enum" EnumName EnumBody	<< X[0], nil >>`,
		Id:         "Enum",
		NTType:     26,
		Index:      81,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumBody : "{" "}"	<< X[0], nil >>`,
		Id:         "EnumBody",
		NTType:     27,
		Index:      82,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumBody : "{" Option "}"	<< X[0], nil >>`,
		Id:         "EnumBody",
		NTType:     27,
		Index:      83,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumBody : "{" EnumField "}"	<< X[0], nil >>`,
		Id:         "EnumBody",
		NTType:     27,
		Index:      84,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumBody : "{" emptyStatement "}"	<< X[0], nil >>`,
		Id:         "EnumBody",
		NTType:     27,
		Index:      85,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumField : ident "=" intLit ";"	<< X[0], nil >>`,
		Id:         "EnumField",
		NTType:     28,
		Index:      86,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumField : ident "=" intLit "[" EnumValueOptions "]" ";"	<< X[0], nil >>`,
		Id:         "EnumField",
		NTType:     28,
		Index:      87,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumValueOptions : EnumValueOption	<< X[0], nil >>`,
		Id:         "EnumValueOptions",
		NTType:     29,
		Index:      88,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumValueOptions : EnumValueOptions "," EnumValueOption	<< X[0], nil >>`,
		Id:         "EnumValueOptions",
		NTType:     29,
		Index:      89,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumValueOption : OptionName "=" Constant	<< X[0], nil >>`,
		Id:         "EnumValueOption",
		NTType:     30,
		Index:      90,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Message : "message" MessageName "{" "}"	<< X[0], nil >>`,
		Id:         "Message",
		NTType:     31,
		Index:      91,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Message : "message" MessageName "{" MessageBody "}"	<< X[0], nil >>`,
		Id:         "Message",
		NTType:     31,
		Index:      92,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageBody : MessageContent	<< X[0], nil >>`,
		Id:         "MessageBody",
		NTType:     32,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageBody : MessageBody MessageContent	<< X[0], nil >>`,
		Id:         "MessageBody",
		NTType:     32,
		Index:      94,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Field	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Enum	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Message	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Option	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      98,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Oneof	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : MapField	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : Reserved	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageContent : emptyStatement	<< X[0], nil >>`,
		Id:         "MessageContent",
		NTType:     33,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Service : "service" ServiceName "{" "}"	<< X[0], nil >>`,
		Id:         "Service",
		NTType:     34,
		Index:      103,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Service : "service" ServiceName "{" ServiceBodies "}"	<< X[0], nil >>`,
		Id:         "Service",
		NTType:     34,
		Index:      104,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceBodies : ServiceBody	<< X[0], nil >>`,
		Id:         "ServiceBodies",
		NTType:     35,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceBodies : ServiceBodies ServiceBody	<< X[0], nil >>`,
		Id:         "ServiceBodies",
		NTType:     35,
		Index:      106,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceBody : Option	<< X[0], nil >>`,
		Id:         "ServiceBody",
		NTType:     36,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceBody : Rpc	<< X[0], nil >>`,
		Id:         "ServiceBody",
		NTType:     36,
		Index:      108,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceBody : emptyStatement	<< X[0], nil >>`,
		Id:         "ServiceBody",
		NTType:     36,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rpc : "rpc" RpcName "(" MessageType ")" "returns" "(" MessageType ")" RpcBodies	<< X[0], nil >>`,
		Id:         "Rpc",
		NTType:     37,
		Index:      110,
		NumSymbols: 10,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rpc : "rpc" RpcName "(" "stream" MessageType ")" "returns" "(" MessageType ")" RpcBodies	<< X[0], nil >>`,
		Id:         "Rpc",
		NTType:     37,
		Index:      111,
		NumSymbols: 11,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rpc : "rpc" RpcName "(" MessageType ")" "returns" "(" "stream" MessageType ")" RpcBodies	<< X[0], nil >>`,
		Id:         "Rpc",
		NTType:     37,
		Index:      112,
		NumSymbols: 11,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rpc : "rpc" RpcName "(" "stream" MessageType ")" "returns" "(" "stream" MessageType ")" RpcBodies	<< X[0], nil >>`,
		Id:         "Rpc",
		NTType:     37,
		Index:      113,
		NumSymbols: 12,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBodies : "{" "}"	<< X[0], nil >>`,
		Id:         "RpcBodies",
		NTType:     38,
		Index:      114,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBodies : "{" RpcBody "}"	<< X[0], nil >>`,
		Id:         "RpcBodies",
		NTType:     38,
		Index:      115,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBodies : ";"	<< X[0], nil >>`,
		Id:         "RpcBodies",
		NTType:     38,
		Index:      116,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBody : Option	<< X[0], nil >>`,
		Id:         "RpcBody",
		NTType:     39,
		Index:      117,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBody : emptyStatement	<< X[0], nil >>`,
		Id:         "RpcBody",
		NTType:     39,
		Index:      118,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBody : RpcBody Option	<< X[0], nil >>`,
		Id:         "RpcBody",
		NTType:     39,
		Index:      119,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcBody : RpcBody emptyStatement	<< X[0], nil >>`,
		Id:         "RpcBody",
		NTType:     39,
		Index:      120,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FullIdent : ident	<< X[0], nil >>`,
		Id:         "FullIdent",
		NTType:     40,
		Index:      121,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FullIdent : FullIdent "." ident	<< X[0], nil >>`,
		Id:         "FullIdent",
		NTType:     40,
		Index:      122,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageName : ident	<< X[0], nil >>`,
		Id:         "MessageName",
		NTType:     41,
		Index:      123,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumName : ident	<< X[0], nil >>`,
		Id:         "EnumName",
		NTType:     42,
		Index:      124,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldName : ident	<< X[0], nil >>`,
		Id:         "FieldName",
		NTType:     43,
		Index:      125,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `OneofName : ident	<< X[0], nil >>`,
		Id:         "OneofName",
		NTType:     44,
		Index:      126,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MapName : ident	<< X[0], nil >>`,
		Id:         "MapName",
		NTType:     45,
		Index:      127,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ServiceName : ident	<< X[0], nil >>`,
		Id:         "ServiceName",
		NTType:     46,
		Index:      128,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RpcName : ident	<< X[0], nil >>`,
		Id:         "RpcName",
		NTType:     47,
		Index:      129,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageType : MessageName	<< X[0], nil >>`,
		Id:         "MessageType",
		NTType:     48,
		Index:      130,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageType : Idents MessageName	<< X[0], nil >>`,
		Id:         "MessageType",
		NTType:     48,
		Index:      131,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageType : "." MessageName	<< X[0], nil >>`,
		Id:         "MessageType",
		NTType:     48,
		Index:      132,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MessageType : "." Idents MessageName	<< X[0], nil >>`,
		Id:         "MessageType",
		NTType:     48,
		Index:      133,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Idents : ident "."	<< X[0], nil >>`,
		Id:         "Idents",
		NTType:     49,
		Index:      134,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Idents : Idents ident "."	<< X[0], nil >>`,
		Id:         "Idents",
		NTType:     49,
		Index:      135,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumType : EnumName	<< X[0], nil >>`,
		Id:         "EnumType",
		NTType:     50,
		Index:      136,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumType : "." EnumName	<< X[0], nil >>`,
		Id:         "EnumType",
		NTType:     50,
		Index:      137,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EnumType : "." Idents EnumName	<< X[0], nil >>`,
		Id:         "EnumType",
		NTType:     50,
		Index:      138,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : FullIdent	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      139,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : "-" intLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      140,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : "+" intLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      141,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : "-" floatLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      142,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : "+" floatLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      143,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : strLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      144,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Constant : boolLit	<< X[0], nil >>`,
		Id:         "Constant",
		NTType:     51,
		Index:      145,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}