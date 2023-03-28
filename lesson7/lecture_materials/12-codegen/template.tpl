package {{ .Package }}

import _ "database/sql"

func (r *{{.Struct}}) List() ([]{{ .Struct }}, error) {
    sql := "select {{ join .Fields ", " }} from {{ .TableName }};"
    _ = sql
    var result []{{ .Struct }}

    // TODO...

    return result, nil
}