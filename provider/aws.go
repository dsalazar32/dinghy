package provider

import "fmt"

type Aws struct {
	Bucket string
	Dir    string
}

func init() {
	Constructors[AWS] = ProviderSpec{
		New:         NewAws,
		description: "Placeholder for Provider description.",
	}
}

func (p Aws) Connect() {
	fmt.Println("Amazon Web Services Provider (unimplemented!)")
}

func NewAws() (Provider, error) {
	return Aws{}, nil
}
