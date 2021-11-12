package types

import "strings"

func wrapperToPrimitive(wrapperType GoType) string {
	switch wrapperType {
	case "wrapperspb.String":
		return "string"
	case "wrapperspb.Bool":
		return "bool"
	case "wrapperspb.Float":
		return "float64"
	case "wrapperspb.Int32":
		return "int"
	case "wrapperspb.Int64":
		return "int"
	}
	return ""
}

func stripPrecision(arg GoType) string {
	if strings.Contains(string(arg), "int") {
		output := strings.Replace(string(arg), "64", "", -1)
		return strings.Replace(output, "32", "", -1)
	}

	if strings.Contains(string(arg), "float") {
		return "float64"
	}
	return string(arg)
}

func primitiveToWrapper(wrapperType GoType) string {
	switch wrapperType {
	case "wrapperspb.Float":
		return "float32"
	case "wrapperspb.String":
		return "string"
	case "wrapperspb.Bool":
		return "bool"
	case "wrapperspb.Int32":
		return "int32"
	case "wrapperspb.Int64":
		return "int64"
	}
	return string(wrapperType)
}
