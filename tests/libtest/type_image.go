package libtest

import (
	"database/sql"

	"testing"
)

// DoTestImage tests the handling of the Image.
func DoTestImage(t *testing.T) {
	TestForEachDB("TestImage", t, testImage)
	//
}

func testImage(t *testing.T, db *sql.DB, tableName string) {
	pass := make([]interface{}, len(samplesImage))
	mySamples := make([][]byte, len(samplesImage))

	for i, sample := range samplesImage {

		mySample := sample

		pass[i] = mySample
		mySamples[i] = mySample
	}

	rows, err := SetupTableInsert(db, tableName, "image", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()

	i := 0
	var recv []byte
	for rows.Next() {
		err = rows.Scan(&recv)
		if err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if compareBinary(recv, mySamples[i]) {

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
