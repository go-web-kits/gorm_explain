package gorm_explain

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/go-web-kits/utils/logx"
	"github.com/jinzhu/gorm"
)

func Register(db *gorm.DB) {
	db.Callback().Query().Register("gorm_explain", callback)
}

func callback(scope *gorm.Scope) {
	if os.Getenv("EXPLAIN") != "true" {
		return
	}

	if !strings.HasPrefix(strings.ToUpper(scope.SQL), "SELECT") {
		return
	}

	rows, err := scope.SQLDB().Query("EXPLAIN ANALYZE "+scope.SQL, scope.SQLVars...)
	if scope.Err(err) != nil {
		return
	}
	defer rows.Close()

	result, err := convertToResult(rows)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(logx.Yello("EXPLAIN ANALYZE for: " + scope.SQL + " <- " + fmt.Sprint(scope.SQLVars)))
	fmt.Println(logx.Yello(makeString((maxLenOf(result)-11)/2, byte(" "[0])) + "QUERY PLAIN"))
	fmt.Println(logx.Yello(makeString(maxLenOf(result), byte("-"[0]))))
	fmt.Println(logx.Yello(strings.Join(result, "\n")))
}

func convertToResult(rows *sql.Rows) (result []string, err error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))

	for i := range values {
		args[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}
		row := []string{}
		for _, col := range values {
			row = append(row, "=> "+string(col))
		}
		result = append(result, row...)
	}

	return result, nil
}

func maxLenOf(slice []string) int {
	result := 0
	for _, v := range slice {
		if len(v) > result {
			result = len(v)
		}
	}
	return result
}

func makeString(len int, defVal byte) string {
	str := make([]byte, len)
	for i := 0; i < len; i++ {
		str[i] = defVal
	}
	return string(str)
}
