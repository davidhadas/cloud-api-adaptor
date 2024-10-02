package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
	"time"
)

func main() {
	finish := make(chan bool)

	var wg sync.WaitGroup

	fmt.Printf("CAA LOG: STARTING\n")
	ctx, cancel := context.WithCancel(context.Background())
	caaLogTailCmd := exec.CommandContext(ctx, "kubectl", "logs", "-l", "app=cloud-api-adaptor", "-n", "confidential-containers-system")
	fmt.Printf("CAA LOG: STARTING 2\n")
	//caaLogTailCmd.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG="+cfg.KubeconfigFile()))
	stdout, err := caaLogTailCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CAA LOG: STARTING 3\n")

	if err := caaLogTailCmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CAA LOG: STARTING 4\n")

	wg.Add(1)
	go func() {
		defer cancel()

		defer wg.Done()
		fmt.Printf("GO CAA LOG: STARTING 1\n")
		rd := bufio.NewReader(stdout)
		fmt.Printf("GO CAA LOG: STARTING 2\n")
		for {
			line, err := rd.ReadString('\n')
			fmt.Printf("GO CAA LOG: STARTING 3\n")

			if err == nil {
				fmt.Printf("CAA LOG: %s\n", line)
				continue
			}
			if err != io.EOF {
				fmt.Printf("CAA LOG: exit with error: %v", err)
				return
			}
			fmt.Printf("GO CAA LOG: STARTING 4\n")

			select {
			case <-finish:
				fmt.Printf("CAA LOG: finished\n")
				return
			case <-time.After(time.Second):
				continue
			}
		}
	}()
	fmt.Printf("CAA LOG: STARTING 5\n")

	wg.Wait()
	fmt.Printf("CAA LOG: STARTING 6\n")

}
