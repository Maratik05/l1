package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	xmlData := []byte(`
        <data>
            <name>Product</name>
            <value>100$</value>
        </data>
    `)
	xmlParser := &XMLParser{}
	adapter := NewAdapterJSON(xmlParser)
	res, _ := adapter.ParseJSON(xmlData)

	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(jsonBytes))
}

type Parser interface {
	ParseJSON(data []byte) (map[string]interface{}, error)
}

type XMLParser struct{}

func (p *XMLParser) ParseXML(data []byte) (XMLData, error) {
	var res XMLData
	err := xml.Unmarshal(data, &res)
	return res, err
}

type XMLData struct {
	XMLName xml.Name `xml:"data"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type AdapterJSON struct {
	xmlParser *XMLParser
}

func NewAdapterJSON(xmlParser *XMLParser) *AdapterJSON {
	return &AdapterJSON{xmlParser: xmlParser}
}

func (aj *AdapterJSON) ParseJSON(data []byte) (map[string]interface{}, error) {
	xmldata, err := aj.xmlParser.ParseXML(data)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"name":  xmldata.Name,
		"value": xmldata.Value,
	}

	return result, nil
}
