//go:build e2e
// +build e2e

package e2e

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

const (
	defaultHelloOpenShiftImage = "openshift/hello-openshift:latest"
	defaultVaultImageRepo      = "docker.io/hashicorp/vault"
	defaultVaultImageTag       = "1.17.2"
	defaultGrpcurlImage        = "fullstorydev/grpcurl:v1.9.2-alpine"
	defaultHttpbinImage        = "docker.io/kennethreitz/httpbin"
	defaultCurlImage           = "curlimages/curl:8.5.0"
)

// e2eHelloOpenShiftImage returns the hello-openshift image to use for e2e tests.
func e2eHelloOpenShiftImage() string {
	if img := os.Getenv("E2E_HELLO_OPENSHIFT_IMAGE"); img != "" {
		return img
	}
	return defaultHelloOpenShiftImage
}

// e2eHelloOpenShiftImageForNS returns the hello-openshift image for a specific namespace.
func e2eHelloOpenShiftImageForNS(ns string) string {
	return e2eHelloOpenShiftImage()
}

// e2eVaultImageRepository returns the Vault container image repository.
func e2eVaultImageRepository() string {
	if repo := os.Getenv("E2E_VAULT_IMAGE_REPOSITORY"); repo != "" {
		return repo
	}
	return defaultVaultImageRepo
}

// e2eVaultImageTag returns the Vault container image tag.
func e2eVaultImageTag() string {
	if tag := os.Getenv("E2E_VAULT_IMAGE_TAG"); tag != "" {
		return tag
	}
	return defaultVaultImageTag
}

// e2eGrpcurlImage returns the grpcurl container image.
func e2eGrpcurlImage() string {
	if img := os.Getenv("E2E_GRPCURL_IMAGE"); img != "" {
		return img
	}
	return defaultGrpcurlImage
}

// e2eGrpcurlImageForNS returns the grpcurl container image for a specific namespace.
func e2eGrpcurlImageForNS(ns string) string {
	return e2eGrpcurlImage()
}

// e2eHttpbinImage returns the httpbin container image.
func e2eHttpbinImage() string {
	if img := os.Getenv("E2E_HTTPBIN_IMAGE"); img != "" {
		return img
	}
	return defaultHttpbinImage
}

// e2eHttpbinImageForNS returns the httpbin container image for a specific namespace.
func e2eHttpbinImageForNS(ns string) string {
	return e2eHttpbinImage()
}

// e2eCurlImage returns the curl container image.
func e2eCurlImage() string {
	if img := os.Getenv("E2E_CURL_IMAGE"); img != "" {
		return img
	}
	return defaultCurlImage
}

// e2eCurlImageForNS returns the curl container image for a specific namespace.
func e2eCurlImageForNS(ns string) string {
	return e2eCurlImage()
}

// loadVaultHelmValues renders testdata/vault/helm-values.yaml with configured repository and tag.
func loadVaultHelmValues(ns string) (string, error) {
	raw, err := testassets.ReadFile("testdata/vault/helm-values.yaml")
	if err != nil {
		return "", fmt.Errorf("read vault helm-values.yaml: %w", err)
	}
	tmpl, err := template.New("helm-values").Parse(string(raw))
	if err != nil {
		return "", fmt.Errorf("parse vault helm-values template: %w", err)
	}
	cfg := VaultHelmValuesConfig{
		Repository: e2eVaultImageRepository(),
		Tag:        e2eVaultImageTag(),
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("execute vault helm-values template: %w", err)
	}
	return buf.String(), nil
}
