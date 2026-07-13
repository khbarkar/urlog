# Streaming Backend

ForgeBoard uses a Redpanda-compatible Kafka API for local realism. The deployable local manifest lives in `deploy/kubernetes/base/redpanda-single-node.yaml` so default `kubectl kustomize` works without relaxed load restrictions.

Future cloud overlays can replace this with managed Kafka, Confluent Cloud, MSK, Event Hubs, or another Kafka-compatible service while keeping the topic contract stable. Air-gapped deployments keep the same topic contract but use customer-approved images, offline SBOMs, and private registry mirrors.
