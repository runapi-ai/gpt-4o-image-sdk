package gpt4oimage

type TaskStatus string

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

type AsyncTaskResponse struct {
	ID       string     `json:"id"`
	Status   TaskStatus `json:"status"`
	Progress string     `json:"progress,omitempty"`
	Error    string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

type Image struct {
	URL string `json:"url"`
}

type TextToImageResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}
