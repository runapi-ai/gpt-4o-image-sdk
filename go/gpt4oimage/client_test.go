package gpt4oimage

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return s.response, nil
}

func TestTextToImageCreate(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_gen_123","status":"processing"}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Create(context.Background(), TextToImageParams{
		Model:           "gpt-4o-image",
		Prompt:          "a still life",
		AspectRatio:     "1:1",
		SourceImageURLs: []string{"https://cdn.runapi.ai/public/samples/input.png"},
		OutputCount:     2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_4o_image/text_to_image" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-4o-image" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["aspect_ratio"] != "1:1" {
		t.Fatalf("unexpected aspect_ratio: %v", body["aspect_ratio"])
	}
	if _, ok := body["size"]; ok {
		t.Fatal("expected removed size field to be absent")
	}
	if _, ok := body["files_url"]; ok {
		t.Fatal("expected removed files_url field to be absent")
	}
	if _, ok := body["fallback_model"]; ok {
		t.Fatal("expected routing fallback field to be absent")
	}
	if _, ok := body["enable_fallback"]; ok {
		t.Fatal("expected routing fallback field to be absent")
	}
	if _, ok := body["upload_cn"]; ok {
		t.Fatal("expected routing upload field to be absent")
	}
	sourceImageURLs, ok := body["source_image_urls"].([]any)
	if !ok || len(sourceImageURLs) != 1 || sourceImageURLs[0] != "https://cdn.runapi.ai/public/samples/input.png" {
		t.Fatalf("unexpected source_image_urls: %v", body["source_image_urls"])
	}
	if body["output_count"] != float64(2) {
		t.Fatalf("unexpected output_count: %v", body["output_count"])
	}
	if resp.ID != "task_gen_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestTextToImageGet(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_gen_456","status":"completed","progress":"1.00","images":[{"url":"https://file.runapi.ai/result.png"}]}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Get(context.Background(), "task_gen_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_4o_image/text_to_image/task_gen_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_gen_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
}
