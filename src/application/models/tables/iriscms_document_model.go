package tables

type IriscmsDocumentModel struct {
	Id          int64 `xorm:"pk"`
	Name        string
	Table 		string
	Enabled     int
	IsSystem    int
	ModelType   int
	FeTplIndex  string
	FeTplList   string
	FeTplDetail string
}