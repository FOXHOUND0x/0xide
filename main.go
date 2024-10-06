package main

import (
	"ragnarok/modules/docker"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var images = []struct {
	tag  string
	path string
}{
	{"traefik-local:2.10", "./dockerfiles/traefik"},
	{"consul-local:1.15.4", "./dockerfiles/consul"},
	{"vault-local:1.13.3", "./dockerfiles/vault"},
	{"prometheus-local:v2.55.0-rc.0", "./dockerfiles/prometheus"},
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		for _, image := range images {
			if _, err := docker.BuildDockerImage(ctx, image.tag, image.path); err != nil {
				return err
			}
		}

		_, err := docker.StartConsulContainer(ctx, "consul-local:1.15.4")
		if err != nil {
			return err
		}

		_, err = docker.StartTraefikContainer(ctx, "traefik-local:2.10", nil)
		if err != nil {
			return err
		}

		_, err = docker.StartVaultContainer(ctx, "vault-local:1.13.3", nil)
		if err != nil {
			return err
		}

		_, err = docker.StartPrometheusContainer(ctx, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
