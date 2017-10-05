/*
* Copyright © 2017. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package app

var schemaMap = map[string]string{
	"0.1": "schema/mashling_schema.json",
	"0.2": "schema/mashling_schema.json",
}

//GetSupportedSchema returns supported schema
func GetSupportedSchema(schemaVal string) (string, bool) {
	supportedSchema, ok := schemaMap[schemaVal]
	return supportedSchema, ok
}
