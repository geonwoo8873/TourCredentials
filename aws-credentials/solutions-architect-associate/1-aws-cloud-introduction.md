# AWS Cloud 사용 케이스

# 1. AWS 글로벌 구조

## 1.1 AWS 지역 [`Regions`]

AWS가 전 세계에서 가용중인 데이터 센터를 클러스터링하는 물리적 위치를 리전이라고 칭한다. 논리적 데이터 센터의 각 그룹을 가용 영역이라고 하며, 각 리전은 지리적 영역 내에서 격리되고 물리적으로 분리된 최소 3개의 AZ로 구성된다. 이는 단일 데이터 센터로 정의하는 다른 클라우드의 공급자와 달리 여러 AZ 설계는 사용자에게 여러 이점을 제공하고 독립된 자원과 물리적 보안을 갖추어 지연 시간을 짧게 구축해 중복 네트워크를 통해 연결된다. 그리고 사용자는 여러 AZ에서 실행되도록 애플리케이션을 설계하여 내결함성을 강화함 동시에 가장 높은 수준의 보안, 규정 준수, 데이터 보호를 충족한다.

#### [Enable location regions](#참조)

* **North America** (`북미`)
  * us-east-1 (Mexico, Georgia, United States of America, Massachusetts, Illinois, Texas)
  * us-west-2 (Colorado, United States of America, Hawaii, Nevada, California, Arizona, Oregon, Washington)
* **South America** (`남아메리카`)
  * us-east-1 (Argentina, Chile, Peru)
* **Europe** (`유럽`)
  * eu-north-1 (Denmark, Finland)
  * eu-central-1 (Germany, Greece, Poland, Turkey)
* **Middle East** (`중동`)
  * me-south-1 (Oman)
* **Africa** (`아프리카`)
  * af-south-1 (Nigeria)
* **Asia Pacific** (`아시아 태평양`)
  * ap-southeast-1 (Australia, Philippines, Thailand, Vietnam)
  * ap-southeast-2 (New Zealand)
  * ap-south-1 (India, India)
  * ap-northeast-1 (Taiwan)

> [!WARNING]
> **상단의 AWS Region 목록 중 현지 데이터 센터의 운영 이슈나 혹은 국가 단위 수준의 재해가 발생 하였을 시 일시적 부터 무기한 사용 불가 이슈가 발생할 수 있어 지속적인 확인이 필요하다.**

## 1.2 AWS 가용 영역 [`Availability Zones`]

AZ는 가용 영역이라는 약자로 AWS 리전의 중복 전력, 네트워킹 및 연결이 제공되는 하나 이상의 개별 데이터 센터로 구성되어 단일 데이터 센터를 사용하는 것보다 높은 가용성, 내결함성 및 확정성을 갖춘 프로덕션 애플리케이션과 데이터베이스를 운영할 수 있다. 모든 AZ는 높은 대역폭, 빠른 네트워킹, 중복성을 갖춘 전용 메트로 광 네트워크와 상호 연결되어 있어 처리량과 지연 시간이 짧은 네트워킹을 제공한다. 가용 영역간의 모든 트랙픽은 암호화 처리로 되어 있어 동기 복제 기능을 충분히 수행할 수 있기 때문에 애플리케이션 분할을 용이하다는 점이다.

> [!NOTE]
> * **AZ는 최소 3개부터 최대 6개로 이루어져 있으며 이는 각 구축된 개별의 데이터 센터로 구성되어 있다.**
>   * ap-southeast-2a
>   * ap-southeast-2b
>   * ap-southeast-2c

<!-- ## 1.3 AWS 엣지 로케이션 / PoP [`Point of Presence`] -->


---

#### 참조

* [AWS Regions & Local zones](https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html)

#### 단어

> [!NOTE]
> * **하부 구조 (`Infrastructure`)**
> * **지역 (`Regions`)**
