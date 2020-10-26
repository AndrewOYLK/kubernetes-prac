module github.com/AndrewOYLK/k8scode

go 1.15

require (
	github.com/imdario/mergo v0.3.11 // indirect
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/net v0.0.0-20201009032441-dbdefad45b89 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20201005171033-6301aaf42dc7 // indirect
)

replace k8s.io/client-go v11.0.0+incompatible => k8s.io/client-go v0.17.0
