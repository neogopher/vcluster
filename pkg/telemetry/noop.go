package telemetry

import (
	"context"

	managementv1 "github.com/loft-sh/api/v3/pkg/apis/management/v1"
	"github.com/loft-sh/vcluster/pkg/config"
	"k8s.io/client-go/kubernetes"
)

type noopCollector struct{}

func (n *noopCollector) RecordStart(_ context.Context, _ *config.VirtualClusterConfig) {}

func (n *noopCollector) RecordError(_ context.Context, _ *config.VirtualClusterConfig, _ ErrorSeverityType, _ error) {
}

func (n *noopCollector) Flush() {}

func (n *noopCollector) SetVirtualClient(_ *kubernetes.Clientset) {}

func (n *noopCollector) RecordCLI(_ *managementv1.Self, _ error) {}
