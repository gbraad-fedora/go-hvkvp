package hvkvp

import (
	"testing"
)

func TestGetKvpRecords(t *testing.T) {
    records := getKvpRecords("./testdata/testpool")
    expected := 2
    actual := len(records)
    
    if actual != expected {
		t.Errorf("Expected '%d' but got '%d'.", expected, actual)
	}
}

func TestGetKvpRecordsByKey(t *testing.T) {
    records := getKvpRecords("./testdata/testpool")

    testData := map[string]string{
        "IpAddress": "10.0.75.128",
        "42": "Answer to the Ultimate Question of Life, the Universe, and Everything",
    }

    for _, record := range records {
        key := record.GetKey()
        actual := record.GetValue()
        expected := testData[key]
        
        if actual != expected {
	    	t.Errorf("Expected '%s' but got '%s' using '%s' as key.", expected, actual, key)
	    }
    }
}