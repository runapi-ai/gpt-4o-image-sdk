package ai.runapi.gpt4oimage.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to image operations. */
public final class TextToImageModel extends Gpt4oimageValue {
  /** gpt-4o-image model slug. */
  public static final TextToImageModel GPT_4O_IMAGE = new TextToImageModel("gpt-4o-image");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToImageModel(String value) {
    super(value);
  }
}
