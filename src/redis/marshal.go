package redis

import (
	"fmt"
	"io"
)

type Any interface{}

func Marshal(writer io.Writer, value Any) {
	switch t := value.(type) {
	case int:
		var intValue int = value.(int)
		fmt.Fprintf(writer, ":%d\r\n", intValue)
	case string:
		var stringValue string = value.(string)
		fmt.Fprintf(writer, "$%d\r\n%s\r\n", len(stringValue), stringValue)
	case *[]Any:
		var arrayPointerValue *[]Any = value.(*[]Any)
		fmt.Fprintf(writer, "*%d\r\n", len(*arrayPointerValue))
		for _, anyValue := range *arrayPointerValue {
			Marshal(writer, anyValue)
		}
	case []Any:
		var arrayValue []Any = value.([]Any)
		Marshal(writer, &arrayValue)
	default:
		panic(fmt.Sprintf("Unrecognized type '%T' for redis.Marshal", t))
	}
}

func MarshalSimpleString(writer io.Writer, value string) {
	fmt.Fprintf(writer, "+%s\r\n", value)
}
