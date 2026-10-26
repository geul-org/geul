# GEUL Go 라이브러리

GEUL 문서를 읽고 쓰고, 다른 표현과 서로 바꾸는 Go 라이브러리다. 누구나 import해서 쓴다.
명세는 geul-dev 저장소의 `draft/t002/docs/`(OVERVIEW, ISSUER, GEUL-ORG, TOULMIN)를 따른다.

```
go get github.com/geul-org/geul/go
```

## 구조

```
pkg/core/              읽기·쓰기, 프레이밍, 정규형, 버전 ID (엔진 내장 규칙)
pkg/toulmin/           TOULMIN 코덱 (판정·실행 그래프 레코드, 코드표, 값)
pkg/delta/             버전 diff(DELTA-BASE·DELTA-REMOVE) 생성·적용과 LOCAL ID 이어받기(같은 서술 같은 번호, auto increment)
pkg/inspect/           해석된 문서 모델 (레코드별 이름·내용·Edge Id·참조, 무손실일 때만)
pkg/resolve/           IMPORT 의존성 해석 (Resolver 주입)
pkg/schema/            형식 검증 (0층 → Codec 규격 → 격틀)
```

- GEUL과 다른 표현 사이의 코덱(`geulnl`, `wikigeul`, `nlgeul`)은 이 라이브러리가 아니라 TANGEUL 문서로 작성한다(geul-dev `draft/t002/codec/`). 이 라이브러리는 코덱 문서가 부르는 호스트 함수와 그 문서를 실행하는 기반을 제공한다.

## 흐름

```
.geul 바이트 → core → Doc ─┬→ schema  (Fact)
                ↑          └→ 코덱 문서(TANGEUL) 실행 → 자연어 등
     resolve: IMPORT-DOC-* → 아카이브 → 외부 Doc 연결

코덱 문서(TANGEUL) 실행 → core 빌더 → 정규형 바이트 → 버전 ID
```

## 원칙

- core는 외부 의존과 네트워크 I/O가 없다. 다른 패키지가 core에 의존하고, 반대 방향은 없다.
- 디코딩은 IMPORT를 조회하지 않는다. 조회는 `resolve`의 명시적 단계이고, 위치는 Resolver로 주입한다.
- 쓰기는 빌더로만 한다. 폭과 ID는 빌더가 정한다. 같은 내용은 같은 바이트가 된다.
- 결과는 `doc, warns, err`로 돌려준다. 경고를 실패로 볼지는 호출하는 쪽이 정한다.
- 모르는 레코드는 버리지 않고 원본 WORD로 보존한다.
- 폭과 개수에 고정 상한을 두지 않는다. 상한은 설정으로 바꿀 수 있게 한다.

## 아카이브

- 로컬 아카이브는 geul-dev 저장소의 `draft/t002/.well-known/geul/`이다. 경로는 GEUL-ORG 공표 주소와 같다.
- 아카이브, 미러, geul.org 덤프는 같은 구조라서 Resolver 하나로 읽는다.

## 단계

1. core 읽기·쓰기 왕복
2. 정규형과 버전 ID
3. 파일 Resolver와 IMPORT
4. schema 격틀 검사와 Fact. 모호함 지적 데모

서명은 보류한다. 검증 클래스 레코드는 보존만 하고 버전 ID 계산에서 뺀다.

## 개발 규칙

- filefunc: 함수·메서드·타입별로 파일을 나누고 `//ff:` 주석을 단다. 어휘는 `codebook.yaml`. `filefunc validate .` 위반 0건.
- TSMA: `tsma scan .` → `tsma next` → 테스트 → `tsma submit`.
- `go build ./...`, `go vet ./...`, `go test ./...` 통과.
- 공통 테스트 벡터는 다른 언어 구현(`ts/` 등)과 함께 쓴다.

## 라이선스

CC0. CC0가 아닌 자료(표준국어대사전 등)의 내용은 코드, 데이터, 테스트에 넣지 않는다.
