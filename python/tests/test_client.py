import pytest

from runapi.core import config
from runapi.core.errors import AuthenticationError, ValidationError
from runapi.gpt_4o_image import Gpt4oImageClient
from runapi.gpt_4o_image.resources.text_to_image import TextToImage
from runapi.gpt_4o_image.types import CompletedTextToImageResponse, TextToImageResponse


class FakeHttp:
    """Records (method, path, body) and replays preset responses by call order."""

    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- authentication -------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(
        Gpt4oImageClient(api_key="param-key", http_client=FakeHttp()), Gpt4oImageClient
    )


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(Gpt4oImageClient(http_client=FakeHttp()), Gpt4oImageClient)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(Gpt4oImageClient(http_client=FakeHttp()), Gpt4oImageClient)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        Gpt4oImageClient()


# --- transport injection / accessors --------------------------------------


def test_uses_injected_http_client():
    fake = FakeHttp()
    client = Gpt4oImageClient(api_key="k", http_client=fake)
    assert client.text_to_image._http is fake


def test_exposes_resource_accessors():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    assert isinstance(client.text_to_image, TextToImage)


# --- request shapes -------------------------------------------------------


def test_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = Gpt4oImageClient(api_key="k", http_client=fake)
    result = client.text_to_image.create(
        model="gpt-4o-image",
        prompt="hello",
        aspect_ratio="1:1",
        mask_url=None,
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gpt_4o_image/text_to_image",
            {"model": "gpt-4o-image", "prompt": "hello", "aspect_ratio": "1:1"},
        ),
    ]
    assert isinstance(result, TextToImageResponse)
    assert result.id == "t1"


def test_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = Gpt4oImageClient(api_key="k", http_client=fake)
    client.text_to_image.get("t1")
    assert fake.calls == [("get", "/api/v1/gpt_4o_image/text_to_image/t1", None)]


def test_run_polls_and_narrows_completed_type():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "images": [{"url": "https://x/y.png"}]},
    )
    client = Gpt4oImageClient(api_key="k", http_client=fake)
    result = client.text_to_image.run(
        model="gpt-4o-image", prompt="hi", aspect_ratio="1:1"
    )

    assert isinstance(result, CompletedTextToImageResponse)
    assert result.images[0].url == "https://x/y.png"
    assert [call[0] for call in fake.calls] == ["post", "get"]


# --- validation -----------------------------------------------------------


def test_create_requires_model():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="model is required"):
        client.text_to_image.create(aspect_ratio="1:1")


def test_create_requires_aspect_ratio():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="aspect_ratio is required"):
        client.text_to_image.create(model="gpt-4o-image")


def test_create_rejects_unknown_model():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match=r"Invalid model: not-a-model\. Must be: gpt-4o-image"):
        client.text_to_image.create(model="not-a-model", aspect_ratio="1:1")


def test_create_rejects_invalid_aspect_ratio():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="Invalid aspect_ratio"):
        client.text_to_image.create(model="gpt-4o-image", aspect_ratio="99:1")


def test_create_rejects_invalid_output_count():
    client = Gpt4oImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="Invalid output_count"):
        client.text_to_image.create(
            model="gpt-4o-image", aspect_ratio="1:1", output_count=3
        )
