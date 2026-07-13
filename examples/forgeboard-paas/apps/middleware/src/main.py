from datetime import datetime, timezone
from enum import StrEnum
from hashlib import sha256
from uuid import uuid4

from fastapi import FastAPI
from pydantic import BaseModel, Field


class BundleRisk(StrEnum):
    LOW = "low"
    WATCH = "watch"
    BLOCK = "block"


class Sector(StrEnum):
    FINTECH = "fintech"
    MEDICAL = "medical"
    GREEN_TECH = "green_tech"


class BundleUpload(BaseModel):
    tenant_id: str = Field(min_length=3)
    bundle_name: str = Field(min_length=3)
    declared_schema_version: str = "v1"
    estimated_events_per_minute: int = Field(default=500, ge=0)
    customer_instructions: str = ""


class TransactionKind(StrEnum):
    MERCHANT_SETTLEMENT = "merchant_settlement"
    MEDICAL_REIMBURSEMENT = "medical_reimbursement"
    CARBON_CREDIT = "carbon_credit"
    ENERGY_SUBSIDY = "energy_subsidy"


class TransactionSubmission(BaseModel):
    tenant_id: str = Field(min_length=3)
    sector: Sector
    transaction_kind: TransactionKind
    amount_minor: int = Field(ge=0)
    currency: str = Field(min_length=3, max_length=3)
    counterparty_ref: str = Field(min_length=3)
    policy_bundle_version: str = "local-dev"
    customer_instructions: str = ""


class BundleReceipt(BaseModel):
    upload_id: str
    accepted_at: datetime
    risk_hint: BundleRisk
    emitted_topic: str


class TransactionReceipt(BaseModel):
    transaction_id: str
    accepted_at: datetime
    risk_hint: BundleRisk
    evidence_hash: str
    emitted_topic: str
    requires_human_review: bool


app = FastAPI(title="ForgeBoard API", version="0.1.0")


@app.get("/healthz")
def healthz() -> dict[str, str]:
    return {"status": "ok"}


@app.post("/bundles", response_model=BundleReceipt)
def submit_bundle(upload: BundleUpload) -> BundleReceipt:
    risk = BundleRisk.LOW
    if upload.estimated_events_per_minute > 10_000:
        risk = BundleRisk.WATCH
    if "ignore previous instructions" in upload.customer_instructions.lower():
        risk = BundleRisk.BLOCK

    return BundleReceipt(
        upload_id=str(uuid4()),
        accepted_at=datetime.now(timezone.utc),
        risk_hint=risk,
        emitted_topic="forgeboard.bundle.submitted",
    )


@app.post("/transactions", response_model=TransactionReceipt)
def submit_transaction(txn: TransactionSubmission) -> TransactionReceipt:
    risk = BundleRisk.LOW
    if txn.amount_minor > 10_000_000:
        risk = BundleRisk.WATCH
    if txn.sector == Sector.MEDICAL and txn.amount_minor > 2_500_000:
        risk = BundleRisk.WATCH
    if "ignore previous instructions" in txn.customer_instructions.lower():
        risk = BundleRisk.BLOCK

    transaction_id = str(uuid4())
    accepted_at = datetime.now(timezone.utc)
    evidence_payload = "|".join(
        [
            transaction_id,
            txn.tenant_id,
            txn.sector,
            txn.transaction_kind,
            str(txn.amount_minor),
            txn.currency.upper(),
            txn.counterparty_ref,
            txn.policy_bundle_version,
            accepted_at.isoformat(),
        ]
    )

    return TransactionReceipt(
        transaction_id=transaction_id,
        accepted_at=accepted_at,
        risk_hint=risk,
        evidence_hash=sha256(evidence_payload.encode("utf-8")).hexdigest(),
        emitted_topic="forgeboard.transactions.submitted",
        requires_human_review=risk == BundleRisk.BLOCK,
    )
