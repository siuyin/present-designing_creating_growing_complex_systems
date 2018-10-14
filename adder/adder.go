//10 OMIT
package adder

//20 OMIT

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//50 OMIT
func sumStr(a, b string) (int, error) {
	ai, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return ai + bi, nil
}

//60 OMIT
//30 OMIT
func AdderHandler(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("a")
	b := r.FormValue("b")
	sum, err := sumStr(a, b)
	if err != nil {
		fmt.Fprintf(w, "invalid number received: %v", err)
		return
	}
	fmt.Fprintf(w, "hello, sum is %d and the time is %s", sum,
		time.Now().
			In(time.FixedZone("SGT", 8*3600)).
			Format("2006-01-02 15:04:05.000 MST"))
}

//40 OMIT
