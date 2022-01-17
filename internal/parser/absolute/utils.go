package absolute

import (
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func print(args ...interface{}) {
	// fmt.Println(args...)
}

func jsonify(args interface{}) string {
	return strutil.Jsonify(args)
}
