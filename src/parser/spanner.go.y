%{
package parser
%}

%union {
  empty     struct{}
  str       string
  strs      []string
  col       Column
  cols      []Column
  key       Key
  keys      []Key
  LastToken int
}

%token<str> PRIMARY KEY ASC DESC
%token<str> INTERLEAVE IN PARENT
%token<str> ARRAY OPTIONS
%token<str> NOT NULL
%token<str> ON DELETE CASCADE NO ACTION
%token<str> MAX UNIQUE NULL_FILTERED STORING
%token<str> true null allow_commit_timestamp
%token<empty> '(' ',' ')' ';'
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
%token<str> decimal_value hex_value

%type<str> key_order_opt
%type<key> key_part
%type<keys> key_part_list
%type<keys> primary_key

%start statements


%%

statements:
    statement
  | statements statement

statement:
    create_database ';'
  | create_table ';'
  | create_index ';'
  /* TODO
  | alter_table ';'
  | drop_table ';'
  | drop_index ';'
  */

create_database:
  CREATE DATABASE database_id
  {
    SetCreateDatabaseStatement(yylex, $1, $2, $3)
  }

create_table:
  CREATE TABLE table_name '(' column_def_list ')' primary_key cluster_opt
  {
    SetCreateTableStatement(yylex, $1, $2, $3, $5, $7)
  }

column_def_list:
  /* empty */
  {
    $$ = make([]Column, 0, 0)
  }
  | column_def
  {
    $$ = make([]Column, 0, 1)
    $$ = append($$, $1)
  }
  | column_def ',' column_def_list
  {
    $$ = append($3, $1)
  }

column_def:
  column_name column_type null_opt options_def
  {
    $$ = Column{Name: $1, Type: $2}
  }

primary_key:
  PRIMARY KEY '(' key_part_list ')'
  {
    $$ = $4
  }

key_part_list:
    key_part
  {
    $$ = make([]Key, 0, 1)
    $$ = append($$, $1)
  }
  | key_part_list ',' key_part
  {
    $$ = append($1, $3)
  }

key_part:
  column_name key_order_opt
  {
    $$ = Key{Name: $1, Order: $2}
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
  | cluster

cluster:
    INTERLEAVE IN PARENT table_name cluster_on_delete

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
    $$ = "2621440"
  }

array_type:
  ARRAY '<' scalar_type '>'
  {
    $$ = $1 + "(" + $3 + ")"
  }

options_def:
  /* empty */
  | OPTIONS '(' allow_commit_timestamp '=' true ')'
  | OPTIONS '(' allow_commit_timestamp '=' null ')'

null_opt:
  /* empty */
  | NOT NULL
  
cluster_on_delete:
  /* empty */
  | ON DELETE CASCADE
  | ON DELETE NO ACTION

create_index:
  CREATE unique_opt null_filtered_opt INDEX index_name ON table_name '(' key_part_list ')' storing_clause_opt interleave_clause_list

unique_opt:
  /* empty */
  | UNIQUE

null_filtered_opt:
  /* empty */
  | NULL_FILTERED

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

/*
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
  {
    $$ = $1
  }
  | hex_value
  {
    $$ = $1
  }

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
