package main

import "testing"

func TestSerialize(t *testing.T) {
	books := []Book{{ID: 1, Size: 200, Year: 1999, Rate: 4.8, Title: "Book1", Author: "Author1"}}

	// Test JSON
	t.Run("JSON test:", func(_ *testing.T) {
		jsonData, _ := SerializeJSON(books)
		DeserializeJSON(jsonData)
	})

	// Test XML
	t.Run("XML Test:", func(_ *testing.T) {
		xmlData, _ := SerializeXML(books)
		DeserializeXML(xmlData)
	})

	// Test YAML
	t.Run("YAML Test:", func(_ *testing.T) {
		yamlData, _ := SerializeYAML(books)
		DeserializeYAML(yamlData)
	})

	// Test GOB
	t.Run("GOB Test:", func(_ *testing.T) {
		gobData, _ := SerializeGOB(books)
		DeserializeGOB(gobData)
	})

	// Test MsgPack
	t.Run("MsgPack Test:", func(_ *testing.T) {
		msgPackData, _ := SerializeMsgPack(books)
		DeserializeMsgPack(msgPackData)
	})
}
