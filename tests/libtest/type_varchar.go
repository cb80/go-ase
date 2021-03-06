package libtest

import (
	"database/sql"

	"testing"
)

// DoTestVarChar tests the handling of the VarChar.
func DoTestVarChar(t *testing.T) {
	TestForEachDB("TestVarChar", t, testVarChar)
	//
}

func testVarChar(t *testing.T, db *sql.DB, tableName string) {
	pass := make([]interface{}, len(samplesVarChar))
	mySamples := make([]string, len(samplesVarChar))

	for i, sample := range samplesVarChar {

		mySample := sample

		pass[i] = mySample
		mySamples[i] = mySample
	}

	rows, err := SetupTableInsert(db, tableName, "varchar(13)", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()

	i := 0
	var recv string
	for rows.Next() {
		err = rows.Scan(&recv)
		if err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if compareChar(recv, mySamples[i]) {

			t.Errorf("Received value does not match passed parameter")
			t.Errorf("Expected: %v", mySamples[i])
			t.Errorf("Received: %v", recv)
		}

		i++
	}

	if err := rows.Err(); err != nil {
		t.Errorf("Error preparing rows: %v", err)
	}
}
