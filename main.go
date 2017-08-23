package main

import (
        "io"
        "fmt"
        "os"
        "bytes"
        "encoding/binary"
)

type kvp_record struct {
        Key   [MAX_KEY_SIZE]byte
        Value [MAX_VALUE_SIZE]byte
}

func (record *kvp_record) GetKey() string {
        return string(record.Key[:MAX_KEY_SIZE])
}

func (record *kvp_record) GetValue() string {
        return string(record.Value[:MAX_VALUE_SIZE])
}

const (
        MAX_KEY_SIZE   = 512
        MAX_VALUE_SIZE = 2048
)

func readNextBytes(file *os.File, number int) ([]byte, error) {
        bytes := make([]byte, number)

        _, err := file.Read(bytes)
        if err != nil {
                return nil, err
        }

        return bytes, nil
}

func main() {
        poolName := "/var/lib/hyperv/.kvp_pool_0"

        file, err := os.Open(poolName)
        if err != nil {
                panic("Oops")
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

        for _, record := range records {
                fmt.Printf("Key: %s, Value: %s\n", record.GetKey(), record.GetValue())
        }
}
