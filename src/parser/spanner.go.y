%{
package parser

import (
	"github.com/syucream/spar/src/types"
)
%}

%union {
  empty     struct{}
  flag      bool
  str       string
  strs      []string
  col       types.Column
  cols      []types.Column
  key       types.Key
  keys      []types.Key
  keyorder  types.KeyOrder
  clstr     types.Cluster
  ondelete  types.OnDelete
  stcls     types.StoringClause
  intlr     types.Interleave
  intlrs    []types.Interleave
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

%type<keyorder> key_order_opt
%type<key> key_part
%type<keys> key_part_list
%type<keys> primary_key

%type<clstr> cluster_opt
%type<clstr> cluster
%type<ondelete> on_delete_opt

%type<flag> not_null_opt
%type<flag> unique_opt
%type<flag> null_filtered_opt

%type<strs> column_name_list
%type<stcls> storing_clause storing_clause_opt
%type<intlr> interleave_clause
%type<intlrs> interleave_clause_list

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
    yylex.(*lexerWrapper).result.CreateDatabases = append(yylex.(*lexerWrapper).result.CreateDatabases, s)
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
    yylex.(*lexerWrapper).result.CreateTables = append(yylex.(*lexerWrapper).result.CreateTables, s)
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
    $$ = types.Column{Name: $1, Type: $2, NotNull: $3, Options: $4}
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
    $$ = types.Key{Name: $1, KeyOrder: $2}
  }

key_order_opt:
  /* empty */
  {
    $$ = types.Asc
  }
  | ASC
  {
    $$ = types.Asc
  }
  | DESC
  {
    $$ = types.Desc
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
    $$ = types.NoAction
  }
  | ON DELETE CASCADE
  {
    $$ = types.Cascade
  }
  | ON DELETE NO ACTION
  {
    $$ = types.NoAction
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
    $$ = types.False
  }
  | NOT NULL
  {
    $$ = types.True
  }
  
create_index:
  CREATE unique_opt null_filtered_opt INDEX index_name ON table_name '(' key_part_list ')' storing_clause_opt interleave_clause_list
  {
    s := types.CreateIndexStatement{
      Unique:        $2,
      NullFiltered:  $3,
      IndexName:     $5,
      TableName:     $7,
      Keys:          $9,
      StoringClause: $11,
      Interleaves:   $12,
    }
    yylex.(*lexerWrapper).result.CreateIndexes = append(yylex.(*lexerWrapper).result.CreateIndexes, s)
  }

unique_opt:
  /* empty */
  {
    $$ = types.False
  }
  | UNIQUE
  {
    $$ = types.True
  }

null_filtered_opt:
  /* empty */
  {
    $$ = types.False
  }
  | NULL_FILTERED
  {
    $$ = types.True
  }

storing_clause_opt:
  /* empty */
  {
    $$ = types.StoringClause{}
  }
  | storing_clause
  {
    $$ = $1
  }

storing_clause:
    STORING '(' column_name_list ')'
  {
    $$ = types.StoringClause{ColumnNames: $3}
  }

column_name_list:
    column_name
  {
    $$ = make([]string, 0, 1)
    $$ = append($$, $1)
  }
  | column_name_list ',' column_name
  {
    $$ = append($1, $3)
  }

interleave_clause_list:
  /* empty */
  {
    $$ = make([]types.Interleave, 0, 0)
  }
  | interleave_clause
  {
    $$ = make([]types.Interleave, 0, 1)
    $$ = append($$, $1)
  }
  | interleave_clause_list ',' interleave_clause
  {
    $$ = append($1, $3)
  }

interleave_clause:
    INTERLEAVE IN table_name
  {
    $$ = types.Interleave{TableName: $3}
  }

alter_table:
    ALTER TABLE table_name table_alteration
  {
    s := types.AlterTableStatement{
      TableName:  $3,
      Alteration: $4,
    }
    yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
  }
  | ALTER TABLE table_name table_column_alteration
  {
    s := types.AlterTableStatement{
      TableName:  $3,
      Alteration: $4,
    }
    yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
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
      NotNull:     $5,
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
    yylex.(*lexerWrapper).result.DropTables = append(yylex.(*lexerWrapper).result.DropTables, s)
  }

drop_index:
    DROP INDEX index_name
  {
    s := types.DropIndexStatement{
      IndexName: $3,
    }
    yylex.(*lexerWrapper).result.DropIndexes = append(yylex.(*lexerWrapper).result.DropIndexes, s)
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
