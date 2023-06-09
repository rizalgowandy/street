package rqProperty

import (
	"github.com/kokizzu/gotro/S"

	"street/conf"
)

// TODO: this is slow, use tarantool-go api instead

func (rq *Property) FindPropertiesBySerialNumber(serialNumber string) (res []*Property) {
	query := `SELECT ` + rq.SqlSelectAllFields() + ` 
FROM ` + rq.SqlTableName() + `
WHERE ` + rq.SqlSerialNumber() + ` = ` + S.Q(serialNumber)
	if conf.IsDebug() {
		//L.Print(query)
	}
	rq.Adapter.QuerySql(query, func(row []any) {
		obj := &Property{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return res
}
