package envoy

import (
    "context"

    "github.com/cortezaproject/corteza/server/pkg/envoyx"
)

func (e YamlEncoder) encode(ctx context.Context, base *yaml.Node, p envoyx.EncodeParams, rt string, nodes envoyx.NodeSet, tt envoyx.Traverser) (out *yaml.Node, err error) {
	return
}
