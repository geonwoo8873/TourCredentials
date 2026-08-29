# 암호화된 암호 관리

# 1. 엔터프라이즈에서 암호화된 비밀 관리 [`GH-100`]

GitHub Actions를 사용하면 암호화된 비밀을 통해 API 키, 인증 토큰, 암호 및 인증서와 같은 중요한 데이터를 안전하게 저장하고 사용할 수 잇다. 이러한 비밀은 안전하게 저장되고 워크플로에 삽입된다. 이 디자인은 로그 또는 소스 코드에 표시되지 않도록 한다.

엔터프라이즈 환경에서는 효과적인 비밀 관리가 중요하고, 보안을 유지하고 규정 준수 요구 사항을 충족하며 운영 효율성을 지원하는 데 도움이 된다. GitHub를 사용하면 엔터프라이즈, 조직, 리포지토리 및 환경의 네 가지 수준에서 비밀을 관리할 수 있다.

#### 암호화된 비밀의 범위

비밀의 범위를 이해하는 것은 엔터프라이즈 환경에서 안전하게 관리하는 데 필수적이다.

| **비밀 수준** | **범위** | **액세스할 수 있는 사용자** | **일반적인 사용 사례** |
| - | - | - | - |
| **Enterprise-Level** | GitHub Enterprise Cloud 조직의 모든 리포지토리에 적용 | 엔터프라이즈 소유자, 보안 관리자 | 여러 리포지토리에서 자격 증명 공유 |
| **Organization-Level** | 조직의 모든 리포지토리에 적용하고, 선택적으로 선택한 리포지토리로 제한 | 조직 소유자, 보안 관리자 | 클라우드 서비스 및 공유 데이터베이스에 액세스한다. |
| **Repository-Level** | 단일 리포지토리에만 적용 | 리포지토리 관리자, 워크플로 실행기 | 한 리포지토리에서 배포를 위한 보안 자격 증명 |
| **Environment-LEVEL** | 스테이징 또는 프로덕션과 같은 리포지토리 내의 특정 배포 환경에 적용 | 지정된 환경의 워크플로 러너 | 배포 환경별로 자격 증명을 구분 |

* 주요 고려 사항
  * 엔터프라이즈 비밀은 GitHub Enterprise Cloud 전용이며 조직 전체에서 중앙 집중식 관리를 지원한다.
  * 조직 비밀은 세분화된 액세스 제어를 제공하며 특정 리포지토리로 제한될 수 있다.
  * 환경 비밀은 워크플로 환경에 대한 액세스를 제한하여 의도하지 않은 노출을 방지하는 데 도움이 된다.

# 2. 조직 수준에서 암호화된 암호 관리 [`GH-100`]

조직 수준에서 암호화된 비밀을 만들면 중요한 정보를 보호하는 동시에 여러 리포지토리에서 비밀을 관리하는 데 필요한 요소들을 줄일 수 있다.

GitHub 조직에서 워크플로를 작성하는 개발자는 일부 워크플로에서 코드를 프로덕션에 배포하기 위한 자격 증명이 필요하며 이러한 중요한 정보를 공유하지 않기 위해 조직 수준에서 자격 증명을 포함하는 암호화된 암호를 생성할 수 있다. 이러한 방식으로, 노출되지 않고 워크플로에서 자격 증명을 사용할 수 있다.

조직 수준에서 비밀을 생성하려면 조직의 설정으로 이동하여 `비밀 및 변수` -> `작업` -> `새 조직 비밀`을 선택한다.

## 2.1 GitHub CLI를 통해 Organization-Level 암호화된 비밀 관리

* 조직에 대한 비밀을 만든다.
  ```sh
  gh secret set SECRET_NAME --org my-org --body "super-secret-value"
  ```
* 모든 조직 비밀을 나열한다.
  ```sh
  gh secret list --org my-org
  ```
* 기존 비밀을 업데이트한다.
  ```sh
  gh secret set SECRET_NAME --org my-org --body "new-secret-value"
  ```
* 비밀을 삭제한다.
  ```sh
  gh secret delete SECRET_NAME --org my-org
  ```

#### 조직 비밀에 대한 보안 고려 사항

* 기본적으로 모든 리포지토리에 대한 액세스 권한을 부여하지 않고 비밀을 특정 리포지토리로 제한한다.
* 권한 있는 구성원만 비밀을 만들거나 업데이트하거나 삭제할 수 있도록 RBAC(Role-Based Access Control)
* 액세스 로그를 정기적으로 모니터링 하여 무단 사용 또는 의심스러운 활동을 식별하고 대응한다.

