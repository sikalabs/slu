package digitalocean_utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/digitalocean/godo"
)

func GetToken() string {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		log.Fatal("environment variable DIGITALOCEAN_TOKEN is missing or empty")
	}
	return token
}

func listKubernetes(client *godo.Client) {
	clusters, _, _ := client.Kubernetes.List(context.TODO(), &godo.ListOptions{})
	if len(clusters) == 0 {
		return
	}
	fmt.Println("Kubernetes")
	for _, cluster := range clusters {
		fmt.Println(cluster.Name)
	}
	fmt.Println("")

}

func listDroplets(client *godo.Client) {
	droplets, _, _ := client.Droplets.List(context.TODO(), &godo.ListOptions{})
	if len(droplets) == 0 {
		return
	}
	fmt.Println("Droplets")
	for _, droplet := range droplets {
		fmt.Println(droplet.Name)
	}
	fmt.Println("")
}

func listLoadBalancers(client *godo.Client) {
	loadbalancers, _, _ := client.LoadBalancers.List(context.TODO(), &godo.ListOptions{})
	if len(loadbalancers) == 0 {
		return
	}
	fmt.Println("LoadBalancers")
	for _, lb := range loadbalancers {
		fmt.Println(lb.Name)
	}
	fmt.Println("")
}

func listVolumes(client *godo.Client) {
	volumes, _, _ := client.Storage.ListVolumes(context.TODO(), &godo.ListVolumeParams{})
	if len(volumes) == 0 {
		return
	}
	fmt.Println("Volumes")
	for _, volume := range volumes {
		fmt.Println(volume.Name, volume.DropletIDs)
	}
	fmt.Println("")
}

func listDomains(client *godo.Client) {
	domains, _, _ := client.Domains.List(context.TODO(), &godo.ListOptions{})
	if len(domains) == 0 {
		return
	}
	fmt.Println("Domains")
	for _, domain := range domains {
		fmt.Println(domain.Name)
	}
	fmt.Println("")
}

func listSSHKeys(client *godo.Client) {
	keys, _, _ := client.Keys.List(context.TODO(), &godo.ListOptions{})
	if len(keys) == 0 {
		return
	}
	fmt.Println("Keys (SSH)")
	for _, key := range keys {
		fmt.Println(key.Name)
	}
	fmt.Println("")
}

func ListAll(token string) {
	client := godo.NewFromToken(token)
	listKubernetes(client)
	listDroplets(client)
	listVolumes(client)
	listLoadBalancers(client)
	listDomains(client)
	listSSHKeys(client)
}

func PrepareVolumesCleanUp(token string) []godo.Volume {
	var volumesForCleanUp []godo.Volume
	client := godo.NewFromToken(token)
	volumes, _, _ := client.Storage.ListVolumes(context.TODO(), &godo.ListVolumeParams{})
	for _, volume := range volumes {
		if len(volume.DropletIDs) == 0 {
			volumesForCleanUp = append(volumesForCleanUp, volume)
		}
	}
	fmt.Println("Volumes marked for clean up:")
	for _, v := range volumesForCleanUp {
		fmt.Println(v.Name)
	}

	return volumesForCleanUp
}

func DoVolumesCleanUp(token string, volumesForCleanUp []godo.Volume) {
	var err error
	client := godo.NewFromToken(token)
	for _, v := range volumesForCleanUp {
		_, err = client.Storage.DeleteVolume(context.TODO(), v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
