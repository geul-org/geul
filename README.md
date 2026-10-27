# gogeul

GEUL(General Encoding Unified Language) 문서(`.geul`)를 읽고 쓰고, 다른 표현과 서로 바꾸는 Go 라이브러리와 `geul` 명령줄 도구.
명세는 [geul-org/geul-dev](https://github.com/geul-org/geul-dev)의 `files/geul/docs/`(OVERVIEW, ISSUER, GEUL-ORG, TOULMIN)를 따른다.

```
go get github.com/geul-org/gogeul          # 라이브러리
go install github.com/geul-org/gogeul/cmd/geul@latest   # CLI
```

- 저장소: 2026-10-27 `geul-org/geul`에서 이름을 바꾸고 `go/`를 루트로 올렸다. 모듈 경로는 `github.com/geul-org/gogeul`(옛 `github.com/geul-org/geul/go`).
- 다른 언어 구현은 별도 저장소로 둔다(TypeScript `geults`, 예정).

## 구조

```
pkg/core/              읽기·쓰기, 프레이밍, 정규형, 버전 ID, 16 WORD Issuer 계산 (엔진 내장 규칙)
pkg/codecspec/         Codec 규격 문서(DECLARE·FIELD·ENUM·CODEC·BIND) → Codec별 필드 목록
pkg/toulmin/           TOULMIN 코덱 (판정·실행 그래프 레코드, 코드표, 값)
pkg/delta/             버전 diff(DELTA-BASE·DELTA-REMOVE) 생성·적용, 변경 목록·blame, LOCAL ID 이어받기(같은 서술 같은 번호, auto increment)
pkg/inspect/           해석된 문서 모델 (레코드별 이름·내용·Edge Id·참조, 무손실일 때만)
pkg/notation/          텍스트 표기 (레코드 한 줄 ↔ 레코드, geul CLI의 show·add·meta, 왕복 보장)
pkg/resolve/           아카이브(공표·버전 목록·이름·폐기)와 IMPORT 의존성 해석 (Resolver 주입)
pkg/schema/            형식 검증 (0층 → Codec 규격 → 격틀)
cmd/geul/              geul CLI 진입점
internal/cli/          CLI 명령 (cobra, 명령마다 파일 하나)
internal/config/       CLI 설정 (플래그 > GEUL… 환경변수 > 설정 파일 > 데이터 폴더 탐색)
internal/view/page/    뷰어 페이지 (inspect.Document → 페이지 모델·그래프 → HTML, page.html 내장)
internal/view/server/  geul serve의 HTTP 서버 (페이지·바이트·목록, Accept 협상)
internal/view/build/   geul build의 정적 굽기 (server의 페이지·목록을 그대로 쓴다)
internal/view/static/  그래프 JS(vendor/, MIT)를 내장해 /.well-known/geul/_static/으로 제공·굽기
```

- GEUL과 다른 표현 사이의 코덱(`geulnl`, `wikigeul`, `nlgeul`)은 이 라이브러리가 아니라 TANGEUL 문서로 작성한다(geul-dev `draft/codec/`). 이 라이브러리는 코덱 문서가 부르는 호스트 함수와 그 문서를 실행하는 기반을 제공한다.

## 흐름

```
.geul 바이트 → core → Doc ─┬→ schema  (Fact)
                ↑          └→ 코덱 문서(TANGEUL) 실행 → 자연어 등
     resolve: IMPORT-DOC-* → 아카이브 → 외부 Doc 연결

코덱 문서(TANGEUL) 실행 → core 빌더 → 정규형 바이트 → 버전 ID
```

## 원칙

- core는 외부 의존과 네트워크 I/O가 없다. 다른 패키지가 core에 의존하고, 반대 방향은 없다.
- `pkg/*`는 외부 모듈에 의존하지 않는다. cobra는 `cmd/geul`·`internal/cli`만, 그래프 JS는 `internal/view`만 쓴다. 예외 예정: BCP 47 검사용 `golang.org/x/text`(BSD-3).
- 디코딩은 IMPORT를 조회하지 않는다. 조회는 `resolve`의 명시적 단계이고, 위치는 Resolver로 주입한다.
- 쓰기는 빌더로만 한다. 폭과 ID는 빌더가 정한다. 같은 내용은 같은 바이트가 된다.
- 결과는 `doc, warns, err`로 돌려준다. 경고를 실패로 볼지는 호출하는 쪽이 정한다.
- 모르는 레코드는 버리지 않고 원본 WORD로 보존한다.
- 폭과 개수에 고정 상한을 두지 않는다. 상한은 설정으로 바꿀 수 있게 한다.

## `geul` CLI

`.geul` 문서와 버전 이력을 터미널에서 보고 다룬다. git의 명령 체계를 따르되 비교 단위는 줄이 아니라 Edge(LOCAL ID)다.

| 묶음 | 명령 |
|---|---|
| 조회·검사 | `ls`, `log`, `show`, `cat`, `hash`, `diff`, `check [--archive]`, `blame`, `grep`, `resolve` |
| 작성(초안 파일) | `new`, `open`, `add [-i]`, `def`, `set`, `rm [--force]`, `meta [--unset]` |
| 공표 | `publish [--amend] [--name]`, `delta`, `deprecate`, `tombstone --reason`, `name` |
| Issuer·설정 | `issuer new\|ls\|show`(Ed25519, 16 WORD Issuer), `env [-w\|-u]` |
| 뷰어 | `serve [--addr]`(브라우저로 보기), `build -o`(정적 웹 호스트용 굽기) |

```
geul issuer new --name me
geul new me/0x0000000000000001 -o d.geul
geul add d.geul LABEL Q42 '"Douglas Adams"@en'
geul add d.geul TRIPLE P31 Q42 Q5 normal
geul publish d.geul --name adams
geul log me/adams && geul show me/adams
```

- 값 표기는 `pkg/notation` 하나다(`show` 출력 = `add` 입력, 왕복 보장).
- 자리 표지(Reference Codec `0000`, 하위 3비트 = 표지 번호): `ABSENT`(`0x0000`, 적지 않음), `UNKNOWN` UNKNOWN-VALUE(`0x0001`, 알 수 없음), `NONE` NO-VALUE(`0x0002`, 없음). 필드의 CODEC 목록이 그 WORD를 정확히 둘 때만 받는다(와일드카드·CODEC 없는 필드는 받지 않는다). 예: `geul add d.geul TRIPLE P1304 Q30 NONE normal`, `geul add d.geul QUALIFIER P582 '#2060' UNKNOWN`.
- 출처는 근거 Edge `SOURCE`(대상 Edge Id, 원천 Reference·`ABSENT`·`UNKNOWN`)와 그것에 붙는 QUALIFIER다(문법 edge/SOURCE). 예: `geul add d.geul SOURCE '#178' Q5375741`, `geul add d.geul QUALIFIER P813 '#1347' 'GREGORIAN-DATE(2013-12-07)'`. 원천 LOCAL의 성질은 `TRIPLE P1476 '#20' '"…"@en' normal`처럼 쓴다. 문서 전체의 원천은 대상 `#0`으로 쓰고(`geul add d.geul SOURCE '#0' Q2013`), 이 SOURCE와 그것을 대상으로 하는 QUALIFIER는 `add`가 HEAD에 둔다(`show`에서 BODY 위). HEAD·BODY의 Edge Id는 한 번호 공간이다. `schema` E8이 레이아웃·배치를 검사하고, `delta.LocalIDs`가 HEAD·BODY Edge의 번호를, `Node`가 TRIPLE로 기술한 LOCAL 노드의 번호를 이어받는다.
- 데이터 폴더(`.well-known/geul/`을 담은 루트)는 `--data`·`GEULDATA`·상위 폴더 탐색 순으로 찾는다. 개인키는 `$XDG_CONFIG_HOME/geul/keys/`(데이터 폴더 밖).
- `serve`와 `build`는 인자가 없으면 데이터 폴더를 쓴다.

## 뷰어 (`serve`, `build`)

2026-10-27 뷰어 저장소 geul-org/gogeulve(옛 visualeditor, 명령 `ve`·`geulve`)를 이 모듈로 합쳤다.

```
geul serve [--addr host:port] [FILE.geul|DIR]   브라우저로 보기 (기본 127.0.0.1, 빈 포트)
geul build -o OUT [DIR]                           OUT/.well-known/geul/ 에 정적 파일로 굽기
```

- 주소는 공표 주소의 형식 지정(GEUL-ORG)을 따른다. `<path>.geul`은 언제나 바이트(`application/vnd.geul`)다. 확장자 없는 `<path>`는 `Accept`가 `application/vnd.geul`이면 바이트, 아니면(브라우저) 뷰어 페이지다(`Vary: Accept`). 디렉터리는 목록, `<dir>/index.json`은 같은 목록의 JSON(탐색 패널용)이다. 이름 경로(`geul.org/quantity`)는 ID 경로로 넘긴다.
- 디코드·해석은 `pkg/inspect`가 한다. 뷰어는 배치하고 그리기만 한다. JS는 화면 조작만 하고 코덱을 모른다. 외부 네트워크에 접속하지 않는다(CSP `script-src 'self' 'unsafe-inline'`).
- `build` 출력: `<path>.geul`(원본), `<path>`(뷰어 페이지, 확장자 없음, text/html), `<dir>/index.html`·`index.json`, `_static/`. 호스트가 할 일: 확장자 없는 `<path>`에 `Accept: application/vnd.geul`이면 `<path>.geul`을 준다(CloudFront Function), `Vary: Accept`·CSP 헤더. 이름 경로와 최신판 302는 아직 없다(CloudFront KeyValueStore로 옮길 예정). 출력 루트가 원본 트리이거나 그 안이면 거부하고, 디코드되지 않는 문서가 있으면 멈춘다.
- 그래프 뷰는 Cytoscape.js 3.30.2, dagre 0.8.5(graphlib 2.1.8·lodash 모듈 포함), cytoscape-dagre 2.5.0(모두 MIT). npm tarball의 sha512를 대조했고 고지는 `internal/view/static/vendor/LICENSES.md`와 [NOTICE](NOTICE)에 있다.

- 설계·명령 표·미결 사항: geul-dev `draft/README.md`. 실사용 시험·패치 기록: `draft/LOG.md`. 남은 일: `draft/TODO.md`.

## 아카이브

- 로컬 아카이브는 geul.org 웹사이트 저장소의 `artifacts/geul/.well-known/geul/`(geul-dev `draft/geul-org/` 아래)이다. 경로는 GEUL-ORG 공표 주소와 같다.
- 아카이브, 미러, geul.org 덤프는 같은 구조라서 Resolver 하나로 읽는다.

## 현황 (2026-10-27)

- 됨: core 읽기·쓰기 왕복과 정규형·내용 해시, 파일 Resolver와 IMPORT, schema(0층·Codec 규격 레이아웃), 공표·버전 목록·이름·폐기·삭제, diff 저장과 LOCAL ID 이어받기(NEXT-LOCAL-ID, 지운 ID 재사용 거부), 텍스트 표기, `geul` CLI.
- 다음(geul-dev `draft/TODO.md`): 버전 문서 PARENT(v2부터), `pkg/verify`(check·publish 공통 검증, Strict), CLI가 다루면 안 되는 Meta·Edge·아카이브 파일 막기, schema LOCAL 해소 검사(F-4), BCP 47 언어 태그 검사.
- 보류: 서명(SIGN 영역). 검증 클래스 레코드는 보존만 하고 내용 해시 계산에서 뺀다. `repack`, `push`, geul.org 등록 Issuer.

## 개발 규칙

- filefunc: 함수·메서드·타입별로 파일을 나누고 `//ff:` 주석을 단다. 어휘는 `codebook.yaml`. `filefunc validate .` 위반 0건.
- TSMA: `tsma scan .` → `tsma next` → 테스트 → `tsma submit`.
- `go build ./...`, `go vet ./...`, `go test ./...` 통과.
- 공통 테스트 벡터는 다른 언어 구현(`geults` 등)과 함께 쓴다.

## 라이선스

우리 저작물은 CC0 1.0([LICENSE](LICENSE)). 가져다 쓴 라이브러리(cobra 등)와 그래프 JS는 각자의 라이선스를 따르며 고지는 [NOTICE](NOTICE). 명세와 데이터(geul-dev)도 CC0. CC0가 아닌 자료(표준국어대사전 등)의 내용은 코드, 데이터, 테스트에 넣지 않는다.
