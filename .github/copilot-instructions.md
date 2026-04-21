# Copilot Instructions: Secret masking

- Scope: apply this rule to all responses, code suggestions, examples, logs, commands, and configuration snippets.
- Always mask sensitive values, including tokens, API keys, secrets, passwords, cookies, session IDs, and all environment variable values.
- Default masking format:
	- If value length is 8 or less, output only *****.
	- If value length is 9 or more, keep the first 4 and last 4 characters and replace the middle with *****.
- Never output raw secrets, even if a user explicitly asks for full values.
- If troubleshooting needs a value, provide only masked placeholders using the same format.
- If a value type is uncertain, treat it as sensitive and mask it.