# 3. 리포지토리 수준에서 암호화된 비밀 관리 [`GH-100`]

비밀을 특정 리포지토리로 범위 지정하려면 GitHub Enterprise Cloud 또는 GitHub Enterprise Server를 사용한다.

## 3.1 리포지토리 수준 비밀 만들기

1. 리포지토리의 설정으로 이동한다
2. Deploy keys를 선택한다.
3. 비밀의 이름과 값을 입력한다.

## 3.2 CLI를 통해 리포지토리 수준 암호화된 비밀 관리

* 리포지토리 비밀을 나열한다.
  ```sh
  gh secret list --repo my-repo
  ```
* 리포지토리 비밀을 업데이트한다.
  ```sh
  gh secret set SECRET_NAME --repo my-repo --body "new-secret-value"
  ```
* 리포지토리 비밀을 삭제한다.
  ```sh
  gh secret delete SECRET_NAME --repo my-repo
  ```

# 4. 작업 및 워크플로 내에서 암호화된 암호 액세스 [`GH-100`]

## 4.1 워크플로에서

`secrets` 컨텍스트를 사용하여 비밀을 열고, `with` 비밀을 입력으로 전달하거나 `env` 환경 변수로 설정하는 데 사용한다.

```yml
steps:
  - name: Hello world action
    uses: actions/hello-world@v1
    with:
      # Pass the secret as an input to the action
      super_secret: ${{ secrets.SuperSecret }}
    env:
      # Set the secret as an environment variable
      super_secret: ${{ secrets.SuperSecret }}
```

* **with** : 암호를 입력 매개 변수로 작업에 전달하고 작업에서 `action.yml`입력을 명시적으로 정의할 때 일반적으로 사용된다.
* **env** : 비밀을 단계에 환경 변수로 노출하여 이 변수는 단계의 명령이나 작업 내의 스크립트에 환경 변수가 예상되는 경우에 유용하다.

## 4.2 작업에서

사용자 지정 작업 내에서 비밀을 사용하려면 메타데이터 파일의 `action.yml` 입력으로 정의하고 작업 코드에서 환경 변수로 액세스할 수 있다.

```yml
inputs:
  super_secret:
    description: 'My secret token'
    required: true
```

```js
// Access the input using the Actions Toolkit
const core = require('@actions/core');
const token = core.getInput('super_secret');
```

* `action.yml` 암호를 필수 또는 선택적 입력으로 정의한다.
* 코드에서 액세스 작업에서 도구 키트를 사용하거나 설정된 경우 환경 변수를 참조하여 비밀을 읽는다.

> [!WARNING]
> **작업 소스 코드에서 암호를 하드 코딩하지 않아야 하며, 입력 및 비밀을 안전하게 관리하려면 작업 도구 키트를 사용하여 코드 논리내에서 값을 처리한다.**

# 5. GitHub Actions에 대한 보안 강화 구성 [`GH-100`]

GitHub Actions에 대한 보안 강화는 소프트웨어 공급망을 안전하게 유지하는 역할을 하며 워크플로에서 사용하는 작업의 보안을 강화하기 위한 권장 사례를 제공한다.

## 5.1 스크립트 삽입 공격 완화하기 위한 사례 식별

GitHub 작업에서 스크립트 삽입 공격을 완화하기 위한 몇 가지 모범 사례는 다음과 같다.

1. **인라인 스크립트 대신 javascript 작업 사용** : 인라인 스크립트에 해당 값을 포함하는 대신 컨텍스트 값을 인수로 허용하는 javascript 작업을 사용하며 컨텍스트 데이터가 셸 명령을 직접 생성하거나 실행하는 데 사용되지 않기 때문에 스크립트 삽입의 위험을 줄인다.

```yml
uses: fakeaction/checktitle@v3
 with:
   title: ${{ github.event.pull_request.title }}
```

2. **인라인 스크립트에섲 중간 환경 변수 사용** : 인라인 스크립트를 사용하는 경우 명령에서 변수를 사용하기 전에 변수를 환경 변수로 평가하기 때문에 스크립트가 실행되기 전에 값이 확인되어 스크립트 삽입의 위험이 줄어든다.

```yml
- name: Check PR title
    env:
      TITLE: ${{ github.event.pull_request.title }}
    run: |
      if [[ "$TITLE" =~ ^octocat ]]; then
      echo "PR title starts with 'octocat'"
      exit 0
      else
      echo "PR title did not start with 'octocat'"
      exit 1
      fi
```

