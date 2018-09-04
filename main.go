package main

import (
	"fmt"
	"github.com/dsalazar32/dinghy/provider"
	"os"
)

const nginxStaticDir = "/usr/share/nginx/html"

func main() {
	bucket, ok := os.LookupEnv("BUCKET")
	if !ok {
		panic("environment variable 'BUCKET' is required.")
	}
	fmt.Println(bucket)

	pSvc, ok := os.LookupEnv("PROVIDER")
	if !ok {
		pSvc = "aws"
	}
	fmt.Println(pSvc)

	sDir, ok := os.LookupEnv("STATIC_DIR")
	if !ok {
		sDir = nginxStaticDir
	}
	fmt.Println(sDir)

	storage, err := provider.Constructors[pSvc].New()
	if err != nil {
		fmt.Println("Error Occurred")
	}

	storage.Connect()
}
