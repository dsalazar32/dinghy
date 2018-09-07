package provider

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
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

// TODO: Set all environmental variable configuration here.
func NewAws() (Provider, error) {
	return &Aws{bucket: "www.iomediums.com", dir: "static", prefix: ""}, nil
}

func (p *Aws) Run() error {
	fmt.Println("Amazon Web Services Provider!")
	sess := session.Must(session.NewSession())
	s3sv := s3.New(sess)

	s3In := &s3.ListObjectsV2Input{
		Bucket: &p.bucket,
		Prefix: &p.prefix,
	}

	// Set the download manager that will be used to download the objects.
	p.Downloader = s3manager.NewDownloader(sess)
	return s3sv.ListObjectsV2Pages(s3In, p.eachPage)
}

func (p Aws) eachPage(output *s3.ListObjectsV2Output, _ bool) bool {
	for _, item := range output.Contents {
		p.downloadToFile(*item.Key)
	}
	return true
}

func (p Aws) downloadToFile(key string) {
	file := filepath.Join(p.dir, key)
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
