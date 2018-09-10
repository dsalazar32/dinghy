package provider

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
	"strings"
)

type Aws struct {
	*s3manager.Downloader

	bucket string
	dir    string
	prefix string
}

func init() {
	Constructors[AWS] = ProviderSpec{
		new:         NewAws,
		description: "Placeholder for Provider description.",
	}
}

func NewAws() (Provider, error) {
	// Bucket where objects will be downloaded from.
	bucket, ok := os.LookupEnv("AWS_BUCKET")
	if !ok {
		panic("environment variable 'AWS_BUCKET' is required.")
	}

	// Limits the response to keys that begin with the specified prefix.
	prefix := os.Getenv("AWS_BUCKET_PREFIX")

	// Target directory objects will be downloaded to.
	dldir, ok := os.LookupEnv("DOWNLOAD_DIR")
	if !ok {
		dldir = nginxStaticDir
	}

	return &Aws{
		bucket:     bucket,
		dir:        dldir,
		prefix:     prefix,
		Downloader: s3manager.NewDownloader(session.Must(session.NewSession())),
	}, nil
}

func (p *Aws) Run() error {
	fmt.Println("Amazon Web Services Provider!")
	sess := session.Must(session.NewSession())
	s3sv := s3.New(sess)

	s3In := &s3.ListObjectsV2Input{
		Bucket: &p.bucket,
		Prefix: &p.prefix,
	}

	return s3sv.ListObjectsV2Pages(s3In, p.eachPage)
}

func (p Aws) eachPage(output *s3.ListObjectsV2Output, _ bool) bool {
	for _, item := range output.Contents {
		p.downloadToFile(*item.Key)
	}
	return true
}

func (p Aws) downloadToFile(key string) {
	file := filepath.Join(p.dir, strings.TrimPrefix(key, p.prefix))
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		panic(err)
	}

	fd, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", p.bucket, key, file)
	s3In := &s3.GetObjectInput{
		Bucket: &p.bucket,
		Key:    &key,
	}
	p.Download(fd, s3In)
}
