# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    # GPT-4o Image type constants and response models.
    # Supports pure generation, editing with source images, and inpainting with a mask.
    module Types
      MODELS = %w[gpt-4o-image].freeze
      # Required output aspect ratio.
      ASPECT_RATIOS = %w[1:1 3:2 2:3].freeze
      # Batch sizes: 1, 2, or 4 images per request.
      OUTPUT_COUNTS = [1, 2, 4].freeze

      # A generated image with its CDN URL.
      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      # Response for GPT-4o Image tasks. Progress contains an intermediate
      # status string while the task is running.
      class TextToImageResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :images, [-> { Image }]
        optional :progress, String
        optional :error, String
      end

      CompletedTextToImageResponse = TextToImageResponse
    end
  end
end
