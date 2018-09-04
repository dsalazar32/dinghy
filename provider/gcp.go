package provider

import "fmt"

type Gcp struct {
	Bucket string
	Dir    string
}

func init() {
	Constructors[GCP] = ProviderSpec{
		New:         NewGcp,
		description: "Placeholder for Provider Description",
	}
}

func (p Gcp) Connect() {
	fmt.Println("Google Cloud Platform Provider (unimplemented!)")
}

func NewGcp() (Provider, error) {
	return Gcp{}, nil
}
