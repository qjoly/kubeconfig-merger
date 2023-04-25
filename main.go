package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type kubeconfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Clusters   []struct {
		Name    string `yaml:"name"`
		Cluster struct {
			Server                   string `yaml:"server"`
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
		} `yaml:"cluster"`
	} `yaml:"clusters"`
	Contexts []struct {
		Name    string `yaml:"name"`
		Context struct {
			Cluster   string `yaml:"cluster"`
			Namespace string `yaml:"namespace"`
			User      string `yaml:"user"`
		} `yaml:"context"`
	} `yaml:"contexts"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Username              string `yaml:"username"`
			Password              string `yaml:"password"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: concat-kubeconfig <first-file> <second-file>")
		os.Exit(1)
	}

	// Load first kubeconfig file
	firstKubeConfig, err := loadKubeConfig(os.Args[1])
	if err != nil {
		fmt.Println("Error loading first kubeconfig file:", err)
		os.Exit(1)
	}

	// Load second kubeconfig file
	secondKubeConfig, err := loadKubeConfig(os.Args[2])
	if err != nil {
		fmt.Println("Error loading second kubeconfig file:", err)
		os.Exit(1)
	}

	for _, firstCluster := range firstKubeConfig.Clusters {
		for _, secondCluster := range secondKubeConfig.Clusters {
			if firstCluster.Name == secondCluster.Name {
				fmt.Printf("Error: Both kubeconfigs have a cluster named '%s'\n", firstCluster.Name)
				os.Exit(1)
			}
		}
	}
//	if hasDefaultCluster(firstKubeConfig) && hasDefaultCluster(secondKubeConfig) {
//		fmt.Println("Both kubeconfig files contain a cluster named 'default'")
//		os.Exit(1)
//	}

	// Concatenate clusters
	for _, cluster := range secondKubeConfig.Clusters {
		firstKubeConfig.Clusters = append(firstKubeConfig.Clusters, cluster)
	}

	// Concatenate contexts
	for _, context := range secondKubeConfig.Contexts {
		firstKubeConfig.Contexts = append(firstKubeConfig.Contexts, context)
	}

	// Concatenate users
	for _, user := range secondKubeConfig.Users {
		firstKubeConfig.Users = append(firstKubeConfig.Users, user)
	}

	// Write concatenated kubeconfig file
	err = writeKubeConfig("combined-kubeconfig.yaml", firstKubeConfig)
	if err != nil {
		fmt.Println("Error writing combined kubeconfig file:", err)
		os.Exit(1)
	}

	fmt.Println("Kubeconfig files combined successfully!")
}

func loadKubeConfig(filename string) (*kubeconfig, error) {
	// Read kubeconfig file
	data, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}

	// Parse kubeconfig file
	kc := kubeconfig{}
	err = yaml.Unmarshal(data, &kc)
	if err != nil {
		return nil, err
	}

	return &kc, nil
}

func writeKubeConfig(filename string, kc *kubeconfig) error {
	// Convert kubeconfig to YAML
	data, err := yaml.Marshal(kc)
	if err != nil {
		return err
	}

	// Write YAML to file
	err = ioutil.WriteFile(filepath.Clean(filename), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

//func hasDefaultCluster(kc *kubeconfig) bool {
//	for _, cluster := range kc.Clusters {
//		if cluster.Name == "default" {
//			return true
//		}
//	}
//	return false
//}

