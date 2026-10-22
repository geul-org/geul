# wikigeul

위키데이터 아이템과 statement를 GEUL 문서로 바꾼다.

## 사상

| 위키데이터 | GEUL |
|---|---|
| 아이템, 속성 | `WIKIDATA-QID-ENTITY`, `WIKIDATA-PID-ENTITY` 참조 |
| statement | Edge |
| qualifier | Edge에 붙는 Edge (`INNER-EDGE` 참조) |
| reference | 근거 계열 Edge |
| 시간, 수량, 좌표 | Value |

## 원칙

- 변환한 문서는 변환자의 Issuer로 낸다. 원천 리비전 ID를 근거 Edge로 남긴다.
- 전체 덤프가 아니라 필요한 하위 그래프부터 변환한다(클래스 계층 P279, 속성 제약 P2302).

## 미결정

- rank(preferred, normal, deprecated)의 표현.
- novalue와 somevalue의 표현. 값 없음과 모름은 다르다.
- 문서 단위(아이템 하나인지, 주제 묶음인지).
