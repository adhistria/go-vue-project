package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

func InitializeDB(db *sqlx.DB) {
	csvFilePaths := []string{"customer_companies.csv", "customers.csv", "orders.csv", "order_items.csv", "deliveries.csv"}
	for _, filePath := range csvFilePaths {
		err := execSchema(db, "test_data/"+filePath)
		if err != nil {
			log.Println(err)
		}
	}
}

func execSchema(db *sqlx.DB, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return err
	}
	headers := data[0]
	columnMapping := []string{}
	columnDataType := map[string]string{}
	for i, v := range headers {
		dataType := checkDataType(data[1][i])
		if i == 0 {
			dataType += " PRIMARY KEY"
		}
		columnMapping = append(columnMapping, v+" "+dataType)
		columnDataType[v] = dataType
	}

	var rows []string
	for i := 1; i < len(data); i++ {
		values := []string{}
		for j := range data[i] {
			if columnDataType[headers[j]] == "INTEGER" {
				values = append(values, data[i][j])
			} else if columnDataType[headers[j]] == "DECIMAL" {
				if _, err = strconv.ParseFloat(data[i][j], 64); err != nil {
					values = append(values, "NULL")
				} else {
					values = append(values, data[i][j])
				}
			} else if columnDataType[headers[j]] == "TEXT[]" {
				arrStr := strings.Split(data[i][j], ",")
				newStr := []string{}
				for _, str := range arrStr {
					newStr = append(newStr, "'"+str+"'")
				}
				values = append(values, "ARRAY"+"["+strings.Join(newStr, ",")+"]")
			} else {
				values = append(values, "'"+data[i][j]+"'")
			}
		}
		rows = append(rows, "("+strings.Join(values, ",")+")")
	}

	columns := strings.Join(columnMapping, ", ")
	tableName := strings.Split(strings.Split(filePath, "/")[1], ".")[0]
	schema := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v);", tableName, columns)
	db.MustExec(schema)

	seedSchema := fmt.Sprintf(`INSERT INTO %v (%v) VALUES %v ; `, tableName, strings.Join(headers, ","), strings.Join(rows, ","))
	_, err = db.Exec(seedSchema)

	return err
}

func checkDataType(val string) string {
	if _, err := strconv.Atoi(val); err == nil {
		return "INTEGER"
	} else if _, err = strconv.ParseFloat(val, 64); err == nil {
		return "DECIMAL"
	} else if _, err = time.Parse(time.RFC3339, val); err == nil {
		return "TIMESTAMP"
	} else if val[0] == '[' && val[len(val)-1] == ']' {
		return "TEXT[]"
	} else {
		return "VARCHAR(255)"
	}
}
