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

### Build examples:

64bit

Linux:
`GOOS=linux GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`

Windows:
`GOOS=windows GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`

MacOS:
`GOOS=darwin GOARCH=amd64 go build -o bin/sshme ./cmd/main.go`