# 1. 엔터프라이즈용 자체 호스팅 실행기 구성 [`GH-100`]

## 1.1 엔터프라이즈용 자체 호스팅 실행기 구성

GitHub Actions의 자체 호스팅 실행기는 사용자 지정된 환경, 네트워크 액세스 및 보안 강화가 필요한 기업에 더 큰 유연성과 제어를 제공한다. 프록시, 레이블 및 네트워킹 고려 사항을 포함하여 자체 호스팅 실행기를 구성하는 모범 사례를 나타낸다.

### 1.1.1 자체 호스팅 실행기 설정

1. 자가 호스팅 러너 생성 및 등록
   * GitHub Enterprise -> 설정 -> Actions -> Runner로 이동한다.
   * 새 실행기를 선택하고 원하는 OS를 선택한다.
   * 제공된 명령에 따라 머신에 실행기를 설치하고 구성한다.
2. 실행기 설치 및 시작
   * OS에 따라 다음과 제공되는 명령을 실행한다.

* Linux / macOS
```sh
./config.sh --url https://github.com/<org-name> --token <generated-token>
./run.sh
```

* Windows
```ps
.\config.cmd --url https://github.com/<org-name> --token <generated-token>
.\run.cmd
```

### 1.1.2 자체 호스팅 실행기를 위한 프록시 구성

기업은 인터넷 액세스를 제한하는 회사 방화벽 및 프록시 뒤에서 작동하는 경우가 많기 때문에, 자체 호스팅 실행기가 GitHub와 통신할 수 있도록 하려면 다음과 같이 프록시 설정을 구성한다.

* Linux / macOS : 프록시 구성

```sh
export http_proxy=http://proxy.company.com:8080
export https_proxy=http://proxy.company.com:8080
export no_proxy=localhost,127.0.0.1
```

```sh
source ~/.bashrc
```

* Widnwos : 프록시 구성

```ps
[System.Environment]::SetEnvironmentVariable("HTTP_PROXY", "http://proxy.company.com:8080", "Machine")
[System.Environment]::SetEnvironmentVariable("HTTPS_PROXY", "http://proxy.company.com:8080", "Machine")
```

### 1.1.3 러너 관리에 레이블 사용

레이블은 OS, 하드웨어 또는 프로젝트 요구 사항에 따라 작업을 구성하고 특정 자체 호스팅 실행기로 라우팅하는 데 도움이 된다.

#### 러너에게 라벨 할당

실행기를 구성할 때 사용자 지정 레이블을 할당할 수 있다.

```sh
./config.sh --url https://github.com/<org-name> --token <generated-token> --labels "high-memory,gpu"
```

#### 워크플로의 특정 러너 타겟팅

레이블이 있는 특정 실행기에서 작업을 실행하려면 워크플로 `.yml` 업데이트한다.

```yml
jobs:
  build:
    runs-on: [self-hosted, high-memory]
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3
```

### 1.1.4 네트워킹 고려 사항

#### GitHub IP 허용 목록

GitHub 호스팅 실행기는 동적 IP에서 작동하지만 자체 호스팅 실행기는 액세스를 허용하기 위해 방화벽 규칙이 필요하며 최신 GitHub IP 범위를 검색한다. 이는 방화벽 설정에서 IP를 허용하여 연결을 보장한다.

```sh
curl -s https://api.github.com/meta | jq .actions
```

#### 프라이빗 네트워크 및 VPN 액세스

프라이빗 시스템에 액세스해야 하는 엔터프라이즈 워크로드의 경우 VPN 또는 내부 네트워크를 통해 견결하도록 실행기를 구성한다.

### 1.1.5 엔터프라이즈 러너를 위한 보안 모범 사례

* **실행기를 신뢰할 수 있는 워크플로로 제한** : 신뢰할 수 없는 코드 자체 호스팅 실행기에서 실행되지 않도록 한다. 일시적인 실행기를 사용하여 지속적인 위협을 방지하기 위해 작업 후 자동으로 실행기를 제거하도록 해야한다.
* **실행기 작업 모니터링** : 모든 실행기 작업을 기록하고 액세스를 감사한다.
* **OS 보안 패치 적용** : 실행기 머신을 정기적으로 업데이트하고 보호한다.

