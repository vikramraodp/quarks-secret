module code.cloudfoundry.org/quarks-secret

go 1.15

require (
	code.cloudfoundry.org/quarks-utils v0.0.3-0.20210303091853-3b41f4b87e33
	github.com/cloudflare/cfssl v1.4.1
	github.com/dchest/uniuri v0.0.0-20200228104902-7aecb25e1fe5
	github.com/go-logr/logr v1.2.2
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.9.2
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	helm.sh/helm/v3 v3.3.0
	k8s.io/api v0.24.8
	k8s.io/apiextensions-apiserver v0.24.8
	k8s.io/apimachinery v0.24.8
	k8s.io/client-go v0.24.8
	sigs.k8s.io/controller-runtime v0.11.0
)

replace code.cloudfoundry.org/quarks-utils => ../quarks-utils
