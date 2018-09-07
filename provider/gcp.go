package provider

import "fmt"

type Gcp struct {
	Bucket string
	Dir    string
}

func init() {
	Constructors[GCP] = ProviderSpec{
		new:         NewGcp,
		description: "Placeholder for Provider Description",
	}
}

func (p Gcp) Run() error {
	fmt.Println("Google Cloud Platform Provider (unimplemented!)")
	return nil
}

func NewGcp() (Provider, error) {
	return Gcp{}, nil
}