### 1.1.6 그룹을 사용하여 자체 호스팅 실행기 관리

실행기 그룹을 사용하면 조직에서 GitHub Actions에서 자체 호스팅 실행기를 위한 액세스를 관리하고 워크로드 배포를 제어하며 보안 정책을 적용할 수 있다. 이는 그룹 간에 실행기를 효과적으로 만들고, 관리하여, 이동하는 방법을 다룬다.

1. **실행기 그룹 이해**

러너 그룹은 GitHub Enterprise 또는 Organization 내에서 자체 호스팅 실행기를 구성하고 제어하는 데 도움이 되고 다음 항목에 대한 허용이 가능하다.

* 특정 실행기를 사용할 수 있는 리포지토리를 제한한다.
* 다른 팀 또는 워크로드에 대한 실행기 가용성 제어
* 특정 분기, 워크플로 또는 환경에 대한 권한 권리

| **GitHub 계획** | **실행기 그룹 사용 여부** |
| - | - |
| **GitHub Free** | ❌ |
| **GitHub Pro** | ❌ |
| **GitHub Organization** | ✅ |
| **GitHub Enterprise** | ✅ |

2. **실행기 그룹 만들기**
   1. GitHub -> 조직 설정 -> 액션 -> 러너로 이동한다.
   2. Self-Hosted 러너에서 `새 그룹`을 클릭한다.
   3. 그룹의 이름을 제공한다.
   4. 그룹에 액세스할 수 있는 사용자를 선택한다.
   5. 저장을 클릭한다.

3. **그룹에 실행기 추가**
    그룹이 만들어지면 수동으로 또는 등록 중에 실행기를 추가할 수 있다.

    1. ***옵션 1 : 등록 중 할당***
        새 실행기를 구성할 때 그룹을 지정한다.
        ```sh
        ./config.sh --url https://github.com/<org-name> --token <generated-token> --runnergroup "Linux-Runners"
        ```
    2. ***옵션 2 : 기존 실행기 이동***
       1. GitHub -> 조직 설정 -> 액션 -> 러너로 이동한다.
       2. 러너를 찾아 편집을 진행한다.
       3. 새 실행기 그룹을 선택하고 변경 내용을 저장한다.

4. **액세스 및 권한 관리**

    조직 수준 실행기는 특정 리포지토리로 사용을 제한하여 선택한 워크플로만 실행기에서 액세스할 수 있도록 할 수 있다.

5. **그룹 간에 주자 이동**

    한 그룹에서 다른 그룹으로 러너를 재배치를 하고자 한다면 다음 항목을 순서대로 수행 하면된다.

    1. GitHub -> 조직 설정 -> Actions -> Runner로 이동한다.
    2. 주자 이름을 클릭한다.
    3. 그룹 변경 -> 새 그룹을 선택한다.
    4. 저장을 클릭하거나 또는 다른 그룹에서 실행기의 등록을 취소한 후 다시 등록한다.
    ```sh
    ./config.sh remove
    ./config.sh --url https://github.com/<org-name> --token <generated-token> --runnergroup "New-Group"
    ```

    * 다양한 OS 유형에 대해 별도의 그룹을 만든다.
    * 레이블을 사용하여 러너를 추가로 분류한다.
    * 실행기 액세스를 신뢰할 수 있는 리포지토리로만 제한한다.
    * 팀의 요구 사항에 따라 그룹을 정기적으로 점검하고 업데이트한다.
    * 실행기 사용량 및 성능을 모니터링하여 CI/CD 워크로드를 최적화한다.

# 2. 자체 호스팅 실행기를 모니터링하고 문제와 업데이트를 해결 [`GH-100`]

