package node_gateway

import (
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/pkg/errors"
	"math/rand"
	"sync"
)

type NodeRegistry interface {
	GetRandomNode(asset v1.Asset, version string) (*v1.Node, error)
	GetNodes(asset v1.Asset, version string) ([]*v1.Node, error)
	Add(node *v1.Node) error
}

type MemoryRegistry struct {
	mu    sync.RWMutex
	nodes map[v1.Asset]map[string][]*v1.Node
}

func (m *MemoryRegistry) GetRandomNode(asset v1.Asset, version string) (*v1.Node, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nodes[asset][version][rand.Intn(len(m.nodes[asset]))], nil
}

func (m *MemoryRegistry) GetNodes(asset v1.Asset, version string) ([]*v1.Node, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nodes[asset][version], nil
}

func (m *MemoryRegistry) Add(node *v1.Node) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if existing, ok := m.nodes[node.Asset][node.Version]; !ok {
		m.nodes[node.Asset][node.Version] = []*v1.Node{node}
		return nil
	} else {
		for _, n := range existing {
			if n.Host == node.Host && n.Port == node.Port && n.Version == node.Version {
				return errors.New("node is already registered")
			}
		}
		m.nodes[node.Asset][node.Version] = append(existing, node)
	}
	return nil
}
