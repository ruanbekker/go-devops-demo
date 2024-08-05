#!/usr/bin/env bash
CLUSTER_NAME=${KIND_CLUSTER_NAME:-go-devops-demo}
CONTEXT_NAME="kind-${CLUSTER_NAME}"
KUBECONFIG_FILENAME="/tmp/${CLUSTER_NAME}.config"

export KUBECONFIG=${KUBECONFIG_FILENAME}

function is_command_available(){
    command -v ${1} > /dev/null 2>&1
    echo ${?}
}

function print_kind_installation() {
    echo "[error] kind is not installed, to install:"
    echo curl -Lo /tmp/kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
    echo sudo install -o root -g root -m 0755 /tmp/kind /usr/local/bin/kind
}

function create_if_not_exists(){
    if [[ "$(kind get clusters | wc -l)" -gt 0 ]]
    then 
      echo "cluster exists with name ${CLUSTER_NAME}"
    else 
      kind create cluster --name ${CLUSTER_NAME} --config kind/kind-config.yaml
    fi
}

function is_cluster_ready(){
    kubectl cluster-info --context ${CONTEXT_NAME} > /dev/null 2>&1
    return ${?} 
}

function is_node_ready(){
    local status
    status=$(kubectl get nodes --context ${CONTEXT_NAME} | awk '/'$CLUSTER_NAME'-control-plane/ {print $2}')
    if [[ "${status}" == "Ready" ]]; then return 0; else return 1; fi
}

function install_apps(){
    # nginx ingress
    helm --kube-context ${CONTEXT_NAME} repo add nginx https://kubernetes.github.io/ingress-nginx
    helm --kube-context ${CONTEXT_NAME} upgrade --install nginx-public nginx/ingress-nginx \
      --version 4.7.3 --namespace kube-system \
      --set controller.admissionWebhooks.enabled=false --set controller.hostPort.enabled=true \
      --set controller.ingressClass=nginx --set controller.service.type=NodePort
    kubectl --context ${CONTEXT_NAME} wait --namespace kube-system \
      --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=300s
    # prometheus
    helm --kube-context ${CONTEXT_NAME} repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm --kube-context ${CONTEXT_NAME} upgrade --install kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 61.7.0 --namespace kube-system
    # go devops demo
    kubectl --context ${CONTEXT_NAME} apply -f kubernetes/manifests/ -n default
}

if [[ "$(is_command_available kind)" != 0 ]]
  then
    print_kind_installation
    exit 1
  else
    create_if_not_exists
    until is_cluster_ready
    do
      echo "waiting for cluster to become ready.."
      sleep 5
    done
    echo "cluster is ready"
    until is_node_ready
    do
      echo "waiting for node to become ready.."
      sleep 5
    done
    echo "node is ready"
    install_apps
fi
