# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    module Types
      MODELS = %w[gpt-4o-image].freeze
      ASPECT_RATIOS = %w[1:1 3:2 2:3].freeze
      OUTPUT_COUNTS = [1, 2, 4].freeze

      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

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
