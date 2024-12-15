package template

import (
	"pg-backend/util"
	"text/template"
	"encoding/json"
	"reflect"
	"bytes"
	"io"
	"os"
)

type Schema = util.Schema
type QuestionSchema = util.QuestionSchema
type Legend = map[string]QuestionSchema

func readSchema() Schema {
	file, err := os.Open("./data/schema.json")

	if err != nil {
		panic("Unable to open the file.")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("Unable to read the file.")
	}

	var data Schema
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic("Unable to parse json data")
	}

	return data
}

func buildTemplate(data QuestionSchema) string {
	tpl, err := template.ParseFiles("./template/question.tmpl")

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	tpl.Execute(&buf, data)

	return buf.String()
}

func generateLegendInitialization(schema Schema) Legend {
	var legend Legend = make(Legend)

	schemaValue := reflect.ValueOf(schema)
	schemaType := reflect.TypeOf(schema)

	for i := 0; i < schemaType.NumField(); i++ {
		category := schemaValue.Field(i)

		for j := 0; j < category.NumField(); j++ {
			technique := category.Field(j)
			id := technique.FieldByName("Id").String()
			questionSchema := category.Field(j).Interface().(QuestionSchema)
			legend[id] = questionSchema
		}
	}

	return legend
}

func Generate(id string) string {
	var schema Schema = readSchema()
	legend := generateLegendInitialization(schema)
	return buildTemplate(legend[id])
}
