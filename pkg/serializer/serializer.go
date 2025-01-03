package rest

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"io"
)

type UpdateHelper interface {
	GetFields() []string
	SetFields([]string)
}

type EasyJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d EasyJSONSerializer) Serialize(c echo.Context, i any, indent string) error {
	switch m := i.(type) {
	case easyjson.Marshaler:
		_, err := easyjson.MarshalToWriter(m, c.Response())
		return err
	default:
		e := json.NewEncoder(c.Response())
		e.SetIndent("", indent)
		return e.Encode(m)
	}
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d EasyJSONSerializer) Deserialize(c echo.Context, i any) error {
	var (
		err error
		ok  bool
	)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	_, ok = i.(UpdateHelper)
	if ok {
		data := make(map[string]any)
		if err = json.Unmarshal(body, &data); err != nil {
			return err
		}

		var updatedFields []string
		for entKey := range data {
			// main entity object
			switch e := data[entKey].(type) {
			case map[string]any:
				updatedFields = make([]string, 0, len(e))
				// normalize
				for j := range e {
					switch ej := e[j].(type) {
					case map[string]any:
						for k := range ej {
							// transform something like data[manufacturer]any{key: val} to data[manufacturer_key]string
							e[fmt.Sprintf("%s_%s", j, k)] = ej[k]
							delete(e, j)
							continue
						}
					}
				}
				// fill updated fields
				for j := range e {
					updatedFields = append(updatedFields, j)
				}
			}
		}
		i.(UpdateHelper).SetFields(updatedFields)
	}

	_, ok = i.(easyjson.Unmarshaler)
	if !ok {
		return fmt.Errorf("easyjson: unsupported type %T", i)
	}

	return easyjson.Unmarshal(body, i.(easyjson.Unmarshaler))
}
