package handleTime

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func RetrieveNtpTime(host string) (time.Time, error) {
	fmt.Printf("Use '%s' as ntp server\n\n", host)
	return ntp.Time(host)
}