package command

import (
	"github.com/eidolon/gonsole"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateCommand() gonsole.Command {
	var input string
	var schema string

	configure := func(d *gonsole.Definition) {
		d.Arg(
			gonsole.StringValue(&input),
			"INPUT",
			"The JSON to validate. Can be a JSON string, a local file (starting with 'file://'), " +
				"or a remote file served over HTTP (starting with 'http://').",
		)

		d.Arg(
			gonsole.StringValue(&schema),
			"SCHEMA",
			"The JSON schema. Can be a JSON string, a local file (starting with 'file://'), or a " +
				"remote file served over HTTP (starting with 'http://').",
		)
	}

	execute := func() int {
		inputLoader := gojsonschema.NewReferenceLoader(input)
		schemaLoader := gojsonschema.NewReferenceLoader(schema)

		result, err := gojsonschema.Validate(schemaLoader, inputLoader)
		if err != nil {
			fmt.Println(err)
			return 1
		}

		if !result.Valid() {
			fmt.Println("The input appears to be invalid. Here are the problems:")

			// @todo: This should show the actual path to the node with the problem. Not just the
			// @todo: field name of the problem.
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}

			return 1
		}

		fmt.Println("The input appears to be valid!")

		return 0
	}

	return gonsole.Command{
		Name: "validate",
		Description: "Validate some JSON against some JSON schema.",
		Configure: configure,
		Execute: execute,
	}
}
