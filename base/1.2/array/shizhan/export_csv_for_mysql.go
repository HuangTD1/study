package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 从mysql 读取数据到csv 文件中
var (
	tables = []string{"user", "orders"}
	count  = len(tables)
	ch     = make(chan bool, count)
)

func main() {
	// 链接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.104:3306)/chapter1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	for _, table := range tables {
		go SqlQuerys(db, table, ch)
	}
	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("done")

}

func SqlQuerys(db *sql.DB, table string, ch chan bool) {
	fmt.Println("开始处理导出:", table)
	rows, _ := db.Query(fmt.Sprintf("select * from %s", table))
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	totalValues := [][]string{}
	for rows.Next() {
		var s []string
		//把每行内容添加到scanArgs，也添加到了values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error() + "====")
		}
		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	exportToCsv(table+".csv", columns, totalValues)
	ch <- true
}

// 导出到csv
func exportToCsv(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	w := csv.NewWriter(f)
	for a, i := range totalValues {
		if a == 0 {
			w.Write(columns)
			w.Write(i)
		} else {
			w.Write(i)
		}
	}
	w.Flush()
	fmt.Println("写入完毕:", file)

}
