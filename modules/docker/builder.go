package docker

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func BuildDockerImage(ctx *pulumi.Context, imageName, contextPath string) (*docker.Image, error) {
	exists, err := imageExists(imageName)
	if err != nil {
		return nil, err
	}

	if exists {
		log.Printf("Image %s already exists. Skipping Build", imageName)
		return nil, nil
	}

	image, err := docker.NewImage(ctx, imageName, &docker.ImageArgs{
		ImageName: pulumi.String(imageName),
		Build: &docker.DockerBuildArgs{
			Context: pulumi.String(contextPath),
		},
		SkipPush: pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	return image, nil
}

func imageExists(imageName string) (bool, error) {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	output, err := cmd.CombinedOutput()

	if err != nil {
		if strings.Contains(string(output), "No such image") {
			return false, nil
		}
		return false, fmt.Errorf("error checking image existence: %w", err)
	}

	return true, nil
}
