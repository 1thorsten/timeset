# timeset
Utility to set time and date from an NTP server (windows)

# Build (for Windows only)
```bash
GOOS=windows GOARCH=amd64 go build -o timeset.exe timeset.go
```
or

```bash
GOOS=windows GOARCH=amd64 go install
```

# Introduction
timeset will help you if your computer is not able to keep the current time (CMOS battery has run out of power).
Of course, you can buy a new battery and replace it. My tip: Save your money and time and use timeset

# Usage
1. take the standard ntp server (pool.ntp.org)
```bash
timeset.exe 
```
2. specify ntp server as first argument
```bash
timeset.exe pool.ntp.org
# if your router acts as a time server for your local network
timeset.exe 10.10.1.1
```

3. Specify the time server in the name of the executable
```bash
timeset_ntp_pool_org.exe
# if your router acts as a time server for your local network
timeset_10_10_1_1.exe
```

!! timeset must be started with administrator rights (otherwise system time and date cannot be set) .!!!
![set admin rights](https://github.com/1thorsten/timeset/blob/master/run_as_admin.png)

# Download
You can also just download the latest release from [here](https://github.com/1thorsten/timeset/releases).
