package docker

import (
	"fmt"

	"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type containerArgs struct {
	Image  pulumi.StringInput
	Ports  docker.ContainerPortArrayInput
	Envs   pulumi.StringArrayInput
	Labels docker.ContainerLabelArrayInput
}

func newContainer(ctx *pulumi.Context, name string, args *containerArgs, dependsOn []pulumi.Resource) (*docker.Container, error) {
	return docker.NewContainer(ctx, name, &docker.ContainerArgs{
		Image:  args.Image,
		Ports:  args.Ports,
		Envs:   args.Envs,
		Labels: args.Labels,
	}, pulumi.DependsOn(dependsOn))
}

func StartConsulContainer(ctx *pulumi.Context, imageName string) (*docker.Container, error) {
	args := &containerArgs{
		Image: pulumi.String(imageName),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(8500),
				External: pulumi.Int(8500),
			},
		},
		Envs: pulumi.StringArray{
			pulumi.String("CONSUL_BIND_INTERFACE=eth0"),
		},
		Labels: getTraefikLabels(8500),
	}
	return newContainer(ctx, "consul-container", args, nil)
}

func StartTraefikContainer(ctx *pulumi.Context, imageName string, dependsOn []pulumi.Resource) (*docker.Container, error) {
	args := &containerArgs{
		Image: pulumi.String(imageName),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(80),
				External: pulumi.Int(80),
			},
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(8080),
				External: pulumi.Int(8080),
			},
		},
	}
	return newContainer(ctx, "traefik-container", args, dependsOn)
}

func StartVaultContainer(ctx *pulumi.Context, imageName string, dependsOn []pulumi.Resource) (*docker.Container, error) {
	args := &containerArgs{
		Image: pulumi.String(imageName),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(8200),
				External: pulumi.Int(8200),
			},
		},
		Envs: pulumi.StringArray{
			pulumi.String("VAULT_DEV_ROOT_TOKEN_ID=root"),
		},
		Labels: getTraefikLabels(8200),
	}
	return newContainer(ctx, "vault-container", args, dependsOn)
}

func StartPrometheusContainer(ctx *pulumi.Context, dependsOn []pulumi.Resource) (*docker.Container, error) {
	args := &containerArgs{
		Image: pulumi.String("prom/prometheus:v2.55.0-rc.0"),
		Ports: docker.ContainerPortArray{
			&docker.ContainerPortArgs{
				Internal: pulumi.Int(9091),
				External: pulumi.Int(9091),
			},
		},
	}
	return newContainer(ctx, "prometheus-container", args, dependsOn)
}

func getTraefikLabels(port int) docker.ContainerLabelArray {
	return docker.ContainerLabelArray{
		&docker.ContainerLabelArgs{
			Label: pulumi.String("traefik.enable"),
			Value: pulumi.String("true"),
		},
		&docker.ContainerLabelArgs{
			Label: pulumi.String("traefik.http.routers"),
			Value: pulumi.String(fmt.Sprintf("Host(`%s.localhost.`)", getHostName(port))),
		},
		&docker.ContainerLabelArgs{
			Label: pulumi.String("traefik.http.services.loadbalancer.server.port"),
			Value: pulumi.String(fmt.Sprintf("%d", port)),
		},
	}
}

func getHostName(port int) string {
	switch port {
	case 8500:
		return "consul"
	case 8200:
		return "vault"
	default:
		return ""
	}
}
