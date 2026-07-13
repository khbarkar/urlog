# Integration Learning Track

This folder turns the Phase 0 roadmap items into a tutorial path. The goal is to learn the tools while building the smallest useful Integration slice.

## Roadmap Items Covered

1. Land MR #1 scaffold.
2. Instrument `retrieval-nord` with OpenLLMetry using OTel GenAI semantic conventions.
3. Self-host Langfuse with Docker Compose and ship spans to it.
4. Self-host Arize Phoenix side by side and ship the same spans to it.
5. Run both for two weeks against real `retrieval-nord` traffic/tests.

## Learning Sequence

| Step | File | Outcome |
|---|---|---|
| 1 | `01-openllmetry.md` | Understand how OpenLLMetry creates OTel traces |
| 2 | `02-langfuse.md` | Run Langfuse locally and identify its data model |
| 3 | `03-phoenix.md` | Run Phoenix locally and identify its data model |
| 4 | `04-dual-shipping.md` | Route the same traces to both tools |
| 5 | `05-two-week-comparison.md` | Compare what each tool can and cannot answer |
| 6 | `06-forgeboard-integration.md` | Apply Integration to the ForgeBoard sample app |

## Sources To Keep Open

- Langfuse self-hosting: https://langfuse.com/self-hosting
- OpenLLMetry Python quickstart: https://www.traceloop.com/docs/openllmetry/getting-started-python
- Phoenix docs: https://arize.com/docs/phoenix

## Rule

This is a learning harness, not a product dependency. Integration should learn from Langfuse/Phoenix data models without becoming a tracing UI or prompt playground.
