package provider

// const nginxStaticDir = "/usr/share/nginx/html"

type Provider interface {
	Run() error
}

func DownloadFrom(pSpec ProviderSpec) error {
	p, err := pSpec.new()
	if err != nil {
		return err
	}

	return p.Run()
}
