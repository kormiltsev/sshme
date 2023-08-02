# SSHMe

Run commands on remote server. There is more simple way, for sure) Like this one:

`ssh user@server command`

### Usage:

```
sshme -u username -ip 255.255.255.0:22 -k ./path/to/key -exec top
```

You can use .env (in the same folder)
```
SSHME_IP="255.255.255.0:22"
SSHME_USER="root"
SSHME_KEY="/Users/user/.ssh/key_rsa"
```

Usage in this case 
```
sshme -exec top
```

In case of using ENV: 
```
export SSHME_IP="255.255.255.0:22"
export SSHME_USER="root"
export SSHME_KEY="/Users/user/.ssh/key_rsa"
```

Usage in this case 
```
sshme -exec top
```

### Example:

run 
```
go run sshme -u root -ip 255.255.255.0:22 -k /Users/user/.ssh/key_rsa
```

In case of empty flag -exec. Command to execute is set in main.go:
```
command := `echo "===================== RAM ====================="; free -h | awk 'NR==1 {print $1 "  " $2 "  " $3 " " $4 " " $5}'; free -h | awk 'NR==2 {print $2 " " $3 "  " $4 "  " $5 " " $6}'; echo "==================== DRIVES ==================="; df -H`
```

Answer:
```
===================== RAM =====================
total  used  free shared buff/cache
968Mi 252Mi  75Mi  0.0Ki 640Mi
==================== DRIVES ===================
Filesystem      Size  Used Avail Use% Mounted on
tmpfs           102M  717k  101M   1% /run
/dev/vda2        63G   60G  520M 100% /
tmpfs           508M     0  508M   0% /dev/shm
tmpfs           5.3M     0  5.3M   0% /run/lock
/dev/vda1       563M  5.6M  558M   1% /boot/efi
tmpfs           102M     0  102M   0% /run/user/0
tmpfs           102M     0  102M   0% /run/user/1001
```
### Build examples:

64bit

Linux:
`GOOS=linux GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`

Windows:
`GOOS=windows GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`

MacOS:
`GOOS=darwin GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`