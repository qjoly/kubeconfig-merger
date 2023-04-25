# KubeConfig-Merger

## Introduction

kubeconfig-merger is a command-line tool that concatenates two Kubernetes configuration files (kubeconfig files) into a single file. The resulting file can be used as a single kubeconfig file to manage multiple Kubernetes clusters.

## Prerequisites

To use kubeconfig-merger, you must have Go version 1.16 or later installed on your system. You can download Go from the official website: https://golang.org/dl/

## Installation

### Downloading the Pre-Compiled Release

You can download a pre-compiled release of kubeconfig-merger from the [release page](https://github.com/QJoly/kubeconfig-merger/releases). Select the latest release and download the binary for your operating system.

Once you have downloaded the binary, you can use it directly without needing to compile the code.

### Compiling from Source

1. Clone this repository:

   ```bash
   git clone https://github.com/QJoly/kubeconfig-merger
   ```

2. Change to the directory where the repository was cloned:

   ```bash
   cd kubeconfig-merger
   ```

3. Compile the Go code into an executable binary:

   ```bash
   go build -o kubeconfig-merger main.go
   ```

   This will create an executable binary named `kubeconfig-merger` in the current directory.

## Usage

To use kubeconfig-merger, run the executable binary with the following command:

```
./kubeconfig-merger <first-file> <second-file>
```

Replace `<first-file>` and `<second-file>` with the names of the two kubeconfig files you want to concatenate.

For example, if you have two kubeconfig files named `kubeconfig-1.yaml` and `kubeconfig-2.yaml`, you can concatenate them into a single file named `combined-kubeconfig.yaml` with the following command:

```
./kubeconfig-merger kubeconfig-1.yaml kubeconfig-2.yaml
```

The concatenated kubeconfig file will be saved as `combined-kubeconfig.yaml` in the current directory.

[![asciicast](https://asciinema.org/a/5Myq4ZzzvYo9uSuhx7LgMMPEH.svg)](https://asciinema.org/a/5Myq4ZzzvYo9uSuhx7LgMMPEH)

## Purpose

The purpose of kubeconfig-merger is to simplify management of multiple Kubernetes clusters by allowing you to use a single kubeconfig file to access them all. By concatenating multiple kubeconfig files into a single file, you can avoid the need to switch between different kubeconfig files when working with different clusters. This can make it easier to manage and automate Kubernetes operations.
