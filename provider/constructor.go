package provider

type ProviderSpec struct {
	new         func() (Provider, error)
	description string
}

var Constructors = map[string]ProviderSpec{}

const (
	AWS = "awsS3"
	GCP = "gcpCloudStorage"
)
