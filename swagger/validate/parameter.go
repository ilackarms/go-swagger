package validate

// // Parameter creates a parameter validator
// func Parameter(param *swagger.Parameter) (*jsonschema.Document, error) {
// 	if param.In == "body" {
// 		return Schema(param.Schema)
// 	}
// 	b, err := param.MarshalJSON()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return loadSchema(b)
// }

// func loadSchema(schemaJSON []byte) (*jsonschema.Document, error) {
// 	var doc interface{}
// 	if err := json.Unmarshal(schemaJSON, &doc); err != nil {
// 		return nil, err
// 	}
// 	return jsonschema.New(doc)
// }

// // Schema creates a schema validator
// func Schema(schema *swagger.Schema) (*jsonschema.Document, error) {
// 	b, err := schema.MarshalJSON()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return loadSchema(b)
// }

// type parameterContext struct {
// 	request *http.Request
// 	route   *router.MatchedRoute
// 	model   map[string]interface{}
// 	param   *swagger.Parameter
// }

// // Parameter validates a request parameter
// func Parameter(request *http.Request, route *router.MatchedRoute, model map[string]interface{}, param *swagger.Parameter) []errors.Validation {
// 	return nil
// }
