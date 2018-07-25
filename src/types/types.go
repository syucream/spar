package types

type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
	CreateIndexes   []CreateIndexStatement
	AlterTables     []AlterTableStatement
	DropTables      []DropTableStatement
	DropIndexes     []DropIndexStatement
}

type Column struct {
	Name        string
	Type        string // BOOL, INT64, ...
	Nullability string
	Options     string
}

type Key struct {
	Name  string
	Order string
}

type Cluster struct {
	TableName string
	OnDelete  string
}

type Alteration interface {
	Alter()
}

type AddColumnTableAlteration struct {
	Column Column
}

// Alter satisfies Alteration interface
func (self *AddColumnTableAlteration) Alter() {}

type DropColumnTableAlteration struct {
	ColumnName string
}

// Alter satisfies Alteration interface
func (self *DropColumnTableAlteration) Alter() {}

type SetTableAlteration struct {
	OnDelete string
}

// Alter satisfies Alteration interface
func (self *SetTableAlteration) Alter() {}

type AlterColumnTypesAlteration struct {
	ColumnName  string
	ColumnType  string
	Nullability string
}

// Alter satisfies Alteration interface
func (self *AlterColumnTypesAlteration) Alter() {}

type AlterColumnSetAlteration struct {
	ColumnName string
	Options    string
}

// Alter satisfies Alteration interface
func (self *AlterColumnSetAlteration) Alter() {}

type CreateDatabaseStatement struct {
	DatabaseId string
}

type CreateTableStatement struct {
	TableName   string
	Columns     []Column
	PrimaryKeys []Key
	Cluster     Cluster
}

type CreateIndexStatement struct {
	IndexName    string
	Unique       string
	NullFiltered string
	TableName    string
	Keys         []Key
}

type AlterTableStatement struct {
	TableName  string
	Alteration Alteration // Get concrete type yourself
}

type DropTableStatement struct {
	TableName string
}

type DropIndexStatement struct {
	IndexName string
}
