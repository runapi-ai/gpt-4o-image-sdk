package ai.runapi.gpt4oimage.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for text to image operations. */
public final class TextToImageParams {
  private final String model;
  private final String prompt;
  private final String aspectRatio;
  private final List<String> sourceImageUrls;
  private final String maskUrl;
  private final Integer outputCount;
  private final String callbackUrl;
  private final Boolean enablePromptExpansion;

  private TextToImageParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.aspectRatio = Gpt4oimageParamUtils.requireNonBlank(builder.aspectRatio, "aspectRatio");
    this.sourceImageUrls = Gpt4oimageParamUtils.strings(builder.sourceImageUrls);
    this.maskUrl = builder.maskUrl;
    this.outputCount = builder.outputCount;
    this.callbackUrl = builder.callbackUrl;
    this.enablePromptExpansion = builder.enablePromptExpansion;
  }

  /** Creates a new TextToImageParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gpt-4o-image/text-to-image";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", Gpt4oimageParamUtils.wireValue(model));
    raw.put("prompt", Gpt4oimageParamUtils.wireValue(prompt));
    raw.put("aspect_ratio", Gpt4oimageParamUtils.wireValue(aspectRatio));
    raw.put("source_image_urls", Gpt4oimageParamUtils.wireValue(sourceImageUrls));
    raw.put("mask_url", Gpt4oimageParamUtils.wireValue(maskUrl));
    raw.put("output_count", Gpt4oimageParamUtils.wireValue(outputCount));
    raw.put("callback_url", Gpt4oimageParamUtils.wireValue(callbackUrl));
    raw.put("enable_prompt_expansion", Gpt4oimageParamUtils.wireValue(enablePromptExpansion));
    return Gpt4oimageParamUtils.compact(raw);
  }



  /** Builder for {@link TextToImageParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private String aspectRatio;
    private List<String> sourceImageUrls;
    private String maskUrl;
    private Integer outputCount;
    private String callbackUrl;
    private Boolean enablePromptExpansion;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(TextToImageModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = Gpt4oimageParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = Gpt4oimageParamUtils.requireNonBlank(value, "prompt");
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = Gpt4oimageParamUtils.requireNonBlank(value, "aspectRatio");
      return this;
    }

    /** Sets the source image URLs. */
    public Builder sourceImageUrls(List<String> value) {
      this.sourceImageUrls = value;
      return this;
    }

    /** Sets the mask URL. */
    public Builder maskUrl(String value) {
      this.maskUrl = Gpt4oimageParamUtils.requireNonBlank(value, "maskUrl");
      return this;
    }

    /** Sets the number of generated outputs. */
    public Builder outputCount(int value) {
      this.outputCount = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = Gpt4oimageParamUtils.requireNonBlank(value, "callbackUrl");
      return this;
    }

    /** Sets the prompt expansion toggle. */
    public Builder enablePromptExpansion(boolean value) {
      this.enablePromptExpansion = value;
      return this;
    }

    /** Builds immutable text to image parameters. */
    public TextToImageParams build() {
      return new TextToImageParams(this);
    }
  }
}
