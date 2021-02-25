package filestore

import (
	"fmt"
	"strings"
)

// DatatugFolder defines a folder name in a repo where to store DataTug project
const (
	BoardsFolder           = "boards"
	ProjectSummaryFileName = "datatug-project.json"
	DataFolder             = "data"
	DatatugFolder          = "datatug"
	DbFolder               = "db"
	DbCatalogsFolder       = "dbcatalogs"
	DbModelsFolder         = "dbmodels"
	EntitiesFolder         = "entities"
	EnvironmentsFolder     = "environments"
	QueriesFolder          = "queries"
	RecordsetsFolder       = "recordsets"
	ServersFolder          = "servers"
	SchemasFolder          = "schemas"
	TablesFolder           = "tables"
	ViewsFolder            = "views"
)

func jsonFileName(id, suffix string) string {
	switch suffix {
	case
		boardFileSuffix,
		dbCatalogFileSuffix,
		dbModelFileSuffix,
		dbServerFileSuffix,
		recordsetFileSuffix,
		environmentFileSuffix,
		entityFileSuffix,
		serverFileSuffix,
		columnsFileSuffix,
		queryFileSuffix:
		// OK
	default:
		panic("unknown JSON file suffix")

	}
	return fmt.Sprintf("%v.%v.json", id, suffix)
}

func getProjItemIdFromFileName(fileName string) (id string, suffix string) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 3 {
		return "", ""
	}
	suffixIndex := len(parts) - 2
	suffix = parts[suffixIndex]
	id = strings.Join(parts[:suffixIndex], ".")
	return
}

const (
	boardFileSuffix       = "board"
	dbCatalogFileSuffix   = "db"
	dbModelFileSuffix     = "dbmodel"
	dbServerFileSuffix    = "dbserver"
	recordsetFileSuffix   = "recordset"
	environmentFileSuffix = "env"
	entityFileSuffix      = "entity"
	serverFileSuffix      = "server"
	columnsFileSuffix     = "columns"
	queryFileSuffix       = "q"
)
