# Copilot Instructions: Secret masking

- Always mask sensitive values (tokens, API keys, secrets, passwords, cookies, session IDs, and any env var values).
- Default masking: keep first 4 and last 4 characters and replace the rest with `*****`.
- If length is 8 or less, replace the entire value with `*****`.
- Never output raw secrets in code blocks, logs, CLI commands, IDE steps, or configuration examples.
- Use masked placeholders in examples (e.g., `abcd*****wxyz`).
