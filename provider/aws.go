package provider

import "fmt"

type Aws struct {
	Bucket string
	Dir    string
}

func (p Aws) Download() {
	fmt.Println("AWS Download (unimplemented!)")
}
