%{
package parser

type Statement struct {
  Action string
  Target string
  Id string
}

func setStatement(yylex interface{}, stmt Statement) {
  yylex.(*Lexer).Stmt = stmt
}

%}

%union {
  empty     struct{}
  bytes     []byte
  str       string

  LastToken int
}

%token<bytes> PRIMARY KEY ASC DESC
%token<bytes> INTERLEAVE IN PARENT
%token<bytes> ARRAY OPTIONS
%token<bytes> NOT NULL
%token<bytes> ON DELETE CASCADE NO ACTION
%token<bytes> MAX
%token<bytes> true null allow_commit_timestamp
%token<empty> '(' ',' ')'
%token<str> CREATE ALTER DROP
%token<str> DATABASE TABLE INDEX

%token<bytes> BOOL INT64 FLOAT64 STRING BYTES DATE TIMESTAMP

%type<statement> create_database

%token<str> database_id
%token<str> table_name
%token<str> column_name
%token<str> decimal_value hex_value
%token<indexName> index_name

%start statement


%%

statement:
    create_database
  | create_table
/* TODO  { create_index | alter_table | drop_table | drop_index } */

create_database:
  CREATE DATABASE database_id
  {
    setStatement(yylex, Statement{Action: $1, Target: $2, Id: $3})
  }

create_table:
  CREATE TABLE table_name '(' column_def_opt ')' primary_keys cluster
  {
    setStatement(yylex, Statement{Action: $1, Target: $2, Id: $3})
  }

column_def_opt:
    column_def
  | column_def_opt ',' column_def

column_def:
  column_name column_type null_opt options_def

primary_keys:
    primary_key
  | primary_keys ',' primary_key

primary_key:
  PRIMARY KEY '(' key_part_opt ')'

key_part_opt:
    key_part
  | key_part_opt ',' key_part

key_part:
    column_name key_order_opt

key_order_opt:
  /* empty */
  | ASC
  | DESC

cluster:
    INTERLEAVE IN PARENT table_name cluster_opt

column_type:
    scalar_type
  | array_type

scalar_type:
    BOOL
  | INT64
  | FLOAT64
  | STRING '(' length ')'
  | BYTES '(' length ')'
  | DATE
  | TIMESTAMP

length:
    int64_value
  | MAX

array_type:
  ARRAY '<' scalar_type '>'

options_def:
  /* empty */
  | OPTIONS '(' allow_commit_timestamp '=' true ')'
  | OPTIONS '(' allow_commit_timestamp '=' null ')'

null_opt:
  /* empty */
  | NOT NULL
  
cluster_opt:
  /* empty */
  | ON DELETE CASCADE
  | ON DELETE NO ACTION

/*
create_index:
    CREATE [UNIQUE] [NULL_FILTERED] INDEX index_name
    ON table_name ( key_part [, ...] ) [ storing_clause ] [ , interleave_clause ]

storing_clause:
    STORING ( column_name [, ...] )

interleave_clause:
    INTERLEAVE IN table_name

alter_table:
    ALTER TABLE table_name { table_alteration | table_column_alteration }

table_alteration:
{ ADD COLUMN column_def | DROP COLUMN column_name |
      SET ON DELETE { CASCADE | NO ACTION } }

table_column_alteration:
    ALTER COLUMN column_name { { scalar_type | array_type } [NOT NULL] | SET options_def }

drop_table:
    DROP TABLE table_name

drop_index:
    DROP INDEX index_name
*/

int64_value:
    decimal_value
  | hex_value

/*
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
