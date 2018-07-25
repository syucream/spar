%{
package parser

import (
	"github.com/syucream/spar/src/types"
)
%}

%union {
  empty     struct{}
  str       string
  strs      []string
  col       types.Column
  cols      []types.Column
  key       types.Key
  keys      []types.Key
  clstr     types.Cluster
  alt       types.Alteration
  LastToken int
}

%token<str> PRIMARY KEY ASC DESC
%token<str> INTERLEAVE IN PARENT
%token<str> ARRAY OPTIONS
%token<str> NOT NULL
%token<str> ON DELETE CASCADE NO ACTION
%token<str> MAX UNIQUE NULL_FILTERED STORING
%token<str> ADD COLUMN SET
%token<str> true null allow_commit_timestamp
%token<empty> '(' ',' ')' ';' '='
%token<str> CREATE ALTER DROP
%token<str> DATABASE TABLE INDEX
%token<str> BOOL INT64 FLOAT64 STRING BYTES DATE TIMESTAMP

%token<str> database_id
%token<str> table_name
%token<str> column_name
%token<str> index_name

%type<col> column_def
%type<cols> column_def_list
%type<str> column_type scalar_type array_type length int64_value
%type<str> options_def
%token<str> decimal_value hex_value

%type<str> key_order_opt
%type<key> key_part
%type<keys> key_part_list
%type<keys> primary_key

%type<clstr> cluster_opt
%type<clstr> cluster
%type<str> on_delete_opt

%type<str> not_null_opt
%type<str> unique_opt
%type<str> null_filtered_opt

%type<alt> table_alteration
%type<alt> table_column_alteration

%start statements


%%

statements:
    statement
  | statements statement

statement:
    create_database ';'
  | create_table ';'
  | create_index ';'
  | alter_table ';'
  | drop_table ';'
  | drop_index ';'

create_database:
  CREATE DATABASE database_id
  {
    s := types.CreateDatabaseStatement{
      DatabaseId: $3,
    }
    yylex.(*LexerWrapper).Result.CreateDatabases = append(yylex.(*LexerWrapper).Result.CreateDatabases, s)
  }

create_table:
  CREATE TABLE table_name '(' column_def_list ')' primary_key cluster_opt
  {
    s := types.CreateTableStatement{
      TableName:   $3,
      Columns:     $5,
      PrimaryKeys: $7,
      Cluster:     $8,
    }
    yylex.(*LexerWrapper).Result.CreateTables = append(yylex.(*LexerWrapper).Result.CreateTables, s)
  }

column_def_list:
  /* empty */
  {
    $$ = make([]types.Column, 0, 0)
  }
  | column_def
  {
    $$ = make([]types.Column, 0, 1)
    $$ = append($$, $1)
  }
  | column_def ',' column_def_list
  {
    $$ = append($3, $1)
  }

column_def:
  column_name column_type not_null_opt options_def
  {
    $$ = types.Column{Name: $1, Type: $2, Nullability: $3, Options: $4}
  }

primary_key:
  PRIMARY KEY '(' key_part_list ')'
  {
    $$ = $4
  }

key_part_list:
    key_part
  {
    $$ = make([]types.Key, 0, 1)
    $$ = append($$, $1)
  }
  | key_part_list ',' key_part
  {
    $$ = append($1, $3)
  }

key_part:
  column_name key_order_opt
  {
    $$ = types.Key{Name: $1, Order: $2}
  }

key_order_opt:
  /* empty */
  {
    $$ = "ASC"
  }
  | ASC
  {
    $$ = $1
  }
  | DESC
  {
    $$ = $1
  }

cluster_opt:
  /* empty */
  {
    $$ = types.Cluster{}
  }
  | ',' cluster
  {
    $$ = $2
  }

cluster:
    INTERLEAVE IN PARENT table_name on_delete_opt
  {
    $$ = types.Cluster{TableName: $4, OnDelete: $5}
  }

on_delete_opt:
  /* empty */
  {
    // default
    $$ = "NO ACTION"
  }
  | ON DELETE CASCADE
  {
    $$ = $3
  }
  | ON DELETE NO ACTION
  {
    $$ = $3 + " " + $4
  }

column_type:
    scalar_type
  {
    $$ = $1
  }
  | array_type
  {
    $$ = $1
  }

scalar_type:
    BOOL
  {
    $$ = $1
  }
  | INT64
  {
    $$ = $1
  }
  | FLOAT64
  {
    $$ = $1
  }
  | STRING '(' length ')'
  {
    $$ = $1 + "(" + $3 + ")"
  }
  | BYTES '(' length ')'
  {
    $$ = $1 + "(" + $3 + ")"
  }
  | DATE
  {
    $$ = $1
  }
  | TIMESTAMP
  {
    $$ = $1
  }

