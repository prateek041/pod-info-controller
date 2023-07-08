package controllertest

import "embed"

//go:embed config/crd/bases/test.prateeksingh.tech_podreplicators.yaml
var CRDv1TetragonPod embed.FS
