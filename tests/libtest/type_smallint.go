package libtest

import (
	"database/sql"

	"testing"
)

// DoTestSmallInt tests the handling of the SmallInt.
func DoTestSmallInt(t *testing.T) {
	TestForEachDB("TestSmallInt", t, testSmallInt)
	//
}

func testSmallInt(t *testing.T, db *sql.DB, tableName string) {
	pass := make([]interface{}, len(samplesSmallInt))
	mySamples := make([]int16, len(samplesSmallInt))

	for i, sample := range samplesSmallInt {

		mySample := sample

		pass[i] = mySample
		mySamples[i] = mySample
	}

	rows, err := SetupTableInsert(db, tableName, "smallint", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()

	i := 0
	var recv int16
	for rows.Next() {
		err = rows.Scan(&recv)
		if err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if recv != mySamples[i] {

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
