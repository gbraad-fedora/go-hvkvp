package main

import (
	"flag"
	"os"

	hvkvp "github.com/gbraad/go-hvkvp"
)

const (
	READ_FORMAT = "Key: %s, Value: %s\n"
	EXPORT_FORMAT = "export %s=%s\n"
)

func main() {
	exportMode := flag.Bool("export", false, "Return all for export as environment variable")
	searchMode := flag.String("key", "", "Search for a specific key and return value")
	flag.Parse()

	if(*searchMode != "") {
		hvkvp.GetKvpRecordByKey(*searchMode)
	} else if(*exportMode) {
		hvkvp.GetAllKvpRecords(EXPORT_FORMAT)
	} else {
		hvkvp.GetAllKvpRecords(READ_FORMAT)
	}
	os.Exit(0)
}
