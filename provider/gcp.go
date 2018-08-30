package provider

import "fmt"

type Gcp struct {
	Bucket string
	Dir    string
}

func (p Gcp) Download() {
	fmt.Println("GCP Download (unimplemented!)")
}