3. **워크플로 템플릿을 할용하여 코드 검색 구현** : 리포지토리의 작업 탭으로 이동하여 새 워크플로 버튼을 선택한다. 워크플로 선택 페이지에서 워크플로 템플릿에 액세스하고 적용할 보안 섹션을 찾는다. 특정 이벤트에서 실행되도록 CodeQL 스캐너를 구성하여 스크립트 주입과 같은 취약성을 포함하여 워크플로내에서 사용되는 작업에서 분기의 파일을 검색하고 노출 CWU(Common Weakness Enumeration, 공통 약점 열거형)에 플래그를 지정하도록 한다.

4. **토큰에 대한 권한 제한** : 생성된 토큰에 `rule of least privilege` 항상 적용해야 하며, 토큰이 만들어진 작업을 달성하기 위한 최소 권한이 토큰에 할당되었는지 확인한다.

# 6. 타사 작업을 안전하게 사용하기 위한 모범 사례 [`GH-100`]

1. **작성자가 신뢰할 수 있는 경우에만 태그에 작업 고정** : 작업의 작성자가 확인되고 신뢰할 수 잇는 경우에만 버전 태그를 사용하고 이후 릴리스에서 예기치 않은 변경의 위험을 줄이는 데 도움이 된다.

```yml
- name: Checkout
  uses: actions/checkout@v4  # Pinned to a specific version tag
```

2. **작업을 전체 커밋 SHA에 고정** : 전체 커밋 SHA에 고정하면 변경할 수 없는 버전의 작업을 사용할 수 있다. 커밋 SHA가 예상 리포지토리에서 제공되는지 항상 확인한다.

```yml
- name: Checkout
  uses: actions/checkout@1e31de5234b9f8995739874a8ce0492dc87873e2  # Pinned to a specific commit SHA
```

3. **작업의 소스 코드 감사** : 작업의 원본을 검토하여 데이터를 안전하게 처리하고 예기치 않거나 악의적인 동작을 포함하지 않는지 확인한다.

## 6.1 신뢰할 수 있는 타사 작업 표시기

신뢰할 수 있는 작업을 사용하여 워크플로의 위험을 줄인다.

* **오피컬 배지를 검색** : 신뢰할 수 있는 작업은 GitHub Marketplace에 표시되고, GitHub에서 작성자를 확인했음을 알리는 타이틀 옆에 확인된 작성자 배지를 표시한다.
* **설명서 확인** : 파일은 `action.yml` 문서화되고 작업 작동 방식을 명확하게 설명해야 한다.

## 6.2 Dependabot 버전 업데이트를 사용하여 작업을 최신 상태로 유지

Dependabot 버전 업데이트를 사용하도록 설정하여 GitHub Actions 종속성을 최신 상태로 안전하게 유지한다.

# 7. 손상된 실행기의 잠재적 영향 [`GH-100`]

## 7.1 실행기에서 데이터 반출

GitHub Actions는 로그에서 비밀을 자동으로 수정하지만 이 수정은 완전한 보안 경계가 아니기에 러너가 손상되면 공격자가 의도적으로 비밀 정보를 로그에 기록하여 노출시킬 수 있다. 손상된 실행기를 사용하여 스크립트된 HTTP 요청을 사용하여 비밀 또는 기타 중요한 리포지토리 데이터를 외부 서버로 전달할 수 있다.

```yml
echo ${SOME_SECRET:0:4}
echo ${SOME_SECRET:4:200}
```

### 7.1.1 비밀에 대한 액세스

이벤트를 사용하요 `PR` 포크된 리포지토리에서 트리거되는 워크플로에는 읽기 전용 권한이 있으며 비밀에 액세스할 수 없다. 그러나 사용 권한은 이벤트 유형에 따라 달라져 실행기가 손상되면 리포지토리 비밀이 노출될 수 있으니 쓰기 권한이 있는 작업이 `GITHUB_TOKEN` 오용될 수 있는 위험이 있다.

GitHub Actions는 워크플로나 포함된 작업에서 참조되지 않은 경우 메모리에서 시크릿을 제거하지만, 실행기가 손상될 경우 사용 중인 모든 시크릿은 유출될 위험이 있다.

* 비밀 또는 토큰이 환경 변수에 할당된 경우를 사용하여 `printenv` 직접 액세스할 수 있다.
* 비밀을 직접 참조하는 경우 확인된 값을 포함하는 생성된 셸 스크립트가 디스크에 저장되고 액세스할 수 있다.
* 사용자 지정 작업의 경우 위험 수준은 작업의 논리 내에서 비밀을 처리하는 방법에 따라 달라진다.

```yml
uses: exampleaction/publish@v3
with:
  key: ${{ secrets.PUBLISH_KEY }}
```

### 7.1.2 도용

