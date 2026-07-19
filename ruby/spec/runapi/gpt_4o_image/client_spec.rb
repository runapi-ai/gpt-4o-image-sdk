# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::Gpt4oImage::Client do
  let(:client) { described_class.new(api_key: "test-key") }

  it "exposes text_to_image resource" do
    expect(client.text_to_image).to be_a(RunApi::Gpt4oImage::Resources::TextToImage)
  end
end
