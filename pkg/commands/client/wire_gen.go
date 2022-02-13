// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package client

import (
	"context"
	"github.com/aquasecurity/fanal/analyzer/config"
	"github.com/aquasecurity/fanal/artifact"
	image2 "github.com/aquasecurity/fanal/artifact/image"
	"github.com/aquasecurity/fanal/cache"
	"github.com/aquasecurity/fanal/image"
	"github.com/aquasecurity/fanal/types"
	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/aquasecurity/trivy/pkg/result"
	"github.com/aquasecurity/trivy/pkg/rpc/client"
	"github.com/aquasecurity/trivy/pkg/scanner"
)

// Injectors from inject.go:

func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache, customHeaders client.CustomHeaders, url client.RemoteURL, dockerOpt types.DockerOption, artifactOption artifact.Option, configScannerOption config.ScannerOption) (scanner.Scanner, func(), error) {
	scannerScanner := client.NewProtobufClient(url)
	clientScanner := client.NewScanner(customHeaders, scannerScanner)
	typesImage, cleanup, err := image.NewDockerImage(ctx, imageName, dockerOpt)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	artifactArtifact, err := image2.NewArtifact(typesImage, artifactCache, artifactOption, configScannerOption)
	if err != nil {
		cleanup()
		return scanner.Scanner{}, nil, err
	}
	scanner2 := scanner.NewScanner(clientScanner, artifactArtifact)
	return scanner2, func() {
		cleanup()
	}, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, artifactCache cache.ArtifactCache, customHeaders client.CustomHeaders, url client.RemoteURL, artifactOption artifact.Option, configScannerOption config.ScannerOption) (scanner.Scanner, error) {
	scannerScanner := client.NewProtobufClient(url)
	clientScanner := client.NewScanner(customHeaders, scannerScanner)
	typesImage, err := image.NewArchiveImage(filePath)
	if err != nil {
		return scanner.Scanner{}, err
	}
	artifactArtifact, err := image2.NewArtifact(typesImage, artifactCache, artifactOption, configScannerOption)
	if err != nil {
		return scanner.Scanner{}, err
	}
	scanner2 := scanner.NewScanner(clientScanner, artifactArtifact)
	return scanner2, nil
}

func initializeResultClient() result.Client {
	dbConfig := db.Config{}
	resultClient := result.NewClient(dbConfig)
	return resultClient
}