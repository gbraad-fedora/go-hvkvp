/*
Copyright (C) 2017 Gerard Braad <me@gbraad.nl>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
