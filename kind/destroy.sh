#!/usr/bin/env bash
CLUSTER_NAME=${KIND_CLUSTER_NAME:-go-devops-demo}

function is_command_available(){
    command -v ${1} > /dev/null 2>&1
    echo ${?}
}

function print_kind_installation() {
    echo "[error] kind is not installed, to install:"
    echo curl -Lo /tmp/kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
    echo sudo install -o root -g root -m 0755 /tmp/kind /usr/local/bin/kind
}

function destroy_cluster(){
    if [[ "$(kind get clusters | wc -l)" -eq 0 ]]
    then 
      echo "no clusters exists"
    else 
      kind delete cluster --name ${CLUSTER_NAME}
    fi
}

if [[ "$(is_command_available kind)" != 0 ]]
  then
    print_kind_installation
    exit 1
  else
    destroy_cluster
fi
