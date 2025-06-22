package bug_test

import (
	"database/sql"
	"testing"

	"bug"

	_ "github.com/glebarez/go-sqlite"
)

const initQuery = `BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS "group_member" (
	"id"	INTEGER NOT NULL
);

COMMIT;
`

func TestStupidity(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	bug.FailIfErr(t, err, "db connection")

	_, err = db.Exec(initQuery)
	bug.FailIfErr(t, err, "SQL init")

	tables := bug.ListTables(db)
	bug.FatalAssertEq(t, "table length before", len(tables), 1)

	stmt := "SELECT id FROM group_member WHERE id = ?;"
	_, err = db.Query(stmt, 0)
	bug.FailIfErr(t, err, "get group member")

	tables = bug.ListTables(db)
	bug.FatalAssertEq(t, "table length after", len(tables), 1)
}
