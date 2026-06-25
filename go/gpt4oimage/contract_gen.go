package gpt4oimage

var contractSchema = map[string]any{"text-to-image": map[string]any{"models": []any{"gpt-4o-image"}, "fields_by_model": map[string]any{"gpt-4o-image": map[string]any{"aspect_ratio": map[string]any{"enum": []any{"1:1", "3:2", "2:3"}, "required": true}, "output_count": map[string]any{"enum": []any{1, 2, 4}, "type": "integer"}}}}}
