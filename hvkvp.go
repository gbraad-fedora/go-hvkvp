package hvkvp

import (
	"io"
	"fmt"
	"os"
	"bytes"
	"strings"
	"encoding/binary"
)

type kvp_record struct {
	Key   [MAX_KEY_SIZE]byte
	Value [MAX_VALUE_SIZE]byte
}

func (record *kvp_record) GetKey() string {
	return strings.Trim(string(record.Key[:MAX_KEY_SIZE]), "\x00")
}

func (record *kvp_record) GetValue() string {
	return strings.Trim(string(record.Value[:MAX_VALUE_SIZE]), "\x00")
}

const (
	MAX_KEY_SIZE   = 512
	MAX_VALUE_SIZE = 2048
	DEFAULT_POOLNAME = "/var/lib/hyperv/.kvp_pool_0"
	READ_FORMAT = "Key: %s, Value: %s\n"
	EXPORT_FORMAT = "export %s=\"%s\"\n"
)

func readNextBytes(file *os.File, number int) ([]byte, error) {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func getKvpRecords(poolFile string) []kvp_record {
	file, err := os.Open(poolFile)
	if err != nil {
		fmt.Println("Error opening pool")
		os.Exit(3)
	}
	
	var records []kvp_record

	for {
		record := kvp_record{}
		data, err := readNextBytes(file, MAX_KEY_SIZE + MAX_VALUE_SIZE)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &record)
		if err == io.EOF {
			break
		}

		records = append(records, record)
	}
	
	return records
}

func GetKvpRecordByKey(key string) {
	for _, record := range getKvpRecords(DEFAULT_POOLNAME) {
		if(strings.EqualFold(record.GetKey(), key)) {
			fmt.Printf(record.GetValue())
			break
		}
	}
}

func GetAllKvpRecords(format string) {
	for _, record := range getKvpRecords(DEFAULT_POOLNAME) {
		fmt.Printf(format, record.GetKey(), record.GetValue())
	}
}
