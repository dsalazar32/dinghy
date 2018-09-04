package provider

type ProviderSpec struct {
	New         func() (Provider, error)
	description string
}

var Constructors = map[string]ProviderSpec{}

const (
	AWS = "awsS3"
	GCP = "gcpCloudStorage"
)
