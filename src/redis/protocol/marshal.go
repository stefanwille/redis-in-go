package protocol

import (
	"fmt"
	"io"
)

func Marshal(writer io.Writer, value Any) error {
	switch t := value.(type) {
	case int:
		var intValue int = value.(int)
		fmt.Fprintf(writer, ":%d\r\n", intValue)
		return nil

	case string:
		var stringValue string = value.(string)
		fmt.Fprintf(writer, "$%d\r\n%s\r\n", len(stringValue), stringValue)
		return nil

	case *[]Any:
		var arrayPointerValue *[]Any = value.(*[]Any)
		fmt.Fprintf(writer, "*%d\r\n", len(*arrayPointerValue))
		for _, anyValue := range *arrayPointerValue {
			error := Marshal(writer, anyValue)
			if error != nil {
				return error
			}
		}
		return nil

	case []Any:
		var arrayValue []Any = value.([]Any)
		return Marshal(writer, &arrayValue)

	default:
		return fmt.Errorf("Unrecognized type '%T' for redis.protocol.Marshal", t)
	}
}

func MarshalSimpleString(writer io.Writer, value string) {
	fmt.Fprintf(writer, "+%s\r\n", value)
}
