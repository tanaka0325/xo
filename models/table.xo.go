// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// Table represents table info.
type Table struct {
	Type      string // type
	TableName string // table_name
}

// PgTables runs a custom query, returning results as Table.
func PgTables(db XODB, schema string, relkind string) ([]*Table, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`c.relkind, ` + // ::varchar AS type
		`c.relname ` + // ::varchar AS table_name
		`FROM pg_class c ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`WHERE n.nspname = $1 AND c.relkind = $2`

	// run query
	XOLog(sqlstr, schema, relkind)
	q, err := db.Query(sqlstr, schema, relkind)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Table{}
	for q.Next() {
		t := Table{}

		// scan
		err = q.Scan(&t.Type, &t.TableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// MyTables runs a custom query, returning results as Table.
func MyTables(db XODB, schema string, relkind string) ([]*Table, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`table_name ` +
		`FROM information_schema.tables ` +
		`WHERE table_schema = ? AND table_type = ?`

	// run query
	XOLog(sqlstr, schema, relkind)
	q, err := db.Query(sqlstr, schema, relkind)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Table{}
	for q.Next() {
		t := Table{}

		// scan
		err = q.Scan(&t.TableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// SqTables runs a custom query, returning results as Table.
func SqTables(db XODB, relkind string) ([]*Table, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tbl_name AS table_name ` +
		`FROM sqlite_master ` +
		`WHERE tbl_name NOT LIKE 'sqlite_%' AND type = ?`

	// run query
	XOLog(sqlstr, relkind)
	q, err := db.Query(sqlstr, relkind)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Table{}
	for q.Next() {
		t := Table{}

		// scan
		err = q.Scan(&t.TableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}