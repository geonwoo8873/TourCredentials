# Copilot Instructions: Secret masking

## Ko-Kr
- 다음과 같이 대화 및 코드 제안에 기밀 유출을 방지를 위해 `*` 일부 표시를 대체한다.  
  [예시 : `AWDX214EW8RF1CB0*******SDA15W52`, `ASBWJBWJSI48451**************`]
  - Tokens
  - API Keys
  - Passwords
  - Session IDs
  - Env values

> [!WARNING]
> 허나 로컬 환경에서 단위 테스트 및 해당 기능의 터미널 출력을 허용해야 한다면 해당 일시적 사용자에게 컨펌을 확인 해야한다.  
> 또한 절대 보안에 철저해야할 값들은 해당 목록외 사용자 및 코파일럿의 재량에 따라 제안 기능에도 마스킹으로 대체하거나 예시 값으로 표현 해야한다.

## EN

- To prevent confidential data leaks in chat responses and code suggestions, partially mask sensitive values with `*`.  
  [Example: `AWDX214EW8RF1CB0*******SDA15W52`, `ASBWJBWJSI48451**************`]
  - Tokens
  - API Keys
  - Passwords
  - Session IDs
  - Env values

> [!WARNING]
> However, if unit tests in a local environment or terminal output for this feature must be allowed, confirm with the relevant user before doing so.  
> Also, for values that require strict security, mask them in suggestions or represent them with sample values at the discretion of the user and Copilot, even if they are not listed above.