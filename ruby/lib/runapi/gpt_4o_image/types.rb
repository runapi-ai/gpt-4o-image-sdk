# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    # GPT-4o Image response models.
    # Supports pure generation, editing with source images, and inpainting with a mask.
    module Types
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
