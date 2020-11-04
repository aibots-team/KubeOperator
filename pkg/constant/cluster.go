package constant

const (
	ClusterRunning      = "Running"
	ClusterInitializing = "Initializing"
	ClusterNotConnected = "NotConnected"
	ClusterFailed       = "Failed"
	ClusterTerminating  = "Terminating"
	ClusterTerminated   = "Terminated"
	ClusterWaiting      = "Waiting"
	ClusterUpgrading    = "Upgrading"
	//日志类型

	ClusterLogTypeAddNode    = "ADD_NODE"
	ClusterLogTypeDeleteNode = "DELETE_NODE"
	ClusterLogTypeBackup     = "CLUSTER_BACKUP"
	ClusterLogTypeRestore    = "CLUSTER_RESTORE"
	ClusterLogTypeUpgrade    = "CLUSTER_UPGRADE"

	ClusterLogStatusSuccess = "SUCCESS"
	ClusterLogStatusFailed  = "FAILED"
	ClusterLogStatusWaiting = "WAITING"
	ClusterLogStatusRunning = "RUNNING"

	// 表示创建资源
	ClusterCreating = "Creating"

	ClusterSourceLocal    = "local"
	ClusterSourceExternal = "external"

	ConditionTrue    = "True"
	ConditionFalse   = "False"
	ConditionUnknown = "Unknown"

	NodeRoleNameMaster = "master"
	NodeRoleNameWorker = "worker"

	ClusterProviderBareMetal = "bareMetal"
	ClusterProviderPlan      = "plan"

	DefaultNamespace     = "kube-operator"
	F5Namespace          = "kube-system"
	DefaultApiServerPort = 8443

	DefaultIngress            = "apps.ko.com"
	DefaultPrometheusIngress  = "prometheus." + DefaultIngress
	DefaultLoggingIngress     = "logging." + DefaultIngress
	DefaultChartmuseumIngress = "chartmuseum." + DefaultIngress
	DefaultRegistryIngress    = "registry." + DefaultIngress
	DefaultDashboardIngress   = "dashboard." + DefaultIngress
	DefaultKubeappsIngress    = "kubeapps." + DefaultIngress

	ChartmuseumChartName    = "nexus/chartmuseum"
	DockerRegistryChartName = "nexus/docker-registry"
	PrometheusChartName     = "nexus/prometheus"
	LoggingChartName        = "nexus/logging"
	DashboardChartName      = "nexus/kubernetes-dashboard"
	KubeappsChartName       = "nexus/kubeapps"

	DefaultRegistryServiceName    = "registry-docker-registry"
	DefaultChartmuseumServiceName = "chartmuseum-chartmuseum"
	DefaultDashboardServiceName   = "dashboard-kubernetes-dashboard"
	DefaultLoggingServiceName     = "elasticsearch-master"
	DefaultPrometheusServiceName  = "prometheus-server"
	DefaultKubeappsServiceName    = "kubeapps"

	DefaultRegistryIngressName    = "docker-registry-ingress"
	DefaultChartmuseumIngressName = "chartmuseum-ingress"
	DefaultDashboardIngressName   = "dashboard-ingress"
	DefaultLoggingIngressName     = "logging-ingress"
	DefaultPrometheusIngressName  = "prometheus-ingress"
	DefaultKubeappsIngressName    = "kubeapps-ingress"

	DefaultRegistryDeploymentName    = "registry-docker-registry"
	DefaultChartmuseumDeploymentName = "chartmuseum-chartmuseum"
	DefaultDashboardDeploymentName   = "dashboard-kubernetes-dashboard"
	DefaultKubeappsDeploymentName    = "kubeapps"
	DefaultLoggingStateSetsfulName   = "elasticsearch-master"
	DefaultPrometheusDeploymentName  = "prometheus-server"
)
