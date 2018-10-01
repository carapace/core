package node_gateway

import (
	"bytes"
	"context"
	"fmt"
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"sync"
)

// compile time assertion to verify we match interface as defined in proto definitions
var _ v1.NodeGatewayServiceServer = &Gateway{}

type Gateway struct {
	nodes  NodeRegistry
	client http.Client
}

func (g *Gateway) Post(ctx context.Context, post *v1.NodePost) (*v1.NodeResponse, error) {
	var targets []*v1.Node
	var err error

	if post.Repeatable {
		targets, err = g.nodes.GetNodes(post.Asset, post.Version)
	} else {
		var node *v1.Node
		node, err = g.nodes.GetRandomNode(post.Asset, post.Version)
		targets = []*v1.Node{node}
	}

	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, errors.New("no nodes are registered with appropriate credentials")
	}

	resps, errs := g.postRequests(targets, post)
	if len(resps) == 0 {
		return nil, errors.New(fmt.Sprintf("unable to make a single postRequest: %s", errs))
	}

	res := &v1.NodeResponse{}
	for _, r := range resps {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		res.Results = append(res.Results, bs)
	}
	return res, nil
}

func (g *Gateway) Get(context.Context, *v1.NodeGet) (*v1.NodeResponse, error) {
	return nil, nil
}

func (g *Gateway) Register(ctx context.Context, node *v1.Node) (*v1.Response, error) {
	return nil, nil
}

func (g *Gateway) postRequests(targets []*v1.Node, post *v1.NodePost) ([]*http.Response, []error) {
	resps := make(chan *http.Response, len(targets))
	errs := make(chan error, len(targets))

	wg := sync.WaitGroup{}
	for _, target := range targets {
		wg.Add(1)
		go func(target *v1.Node) {
			resp, err := g.postRequest(target, post)
			resps <- resp
			errs <- err
			wg.Done()
		}(target)
	}
	wg.Wait()

	r := []*http.Response{}

	for {
		v, ok := <-resps
		if !ok {
			break
		}
		r = append(r, v)
	}

	e := []error{}
	for {
		v, ok := <-errs
		if !ok {
			break
		}
		e = append(e, v)
	}
	return r, e
}

func (g *Gateway) postRequest(target *v1.Node, post *v1.NodePost) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s:%s/%s", target.Host, target.Port, post.Endpoint),
		bytes.NewReader(post.Body),
	)
	if err != nil {
		return nil, err
	}
	return g.client.Do(req)
}
