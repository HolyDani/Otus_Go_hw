package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/vmihailenco/msgpack"
	"gopkg.in/yaml.v2"
)

type Book struct {
	ID     uint    `json:"id" xml:"id" yaml:"id"`
	Size   uint    `json:"size" xml:"size" yaml:"size"`
	Year   uint    `json:"year" xml:"year" yaml:"year"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate"`
	Title  string  `json:"title" xml:"title" yaml:"title"`
	Author string  `json:"author" xml:"author" yaml:"author"`
}

func SerializeJSON(book []Book) ([]byte, error) {
	return json.Marshal(book)
}

func DeserializeJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("deserializeJSON: %w", err)
	}
	return books, nil
}

func SerializeXML(book []Book) ([]byte, error) {
	return xml.Marshal(book)
}

func DeserializeXML(data []byte) ([]Book, error) {
	var books []Book
	err := xml.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("deserializeXML: %w", err)
	}
	return books, nil
}

func SerializeYAML(book []Book) ([]byte, error) {
	return yaml.Marshal(book)
}

func DeserializeYAML(data []byte) ([]Book, error) {
	var books []Book
	err := yaml.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("deserializeYAML: %w", err)
	}
	return books, nil
}

func SerializeGOB(book []Book) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(book)
	return buff.Bytes(), err
}

func DeserializeGOB(data []byte) ([]Book, error) {
	var books []Book
	buff := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buff)
	err := dec.Decode(&books)
	return books, err
}

func SerializeMsgPack(book []Book) ([]byte, error) {
	return msgpack.Marshal(book)
}

func DeserializeMsgPack(data []byte) ([]Book, error) {
	var books []Book
	err := msgpack.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("deserializeMsgPack: %w", err)
	}
	return books, nil
}

func main() {}
