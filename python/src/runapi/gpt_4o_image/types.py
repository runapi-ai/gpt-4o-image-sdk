"""GPT-4o Image model lists, enums, and response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required

MODELS = ["gpt-4o-image"]
ASPECT_RATIOS = ["1:1", "3:2", "2:3"]
OUTPUT_COUNTS = [1, 2, 4]


class Image(BaseModel):
    url = optional(str)


class TextToImageResponse(TaskResponse):
    """Task status/result for GPT-4o Image text-to-image."""
    id = required(str)
    status = optional(str, enum=lambda: TaskResponse.Status.ALL)
    images = optional([lambda: Image])
    progress = optional(str)
    error = optional(str)


CompletedTextToImageResponse = TextToImageResponse
