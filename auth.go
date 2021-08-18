package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rpc "github.com/thepagan/bzemsg/rpc"
)

func init() {
	u, p, err := readAuthCreds()
	if err != nil {
		fmt.Println("Error reading bzedge config: ", err)
	}

	rpc.DefaultClient.User = u
	rpc.DefaultClient.Pass = p
}

func readAuthCreds() (string, string, error) {
	homedir := os.Getenv("HOME")
	confpath := filepath.Join(homedir, ".bzedge/bzedge.conf")
	fi, err := os.Open(confpath)
	if err != nil {
		return "", "", err
	}
	defer fi.Close()

	var user string
	var pass string
	scan := bufio.NewScanner(fi)
	for scan.Scan() {
		parts := strings.SplitN(scan.Text(), "=", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "rpcuser":
			user = val
		case "rpcpassword":
			pass = val
		}
	}

	return user, pass, nil
}
