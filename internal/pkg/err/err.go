package err

import (
	"errors"
	"strings"
)

func JavaScriptError(str string) error {
	if strings.Contains(str, "请开启JavaScript并刷新该页") {
		return errors.New("请开启JavaScript并刷新该页")
	}
	return nil
}