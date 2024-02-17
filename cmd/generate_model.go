package cmd

import (
	"slices"

	"github.com/kberov/slovo2/model"
	"github.com/kberov/slovo2/slovo"
	"github.com/spf13/cobra"
)

// generate/modelCmd represents the generate/model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Generates Go code for sqlx",
	Long: `
This is WIP - pre alpha - and is not a priority at all.
We used pieces of code generated by 'xo' and 'jet-gen'.
This command generates idiomatic (I'm trying) Go code for using sqlx with your
own queries and data objects almost like an ORM. We use sqlite3 and extract all
the meta data for the model from the database. Please, pass the path to your
database file or add it to the configuration section Cfg.Db.DSN
`,
	Run: func(cmd *cobra.Command, args []string) {
		// logger.Warnf("%#v", slovo.Cfg.DB.Tables)
		Logger.Print("generate/model called")
		// generateRecordTypes(slovo.Cfg.DB.Tables)

	},
}

func init() {
	generateCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.
	modelCmd.Flags().StringVarP(&slovo.Cfg.DB.DSN, "DSN", "D", slovo.Cfg.DB.DSN, "DSN for the database")
	// modelCmd.Flags().StringSliceVarP(&slovo.Cfg.DB.Tables, "tables", "t", slovo.Cfg.DB.Tables, "Tables for which to generate model types")
}

// get table names
var selectTables = `
SELECT name, type FROM sqlite_schema
WHERE type ='table'AND  name NOT LIKE 'sqlite_%';
`

func generateRecordTypes(tables []string) {
	Logger.Printf("tables from the command line: %#v", tables)
	dbh := model.DB()
	rows, err := dbh.Query(selectTables)
	if err != nil {
		Logger.Warn("Error" + err.Error())
		return
	}
	var tablesInDB []string
	for rows.Next() {
		var objectName string
		var objType string
		err = rows.Scan(&objectName, &objType)
		if err != nil {
			Logger.Warn("Error" + err.Error())
			return
		}
		Logger.Debug(objectName + " " + objType)
		if slices.Contains(tables, objectName) {
			tablesInDB = append(tablesInDB, objectName)
		}
	}
	Logger.Printf("The following of the requested tables wer found in the database: %#v", tablesInDB)
	//p := model.Products{}
	p := model.Stranici{}
	model.GetByID(&p, 1)
}
