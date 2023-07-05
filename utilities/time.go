package utilities

import (
	"time"
)

//
func DateTimeNow() time.Time {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	return time.Now().UTC().Local()

}
