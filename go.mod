module ca-vuln-scan

go 1.18

require (
	github.com/anchore/grype v0.35.0
	github.com/armosec/armoapi-go v0.0.99
	github.com/armosec/cluster-container-scanner-api v0.0.36
	github.com/armosec/logger-go v0.0.11
	github.com/armosec/utils-go v0.0.7
	github.com/armosec/utils-k8s-go v0.0.7
	github.com/docker/docker v20.10.14+incompatible
	github.com/golang/glog v1.0.0
	github.com/google/go-cmp v0.5.8
	github.com/hashicorp/go-multierror v1.1.1
	github.com/stretchr/testify v1.7.1
	github.com/xyproto/randomstring v0.0.0-20220311122621-4964ec363205
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9
)

require (
	github.com/CycloneDX/cyclonedx-go v0.5.0 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/acobaugh/osrelease v0.1.0 // indirect
	github.com/anchore/go-macholibre v0.0.0-20220308212642-53e6d0aaf6fb // indirect
	github.com/anchore/go-rpmdb v0.0.0-20210914181456-a9c52348da63 // indirect
	github.com/anchore/go-version v1.2.2-0.20210903204242-51efa5b487c4 // indirect
	github.com/anchore/packageurl-go v0.1.1-0.20220314153042-1bcd40e5206b // indirect
	github.com/anchore/stereoscope v0.0.0-20220406160859-c03a18a6b270 // indirect
	github.com/anchore/syft v0.44.0 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/bmatcuk/doublestar/v2 v2.0.4 // indirect
	github.com/bmatcuk/doublestar/v4 v4.0.2 // indirect
	github.com/containerd/containerd v1.5.10 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.10.1 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/cli v20.10.12+incompatible // indirect
	github.com/docker/distribution v2.8.0+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dsnet/compress v0.0.2-0.20210315054119-f66993602bf5 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/facebookincubator/nvdtools v0.1.4 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/gabriel-vasile/mimetype v1.4.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-restruct/restruct v1.2.0-alpha // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-containerregistry v0.8.1-0.20220209165246-a44adc326839 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/jinzhu/copier v0.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/knqyf263/go-apk-version v0.0.0-20200609155635-041fdbb8563f // indirect
	github.com/knqyf263/go-deb-version v0.0.0-20190517075300-09fca494f03d // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mholt/archiver/v3 v3.5.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20220114050600-8b9d41f48198 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/scylladb/go-set v1.0.3-0.20200225121959-cc7b2070d91e // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spdx/tools-golang v0.2.0 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vbatts/tar-split v0.11.2 // indirect
	github.com/vifraa/gopom v0.1.0 // indirect
	github.com/wagoodman/go-partybus v0.0.0-20210627031916-db1f5573bbc5 // indirect
	github.com/wagoodman/go-progress v0.0.0-20200807221327-51d465df1451 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	golang.org/x/crypto v0.0.0-20220213190939-1e6e3497d506 // indirect
	golang.org/x/mod v0.6.0-dev.0.20211013180041-c96bc1413d57 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220218161850-94dd64e39d7c // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.23.3 // indirect
	k8s.io/klog/v2 v2.40.1 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)
