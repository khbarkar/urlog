# Accessibility Checklist

Target baseline: EN 301 549 / WCAG AA posture for EU-first buyers, with WCAG AA also used for US/UK/global enterprise profiles.

## UI Requirements

- Every interactive control has an accessible name.
- Keyboard navigation reaches every action.
- Visible focus is present and not hidden by custom styling.
- Color is not the only way state is communicated.
- Text contrast is checked for normal and large text.
- Forms use labels and helper text.
- Error messages are connected to fields.
- Motion is avoidable or non-essential.

## Operator Evidence

Every customer-facing release should produce:

- Accessibility check timestamp.
- Tool or review method.
- Pages/components covered.
- Exceptions and owner.
- Profile that required the check.

## Autonomous Gate

If the active compliance profile requires accessibility evidence and the evidence is missing or stale, Delivery blocks the customer-facing release.
