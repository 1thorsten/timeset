package main

import (
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
	"timeset/handleTime"
)

const (
	fallbackNtpHost = "pool.ntp.org"
	defaultFilename = "timeset"
)

// DetermineNtpHost by
// 1. take the given argument -> will not be checked
// 2. try to determine ntp server by the executable (timeset_ntp_pool_org.exe -> ntp.pool.org) - will be checked (if ip or host is not valid the server (host string) is taken)
// 3. if nothing fits, the server (host string) is taken
func DetermineNtpHost(host string) string {
	if len(os.Args) >= 2 {
		host = os.Args[1]
		fmt.Printf("Ntp-Server by argument: %s\n", host)
		return host
	}
	_, file, _, ok := runtime.Caller(0)

	if ok {
		// timeset.go
		base := path.Base(file)
		// timeset
		name := base[:len(base)-3]
		// timeset_192_168_178_1.exe
		exeName := os.Args[0]
		// exeName = "timeset_pool_ntp_or.exe"
		suffix := path.Ext(exeName)

		// look for time_set_ (e.g. timeset_192_168_178_1.exe)
		if index := strings.LastIndex(exeName, name+"_"); index != -1 {
			// 192_168_178_1
			potentialMaskedIpOrHost := exeName[index+len(name)+1 : len(exeName)-len(suffix)]
			// replace _ with .
			potentialMaskedIpOrHost = strings.Replace(potentialMaskedIpOrHost, "_", ".", -1)
			// too few dots -> take default host
			if strings.Count(potentialMaskedIpOrHost, ".") < 1 {
				return HostWithFormatInformation(name, host)
			}
			if net.ParseIP(potentialMaskedIpOrHost) == nil {
				if _, err := net.LookupHost(potentialMaskedIpOrHost); err != nil {
					fmt.Printf("no valid ip address and no valid host name: %s\n", potentialMaskedIpOrHost)
					return host
				}
			}
			fmt.Printf("Ntp-Server by name of the executable: %s\n", potentialMaskedIpOrHost)
			return potentialMaskedIpOrHost
		}
	}
	return HostWithFormatInformation(defaultFilename, host)
}

func HostWithFormatInformation(baseName string, host string) string {
	fmt.Println("Specify time server by")
	fmt.Printf("- argument '%s [ntp-server - e.g. pool.ntp.org]'\n", os.Args[0])
	fmt.Printf("- exe-name '%s_10_10_0_1.exe\n", baseName)
	fmt.Printf("- exe-name '%s_pool_nt_org.exe\n", baseName)
	return host
}

const VERSION = "1.0.1"

func main() {
	fmt.Printf("timeset version: %s\n", VERSION)

	if operatingSystem := runtime.GOOS; operatingSystem != "windows" {
		fmt.Printf("invalid os (%s). Only windows is allowed.\n", operatingSystem)
		os.Exit(1)
	}

	ntpTime, err := handleTime.RetrieveNtpTime(DetermineNtpHost(fallbackNtpHost))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err = handleTime.SetSystemDate(ntpTime); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := handleTime.SetSystemTime(ntpTime); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
