package dump

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func HttpResponseDump(res *http.Response) string {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println(err)
	}
	return string(dump)
}