length:
    int64_value
  {
    $$ = $1
  }
  | MAX
  {
    $$ = $1
  }

array_type:
  ARRAY '<' scalar_type '>'
  {
    $$ = $1 + "(" + $3 + ")"
  }

options_def:
  /* empty */
  {
    $$ = ""
  }
  | OPTIONS '(' allow_commit_timestamp '=' true ')'
  {
    $$ = $3 + "=" + $5
  }
  | OPTIONS '(' allow_commit_timestamp '=' null ')'
  {
    $$ = $3 + "=" + $5
  }

not_null_opt:
  /* empty */
  {
    $$ = "NULL"
  }
  | NOT NULL
  {
    $$ = $1 + " " + $2
  }
  
create_index:
  CREATE unique_opt null_filtered_opt INDEX index_name ON table_name '(' key_part_list ')' storing_clause_opt interleave_clause_list
  {
    // TODO Support options
    s := types.CreateIndexStatement{
      Unique:       $2,
      NullFiltered: $3,
      IndexName:    $5,
      TableName:    $7,
      Keys:         $9,
    }
    yylex.(*LexerWrapper).Result.CreateIndexes = append(yylex.(*LexerWrapper).Result.CreateIndexes, s)
  }

unique_opt:
  /* empty */
  {
    $$ = ""
  }
  | UNIQUE
  {
    $$ = $1
  }

null_filtered_opt:
  /* empty */
  {
    $$ = ""
  }
  | NULL_FILTERED
  {
    $$ = $1
  }

storing_clause_opt:
  /* empty */
  | storing_clause

storing_clause:
    STORING '(' column_name_list ')'

column_name_list:
    column_name
  | column_name_list ',' column_name

interleave_clause_list:
  /* empty */
  | interleave_clause
  | interleave_clause_list ',' interleave_clause

interleave_clause:
    INTERLEAVE IN table_name

alter_table:
    ALTER TABLE table_name table_alteration
  {
    s := types.AlterTableStatement{
      TableName:  $3,
      Alteration: $4,
    }
    yylex.(*LexerWrapper).Result.AlterTables = append(yylex.(*LexerWrapper).Result.AlterTables, s)
  }
  | ALTER TABLE table_name table_column_alteration
  {
    s := types.AlterTableStatement{
      TableName:  $3,
      Alteration: $4,
    }
    yylex.(*LexerWrapper).Result.AlterTables = append(yylex.(*LexerWrapper).Result.AlterTables, s)
  }

table_alteration:
    ADD COLUMN column_def
  {
    $$ = &types.AddColumnTableAlteration{
      Column: $3,
    }
  }
  | DROP COLUMN column_name
  {
    $$ = &types.DropColumnTableAlteration{
      ColumnName: $3,
    }
  }
  | SET on_delete_opt
  {
    $$ = &types.SetTableAlteration{
      OnDelete: $2,
    }
  }

table_column_alteration:
    ALTER COLUMN column_name column_type not_null_opt
  {
    $$ = &types.AlterColumnTypesAlteration{
      ColumnName:  $3,
      ColumnType:  $4,
      Nullability: $5,
    }
  }
  | ALTER COLUMN column_name SET options_def
  {
    $$ = &types.AlterColumnSetAlteration{
      ColumnName: $3,
      Options:    $5,
    }
  }

drop_table:
    DROP TABLE table_name
  {
    s := types.DropTableStatement{
      TableName: $3,
    }
    yylex.(*LexerWrapper).Result.DropTables = append(yylex.(*LexerWrapper).Result.DropTables, s)
  }

drop_index:
    DROP INDEX index_name
  {
    s := types.DropIndexStatement{
      IndexName: $3,
    }
    yylex.(*LexerWrapper).Result.DropIndexes = append(yylex.(*LexerWrapper).Result.DropIndexes, s)
  }

int64_value:
    decimal_value
  {
    $$ = $1
  }
  | hex_value
  {
    $$ = $1
  }

/* TODO Check some literals satisfy regexp specs on tokenizer.
decimal_value:
  [-]0—9+

hex_value:
  [-]0x{0—9|a—f|A—F}+

database_id:
  {a—z}[{a—z|0—9|_|-}+]{a—z|0—9}

table_name:
  {a—z|A—Z}[{a—z|A—Z|0—9|_}+]

column_name:
  {a—z|A—Z}[{a—z|A—Z|0—9|_}+]

index_name:
  {a—z|A—Z}[{a—z|A—Z|0—9|_}+]
*/
