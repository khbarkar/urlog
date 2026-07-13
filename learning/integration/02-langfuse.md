# 02 - Langfuse

Goal: self-host Langfuse and understand its trace model.

Langfuse v3 self-hosting runs multiple components: application containers, Postgres, ClickHouse, Redis/Valkey, and S3-compatible object storage. Use the official low-scale Docker Compose setup rather than copying a stale compose file into this repo.

## Try

1. Clone or download the official Langfuse self-hosting Docker Compose example into a scratch directory outside this repo.
2. Start it locally.
3. Create a project and API key.
4. Send a small OpenLLMetry trace to Langfuse.

## Learn

Record:

- How Langfuse models traces, observations, scores, sessions, and prompts.
- What goes to ClickHouse versus object storage.
- How scores are represented.
- What is easy to ask in the UI.
- What is hard or impossible to ask for ops-loop decisions.

## Notes For Urlog

Langfuse is useful for learning and comparison. Urlog should not become a Langfuse clone.
