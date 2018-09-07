package main

import (
	"fmt"
	"github.com/dsalazar32/dinghy/provider"
	"log"
	"os"
)

func main() {
	// default storage cloud provider to aws s3
	p, ok := os.LookupEnv("PROVIDER")
	if !ok {
		p = provider.AWS
	}

	storage, ok := provider.Constructors[p]
	if !ok {
		fmt.Println("Error Occurred")
	}

	if err := provider.DownloadFrom(storage); err != nil {
		log.Fatalf("execution failed: %v", err)
	}
}
