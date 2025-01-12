package install

const (
	AgentContainerName        = "traffic-agent"
	AgentAnnotationVolumeName = "traffic-annotations"
	AgentInjectorName         = "agent-injector"
	DomainPrefix              = "telepresence.getambassador.io/"
	InjectAnnotation          = DomainPrefix + "inject-" + AgentContainerName
	ServicePortAnnotation     = DomainPrefix + "inject-service-port"
	ServiceNameAnnotation     = DomainPrefix + "inject-service-name"
	ManualInjectAnnotation    = DomainPrefix + "manually-injected"
	ManagerAppName            = "traffic-manager"
	MutatorWebhookTLSName     = "mutator-webhook-tls"
	TelAppMountPoint          = "/tel_app_mounts"
)