자체 호스팅 실행기를 효과적으로 관리하려면 지속적인 모니터링, 사전 문제 해결 및 정기적인 업데이트가 필요하고 자체 호스팅 실행기의 고가용성, 보안 및 성능을 보장하기 위한 모범 사례 및 권장 방법을 제공한다.

## 2.1 자체 호스팅 실행기 모니터링

#### 실행기 상태 확인

1. GitHub -> 조직 설정 -> Actions -> Runner로 이동한다.
2. 상태를 검토한다.
   * ✅ 대기 중 → 워크플로 준비 완료된 상태
   * 🔄 활성 → 현재 작업을 실행 중
   * ❌ 오프라인 → Runner가 다운되었거나 연결이 끊김

#### GitHub API를 사용하여 실행기 상태 가져오기

자체 호스팅 실행기의 상태를 프로그래밍 방식으로 확인할 수 있다.

```sh
curl -H "Authorization: token <your_github_token>" \
     -H "Accept: application/vnd.github.v3+json" \
     https://api.github.com/orgs/<org-name>/actions/runners
```

#### 로깅 및 메트릭

* **시스템 로그** : 실행기 설치 폴더 내의 `_diag/` 디렉터리에서 로그를 확인한다.
* **GitHub Actions 워크플로 로그** : 작업 -> 워크플로 실행 -> 로그로 이동하여 실행기 실행 세부 정보를 확인한다.
* **Prometheus/Grafana**를 통한 모니터링 : CPU, Memory 및 작업 실행 시간을 추적하도록 Prometheus 내보내기를 구성한다.

## 2.2 자체 호스팅 실행기 문제 해결

#### 일반적인 문제 및 수정 사항

| **문제** | **가능한 원인** | **수정** |
| - | - | - |
| **실행기가 오프라인으로 표시** | 네트워크 문제, 토큰이 만료되었거나 실행기가 충돌 | 러너 재시작 |
| **작업이 대기 상태로 중단** | 필수 레이블이 있는 사용 가능한 러너 없음 | 실행기 추가 또는 레이블 업데이트 |
| **권한 오류로 작업 실패** | 잘못된 실행기 권한 | 실행기의 올바른 액세스 권한 확인 |
| **워크플로 실행 속도 저하** | 높은 CPU/Memory 사용량 | 시스템 메트릭 모니터링 및 실행기 크기 조정 |

#### 실행기 다시 시작

```sh
./svc.sh stop
./svc.sh start
```

```sh
sudo systemctl restart actions.runner.<org-name>.<runner-name>.service
```

* 로그에서 오류 확인
  * **실행기 로그** : `<runner_dir>/_diag/Runner_<timestamp>.log`
  * **GitHub Actions 로그** : GitHub UI에서 워크플로 실행 로그를 확인

## 2.3 자체 호스팅 실행기 업데이트

#### 실행기 업데이트 확인

GitHub는 주기적으로 러너 바이너리를 업데이트 하며 업데이트 확인을 하고자 하면 다음을 수행하면 된다.

```sh
./config.sh --version
```

API를 통해 러너 버전을 확인할 수 있다.

```sh
curl -H "Authorization: token <your_github_token>" \
     -H "Accept: application/vnd.github.v3+json" \
     https://api.github.com/repos/actions/runner/releases/latest
```

#### 실행기 업데이트

1. 수동 업데이트 실행기 중지
```sh
./svc.sh stop
```
2. 최신 실행기를 다운로드
```sh
curl -o actions-runner-linux-x64.tar.gz -L \
     https://github.com/actions/runner/releases/latest/download/actions-runner-linux-x64.tar.gz
```
3. 추출 및 다시 구성
```sh
tar xzf ./actions-runner-linux-x64.tar.gz
./config.sh --url https://github.com/<org-name> --token <generated-token>
./svc.sh install
./svc.sh start
```

#### GitHub 작업을 사용하여 자동화된 업데이트

[자동으로 실행기를 확인하고 업데이트](../github-crdential/example-sources/example-actions-update.yml)
