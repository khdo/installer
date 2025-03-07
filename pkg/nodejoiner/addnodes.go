package nodejoiner

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/agent/joiner"
	"github.com/openshift/installer/pkg/asset/agent/workflow"
	"github.com/openshift/installer/pkg/asset/store"
)

// NewAddNodesCommand creates a new command for add nodes.
func NewAddNodesCommand(directory string, kubeConfig string) error {
	// Store the current parameters into the assets folder, so
	// that they could be retrieved later by the assets
	params := joiner.Params{
		Kubeconfig: kubeConfig,
	}
	err := params.Save(directory)
	if err != nil {
		return err
	}

	fetcher := store.NewAssetsFetcher(directory)
	return fetcher.FetchAndPersist([]asset.WritableAsset{
		&workflow.AgentWorkflowAddNodes{},
		// To be completed
	})
}
