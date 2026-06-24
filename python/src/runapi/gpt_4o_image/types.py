"""GPT-4o Image response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required


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
