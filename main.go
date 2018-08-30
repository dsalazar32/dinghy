package main

import (
	"github.com/dsalazar32/dinghy/provider"
	"os"
)

const nginxStaticDir = "/usr/share/nginx/html"

var storage provider.Provider

func main() {
	bucket, ok := os.LookupEnv("BUCKET")
	if !ok {
		panic("environment variable 'BUCKET' is required.")
	}

	pSvc, ok := os.LookupEnv("PROVIDER")
	if !ok {
		pSvc = "aws"
	}

	sDir, ok := os.LookupEnv("STATIC_DIR")
	if !ok {
		sDir = nginxStaticDir
	}

	switch pSvc {
	case "aws":
		storage = provider.Aws{bucket, sDir}
	case "gcp":
		storage = provider.Gcp{bucket, sDir}
	}

	storage.Download()
}
