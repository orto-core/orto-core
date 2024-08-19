#!/bin/bash

# Reusable function to apply Kubernetes manifests
kubectlApply() {
    local file="$1"
    local success_msg="$2"
    local failure_msg="$3"

    kubectl apply -f "$file"
    if [ $? -eq 0 ]; then
        echo "$success_msg"
    else
        echo "$failure_msg"
    fi
}

# Install Prometheus Operator
cmd="kubectl get deployment prometheus-operator"
$cmd &> /dev/null
STATUS=$?

# Check status of the command
if [ $STATUS -eq 0 ]; then
    echo "Prometheus operator deployment already exists!"
else
    echo "No deployment found for prometheus-operator!"
    # Installing the Prometheus Operator
    kubectl create -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/master/bundle.yaml
    # Re-trying the command
    $cmd &> /dev/null
    STATUS=$?
    if [ $STATUS -ne 0 ]; then
        echo "Failed to create Prometheus operator deployment."
        exit $STATUS
    fi
fi

# Configure Prometheus RBAC Permissions
mkdir -p operator_k8
file="./operator_k8/prom_rbac.yaml"

if [ -f "$file" ]; then
    echo "RBAC manifest file already exists."
else
    cat <<EOF > "$file"
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
  - apiGroups: [""]
    resources:
      - nodes
      - nodes/metrics
      - services
      - endpoints
      - pods
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources:
      - configmaps
    verbs: ["get"]
  - apiGroups:
      - networking.k8s.io
    resources:
      - ingresses
    verbs: ["get", "list", "watch"]
  - nonResourceURLs: ["/metrics"]
    verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
  - kind: ServiceAccount
    name: prometheus
    namespace: default
EOF
    echo "RBAC file created and populated successfully!"
fi

kubectlApply "$file" "Role created successfully" "Failed to create role"

# Deploy Prometheus
mkdir -p prometheus

file="./prometheus/prometheus.yaml"

if [ -f "$file" ]; then
    echo "Prometheus manifest file already exists."
else
    cat <<EOF > "$file"
apiVersion: monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: prometheus
  labels:
    app: prometheus
spec:
  image: quay.io/prometheus/prometheus:v2.22.1
  nodeSelector:
    kubernetes.io/os: linux
  replicas: 2
  resources:
    requests:
      memory: 400Mi
  securityContext:
    fsGroup: 2000
    runAsNonRoot: true
    runAsUser: 1000
  serviceAccountName: prometheus
  version: v2.22.1
  serviceMonitorSelector: {}
EOF
    echo "Prometheus file created and populated successfully!"
fi

kubectlApply "$file" "Prometheus created successfully" "Failed to create Prometheus"

file="./prometheus/prom_svc.yaml"

if [ -f "$file" ]; then
    echo "Prometheus service manifest file already exists."
else
    cat <<EOF > "$file"
apiVersion: v1
kind: Service
metadata:
  name: prometheus
  labels:
    app: prometheus
spec:
  ports:
    - name: web
      port: 9090
      targetPort: web
  selector:
    app.kubernetes.io/name: prometheus
  sessionAffinity: ClientIP
EOF
    echo "Prometheus service file created and populated successfully!"
fi

kubectlApply "$file" "Prometheus service created successfully" "Failed to create Prometheus service"

file="./prometheus/prometheus_servicemonitor.yaml"

if [ -f "$file" ]; then
    echo "Prometheus service monitor manifest file already exists."
else
    cat <<EOF > "$file"
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: prometheus-self
  labels:
    app: prometheus
spec:
  endpoints:
    - interval: 30s
      port: web
  selector:
    matchLabels:
      app: prometheus
EOF
    echo "Prometheus service monitor file created and populated successfully!"
fi

kubectlApply "$file" "Prometheus service monitor created successfully" "Failed to create Prometheus service monitor"


# Create a grafana Deployment and service
mkdir -p grafana

file="./grafana/grafana.yaml"

if [ -f "$file" ]; then
    echo "Prometheus service monitor manifest file already exists."
else
    cat <<EOF > "$file"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grafana
  labels:
    app: grafana
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      labels:
        app: grafana
    spec:
      containers:
        - name: grafana
          image: grafana/grafana
          ports:
            - containerPort: 3000
          volumeMounts:
            - name: grafana-storage
              mountPath: /var/lib/grafana
      volumes:
        - name: grafana-storage
          emptyDir: {}

---
apiVersion: v1
kind: Service
metadata:
  name: grafana
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 3000
  selector:
    app: grafana
EOF
echo "Grafana manifest file created and populated successfully!"
fi

kubectlApply "$file" "Grafana deployment created successfully" "Failed to create grafana deployment"
