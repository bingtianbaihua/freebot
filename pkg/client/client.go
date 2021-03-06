package client

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type ClientInterface interface {
	DoOperation(ctx context.Context, op interface{}) error
}

type githubClient struct {
	client *github.Client
}

func NewGithubClient(client *github.Client) ClientInterface {
	return &githubClient{
		client: client,
	}
}

func (cli *githubClient) DoOperation(ctx context.Context, op interface{}) (err error) {
	switch v := op.(type) {
	case *ReplaceLabelOperation:
		err = cli.doReplaceLabelOperation(ctx, v)
	case *RequestReviewsOperation:
		err = cli.doRequestReviewsOperation(ctx, v)
	case *RequestReviewsCancelOperation:
		err = cli.doRequestReviewsCancelOperation(ctx, v)
	case *AddAssignOperation:
		err = cli.doAddAssignOperation(ctx, v)
	case *RemoveAssignOperation:
		err = cli.doRemoveAssignOperation(ctx, v)
	case *MergeOperation:
		err = cli.doMergeOperation(ctx, v)
	default:
		err = fmt.Errorf("no support operation")
	}
	return
}
