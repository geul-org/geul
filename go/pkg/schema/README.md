# schema

문서가 스키마에 맞는지 검사한다. 위반은 Fact(Where, Expected, Actual)로 낸다.

| 층 | 내용 |
|---|---|
| 형식 | ISSUER 형식 번호의 Value·Edge 레이아웃 |
| 의미 | 격틀, 선택 제약, 클래스 계층(P279), 공리 |

- 판정은 Pass, Fail, Review다. 판정할 근거가 없으면 Review로 남긴다.
- 격틀은 Defeasible이다. 관용 표현, 비유, 허구는 예외(EXCEPT)로 허용한다.
- 스키마도 GEUL 문서로 공표하고 IMPORT로 가져온다. `0x0000` 핵심 스키마는 엔진에 내장한다.
