package gpt4oimage

// TaskStatus is the async task lifecycle state (e.g. "processing", "completed", "failed").
type TaskStatus string

// TextToImageParams configures GPT-4o image generation.
// Prompt is required for pure generation; when SourceImageURLs is provided,
// Prompt becomes optional (the model can generate variants without one).
// MaskURL is only meaningful when SourceImageURLs is also set, indicating which
// regions of the source to regenerate (inpainting).
// AspectRatio is required: "1:1", "3:2", or "2:3".
type TextToImageParams struct {
	Model                 string   `json:"model" help:"required; model slug"`
	Prompt                string   `json:"prompt,omitempty" help:"optional; prompt text, required when source_image_urls is blank"`
	AspectRatio           string   `json:"aspect_ratio" help:"required; output aspect ratio"`
	SourceImageURLs       []string `json:"source_image_urls,omitempty" help:"optional; up to 5 source image URLs for edits or variants"`
	MaskURL               string   `json:"mask_url,omitempty" help:"optional; mask image URL"`
	OutputCount           int      `json:"output_count,omitempty" help:"optional; number of generated images"`
	CallbackURL           string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
	EnablePromptExpansion *bool    `json:"enable_prompt_expansion,omitempty" help:"optional; prompt expansion toggle"`
}

// AsyncTaskResponse carries the task ID, lifecycle status, and error for GPT-4o Image async operations.
// Progress may contain an intermediate status string while the task is running.
type AsyncTaskResponse struct {
	ID       string     `json:"id"`
	Status   TaskStatus `json:"status"`
	Progress string     `json:"progress,omitempty"`
	Error    string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

// Image holds a URL to a generated image.
type Image struct {
	URL string `json:"url"`
}

// TextToImageResponse is the result of a text-to-image task, containing one or more generated images.
type TextToImageResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}
