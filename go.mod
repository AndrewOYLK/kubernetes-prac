module github.com/AndrewOYLK/kubernetes-prac

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-openapi/spec v0.19.6
	github.com/go-redis/redis/v8 v8.3.3
	github.com/minio/minio-go/v7 v7.0.5
	github.com/operator-framework/operator-sdk v0.14.0
	github.com/redhat-cop/operator-utils v0.1.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5
	github.com/tektoncd/pipeline v0.18.0
	github.com/thedevsaddam/gojsonq v2.3.0+incompatible // indirect
	github.com/urfave/cli v1.22.2
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
	k8s.io/metrics v0.18.6
	sigs.k8s.io/controller-runtime v0.6.1
	sigs.k8s.io/controller-tools v0.1.10
)

// Pinned to kubernetes-1.13.4
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190222213804-5cb15d344471
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190228180357-d002e88f6236
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221213512-86fb29eff628
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228174230-b40b2a5939e4
)

replace (
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.29.0
	k8s.io/kube-state-metrics => k8s.io/kube-state-metrics v1.6.0
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.1.12
	sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.1.11-0.20190411181648-9d55346c2bde
)

replace github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.9.0

replace knative.dev/pkg => knative.dev/pkg v0.0.0-20200113182502-b8dc5fbc6d2f

go 1.13
