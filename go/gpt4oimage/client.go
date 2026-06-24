// Package gpt4oimage provides the GPT-4o Image generation API client.
//
//	client, err := gpt4oimage.NewClient(option.WithAPIKey("sk-your-api-key"))
//	result, err := client.TextToImage.Run(ctx, gpt4oimage.TextToImageParams{
//	    Model: "gpt-4o-image", Prompt: "A watercolor painting of a sunset", AspectRatio: "3:2",
//	})
package gpt4oimage

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const textToImagePath = "/api/v1/gpt_4o_image/text_to_image"

// Client provides GPT-4o powered image generation with optional editing via source images and masks.
type Client struct {
	base.Base
	TextToImage *TextToImage
}

// NewClient creates a GPT-4o Image client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a GPT-4o Image client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{Base: base.New(httpClient), TextToImage: &TextToImage{http: httpClient}}
}

// TextToImage generates images from a text prompt, optionally guided by source images and a mask.
// For pure generation, provide a Prompt. For editing, provide SourceImageURLs (and optionally MaskURL)
// to modify specific regions of the source. At least one of Prompt or SourceImageURLs must be set.
type TextToImage struct{ http core.HTTPClient }

// Create submits a GPT-4o Image text-to-image task and returns immediately with a task id.
func (r *TextToImage) Create(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := core.ValidateParams(contractSchema["text-to-image"], body); err != nil {
		return nil, err
	}
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToImagePath, body, requestOptions)
}

// Get fetches the current status of a GPT-4o Image text-to-image task by id.
func (r *TextToImage) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TextToImageResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TextToImageResponse](ctx, r.http, core.ResourcePath(textToImagePath, id), requestOptions)
}

// Run submits a GPT-4o Image text-to-image task and polls until it completes.
func (r *TextToImage) Run(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*TextToImageResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*TextToImageResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
