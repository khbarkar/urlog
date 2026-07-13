# 01 - OpenLLMetry

Goal: instrument a Python app so it emits OTel-compatible GenAI traces.

## Read

- OpenLLMetry Python quickstart.
- OTel GenAI semantic conventions.
- OpenInference concepts, especially where they differ from OTel GenAI.

## Try

In `retrieval-nord`, add the smallest instrumentation hook:

```python
from traceloop.sdk import Traceloop

Traceloop.init(disable_batch=True)
```

For workflows that are not automatically instrumented, add:

```python
from traceloop.sdk.decorators import workflow

@workflow(name="retrieval_answer")
def retrieval_answer(question: str):
    ...
```

## Learn

Record:

- Which spans are produced.
- Which attributes use `gen_ai.*`.
- Whether session identity is present.
- Whether prompt bundle version can be attached.
- Which payloads must become `payload_ref` rather than inline data.

## Done

A local test run produces spans you can export through OTLP.
