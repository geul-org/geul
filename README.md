# GEUL

General Encoding Unified Language — 인공지능과 시스템을 위한 심볼릭 언어의 라이브러리.
GEUL 문서(`.geul`, `application/vnd.geul`)를 읽고 쓰고, 다른 표현과 서로 바꾼다.

| 디렉터리 | 언어 | 상태 |
|---|---|---|
| `go/` | Go | `pkg/core`·`toulmin`·`delta`·`inspect`·`resolve`·`schema` ([go/README.md](go/README.md)), `cmd/geul` CLI 계획(아래) |
| `ts/` | TypeScript | 예정 |

- 명세: [geul-org/geul-dev](https://github.com/geul-org/geul-dev) `draft/t002/docs/`
- 언어 구현은 공통 테스트 벡터로 서로를 검증한다.
- 라이선스: CC0.

## 계획: `geul` CLI (2026-10-27)

`.geul` 문서와 버전 이력을 터미널에서 보고 다루는 도구. git의 `log`·`show`·`diff`에 해당한다. 지금은 ve(웹 뷰어)나 일회용 도구로만 볼 수 있다.

### 위치와 의존

- `go/cmd/geul`(진입점) + `go/internal/cli`(명령 정의). 같은 모듈 `github.com/geul-org/geul/go`.
- 명령 라이브러리: [cobra](https://github.com/spf13/cobra) v1.10.2(Apache-2.0). filefunc·tsma와 같은 판·같은 배치(`internal/cli`에 명령마다 파일 하나, `//ff:func feature=cli type=command`).
- cobra는 `cmd`·`internal/cli`만 쓴다. `pkg/*`는 외부 의존 없음을 유지한다(라이브러리를 import하는 쪽은 cobra를 빌드하지 않는다).
- 로직은 `pkg/*`에 두고 CLI는 인자 해석과 출력만 한다. CLI에만 있는 기능은 만들지 않는다.

### git과의 관계

- git의 저장(내용 주소 객체, 이진 delta 압축)과 이력 원리는 GEUL과 같다: 해시 경로 `h/`, PARENT, 스냅숏 기본·diff는 저장 형식.
- git의 비교 단위는 줄이라 `.geul`(16비트 WORD 이진)에는 맞지 않는다. `geul`은 같은 명령 체계를 쓰되 비교 단위를 **Edge(LOCAL ID)** 로 바꾼다. LOCAL ID는 버전을 넘어 유지되므로(OVERVIEW 개정) 같은 서술을 추측 없이 짝짓는다.

### 1단계 명령 (조회·검사)

| 명령 | git | 하는 일 | 라이브러리 |
|---|---|---|---|
| `geul log <문서>` | `log` | 버전 목록: 버전, 해시, 공표 시각, 폐기·삭제, NEXT-LOCAL-ID | resolve |
| `geul show <문서>[@버전] \| <파일>` | `show` | 레코드를 한 줄씩 텍스트로(이름, LOCAL ID, 값). diff로 저장된 버전은 복원해서 | resolve, inspect |
| `geul diff <문서>@v1 [@v2] \| <파일> <파일>` | `diff` | LOCAL ID 단위 변경: `+ #40 …`, `- #7 …`, `~ #5 …` | delta, inspect |
| `geul check <파일\|디렉터리>…` | `fsck` | schema 검사 Fact 출력, Fact가 있으면 종료 코드 1. `--archive`면 버전 목록·해시 경로·diff 사슬까지 | schema, resolve |
| `geul cat <문서>[@버전]` | `cat-file` | 스냅숏 바이트 출력(파이프용, diff는 복원) | resolve |
| `geul hash <파일>` | `hash-object` | 내용 해시(버전 목록에 적히는 SHA-256) | core |
| `geul ls [<Issuer>]` | `ls-tree` | Issuer·문서 목록과 이름(documents.geul, issuers.geul) | resolve |

### 2단계 명령 (이력 활용)

| 명령 | git | 하는 일 | 라이브러리 |
|---|---|---|---|
| `geul blame <문서>[@버전]` | `blame` | 서술(LOCAL ID)마다 처음 나타난 버전과 마지막으로 바뀐 버전 | resolve, delta |
| `geul grep <패턴> [<문서>…]` | `grep` | 텍스트 Value·Reference·코덱으로 레코드 찾기 | inspect |
| `geul resolve <문서>` | — | IMPORT 의존성을 따라가 해석 결과·누락 출력 | resolve |

### 3단계 명령 (쓰기, 미결정 해소 뒤)

| 명령 | git | 하는 일 | 라이브러리 |
|---|---|---|---|
| `geul publish <파일>` | `commit` | 아카이브에 새 버전 공표(버전 목록·NEXT-LOCAL-ID·`h/` 갱신) | resolve |
| `geul publish --amend <파일>` | `commit --amend` | 마지막 버전을 덮어씀(`Dir.Overwrite`). 초기 개발 기간 전용 | resolve |
| `geul delta <문서>@v1 @v2 -o <파일>` | — | DELTA 문서 쓰기(저장용 diff) | delta |
| `geul repack <문서>` | `gc`·`repack` | 스냅숏 간격 정책에 따라 버전 파일을 스냅숏·diff로 다시 저장. 해시는 바뀌지 않는다 | delta, resolve |
| `geul deprecate`·`geul tombstone <문서>@버전` | — | DEPRECATION·TOMBSTONE을 버전 목록에 추가 | resolve |

- 용어는 **publish(공표)** 하나로 한다. 명세의 PUBLICATION Edge, 라이브러리 `Dir.Publish`와 같은 말이다.
  - commit이 아닌 이유: git commit은 고칠 수 있는 개인 기록이지만, 공표한 버전은 바뀌지 않는다(해시가 목록에 고정, LOCAL ID 재사용 없음, 철회도 TOMBSTONE 추가로만).
  - release가 아닌 이유: 여러 버전 중 하나를 골라 내놓는 중간층이 GEUL에는 없다. 모든 버전이 공개본이다. 골라 이름 붙이는 일은 태그에 가깝다.
  - 공개 전에 고칠 단계가 필요하면 새 단어 대신 `--amend`를 쓴다.
- 공통: `<문서>`는 이름(`geul.org/quantity`, `wikidata.geul.org/q42`) 또는 ID(`0x0000/0x0806`). 이름 해석은 resolve의 documents.geul.
- 공통 플래그: `--data <디렉터리>`, `--format text|json`(아래 환경 설정).
- 값 표기(날짜, QUANTITY, 언어 텍스트 등)는 ve와 같아야 한다. ve의 표기 코드를 `pkg/inspect`로 옮겨 둘이 함께 쓰는 것을 먼저 검토한다.

### 작성 명령 (문서 생성·Edge 편집)

- 아카이브를 직접 고치지 않는다. 작업 파일(초안)을 편집하고 `publish`로 낸다(git 작업 트리처럼).
- 입력 표기는 `show` 출력 표기와 같다(왕복).
- Edge 형식(부분 수·종류)은 geul.org 규격 문서로 검사하고, 위반이면 쓰지 않는다(schema).
- Edge ID는 사람이 정하지 않는다. 라이브러리가 NEXT-LOCAL-ID를 이어받아 배정한다(`delta.LocalIDs`).

| 명령 | git | 하는 일 | 라이브러리 |
|---|---|---|---|
| `geul new <Issuer>/<문서> -o <파일> [--name <이름>]` | `init` | DOC-ID, SPEC-VERSION, 빈 BODY로 새 문서. `--name`이면 documents.geul에 이름 등록 | core, resolve |
| `geul open <문서> -o <파일>` | `checkout` | 현재 버전 스냅숏을 초안으로(다음 버전 편집) | resolve |
| `geul add <파일> <EDGE> <부분>…` | — | BODY 끝에 Edge 추가, 배정된 `#n` 출력 | delta, schema |
| `geul add <파일> -i <목록>` | — | 한 줄에 Edge 하나씩 일괄 추가(`-`는 표준 입력) | delta, schema |
| `geul def <파일> <객체>` | — | REFERENCE-NODE로 새 LOCAL 노드 `#n` | delta |
| `geul set <파일> #n <부분>…` | — | 같은 ID로 내용 교체(수정) | schema |
| `geul rm <파일> #n…` | `rm` | Edge 삭제(번호는 다시 쓰지 않음) | delta |
| `geul meta <파일> <이름> <값>` | — | HEAD Meta 쓰기 | core |

예:

```
geul new 0x00010000/0x510000000000002a -o q42.geul
geul add q42.geul LABEL Q42 "Douglas Adams"@en        # → #1
geul add q42.geul TRIPLE P31 Q42 Q5 normal              # → #2
geul add q42.geul QUALIFIER P580 #2 GREGORIAN-DATE(1952-03-11)
geul publish q42.geul
```

부분 표기(안):

- Reference: `Q42`, `P31`(위키데이터), `#12`(LOCAL), `geul.org/quantity`, `geul.org/quantity#3`(외부 문서·LOCAL)
- Value: `"text"`, `"Douglas Adams"@en`, `<https://…>`(URI), `1.20`(SIGNED-DECIMAL, 자릿수 보존), `QUANTITY(652, Q828224, 651, 653)`, `EDTF("1971~")`. 그 밖은 `코덱이름(부분, …)`
- 없음 표지 `_`, 낱말 필드(rank 등)는 규격 문서의 이름(`normal`)

결정할 것:

1. **텍스트 표기 문법.** 사실상 GEUL의 텍스트 직렬화라 명세 문서가 필요하다(왕복 보장, 모호성 규칙: `1.20`과 날짜·정수 구별 등).
2. **수정의 정체성.** 같은 서술을 고친 것(같은 ID)인지 지우고 새로 쓴 것(새 ID)인지는 작성자만 안다. CLI는 `set`과 `rm`+`add`를 모두 둔다. wikiport는 구별할 근거가 없어 늘 삭제+신규다.
3. **초안 파일 대 아카이브 직접 편집.** 초안 파일을 권한다. 초안은 버전 목록에 없으므로 NEXT-LOCAL-ID는 add 때 아카이브 목록과 초안의 최대 LOCAL ID에서 계산한다.

### 환경 설정

우선순위: **플래그 > 환경변수 > 설정 파일 > 자동 탐색·기본값**. git(`GIT_DIR`, `.git` 상위 탐색, `~/.gitconfig`)과 같은 구조다.

| 환경변수 | 플래그 | 뜻 | 기본값 |
|---|---|---|---|
| `GEUL_DATA` | `--data` | 데이터 폴더: `.well-known/geul/`을 담은 아카이브 루트(`resolve.Dir.Root`) | 현재 폴더부터 상위로 `.well-known/geul/`을 찾음. 없으면 설정 파일의 `data` |
| `GEUL_CONFIG` | `--config` | 설정 파일 경로 | `$XDG_CONFIG_HOME/geul/config`(없으면 `~/.config/geul/config`) |
| `GEUL_CACHE` | — | 외부 Issuer 문서(IMPORT 해석, HTTP Resolver) 캐시 | `$XDG_CACHE_HOME/geul`(없으면 `~/.cache/geul`) |
| `GEUL_ISSUER` | `--issuer` | `new`의 기본 Issuer | 없음(필수 지정) |
| `GEUL_REMOTE` | `--remote` | `push`·원격 조회 주소 | `https://geul.org` |
| `GEUL_FORMAT` | `--format` | 출력 형식 `text`·`json` | `text` |

- 설정 파일은 같은 키를 소문자로 둔다(`data`, `cache`, `issuer`, `remote`, `format`). 형식은 한 줄에 `키 = 값`. 설정 파일은 2단계로 미루고, 1단계는 플래그·환경변수·자동 탐색만 구현한다.
- `geul env`(git `config --list`에 해당): 각 값과 그 출처(플래그·환경변수·설정 파일·탐색·기본값)를 출력한다.
- 데이터 폴더를 찾지 못하면 쓰기 명령(`publish`, `name` 등)은 실패한다. 조회 명령은 파일 경로 인자만으로도 동작한다.
- 비밀값(Issuer 서명 키, 원격 인증 토큰)은 설정 파일·환경변수에 평문으로 두지 않는다. `push`·SIGN 설계 때 정한다(OS 키체인, 키 파일 경로만 설정).
- 설정 해석은 `internal/config` 한 곳에 둔다. 외부 의존 없이(`os.Getenv` + 작은 파서) 구현하고 viper는 쓰지 않는다. `pkg/*`는 환경변수를 읽지 않는다(값은 호출하는 쪽이 넘긴다).

### Issuer 명령

- 자체 발급은 16 WORD Issuer만 가능하다(공개키에서 계산, docs OVERVIEW 16 WORD Issuer 생성 규칙). 지금 CLI는 16 WORD만 다룬다. geul.org 등록 Issuer(1·2·4 WORD)는 차후 과제(아래)다.

| 명령 | 하는 일 | 라이브러리 |
|---|---|---|
| `geul issuer new [--alg ed25519] [--name <이름>]` | 키 쌍 생성, 16 WORD Issuer ID 계산, `ISSUER-KEY`를 담은 Issuer 문서 공표 | core, resolve |
| `geul issuer ls` | 로컬에서 쓸 수 있는 Issuer와 키 목록 | — |
| `geul issuer show <Issuer>` | Issuer ID, 서명 알고리즘, 공개키, 로컬 이름 | core |

- 개인키는 데이터 폴더 밖에 둔다: `$XDG_CONFIG_HOME/geul/keys/`(권한 0600) 또는 OS 키체인. 데이터 폴더는 공개·git 추적 대상이다.
- 16 WORD Issuer의 이름은 로컬 설정에만 둔다(`issuer.<이름> = 0x…`). `issuers.geul`은 geul.org가 관리하는 등록부라 자체 발급 Issuer를 넣지 않는다.
- 결정할 것: **Issuer 이전.** 16 WORD로 낸 문서는 나중에 2 WORD를 받아도 ID가 다르다. 다시 공표할지, sameAs로 이을지. 권고: 공표 전에 Issuer를 확정한다.

### 자동 관리 파일

- 각 Issuer의 `documents.geul`(문서 이름 목록)과 각 문서의 `versions.geul`(버전 목록)은 `geul` CLI가 자동으로 관리한다. 사람이 직접 편집하지 않는다.
- `versions.geul`: `publish`·`publish --amend`·`deprecate`·`tombstone`이 갱신한다(PUBLICATION·DEPRECATION·TOMBSTONE, PARENT, RECORDED-AT, NEXT-LOCAL-ID, `h/` 사본). 지금 `resolve.Dir.Publish`·`Overwrite`가 하는 일이다.
- `documents.geul`: `new --name <이름>`이 등록하고, `geul name <문서> <이름>`이 붙이거나 바꾼다. 이름은 Issuer 안에서 하나다. 라이브러리에 쓰기 함수가 아직 없어(지금은 일회용 도구로 썼다) `resolve`에 추가한다.
- 두 파일은 버전 관리를 하지 않는 문서라 diff 대상이 아니다(OVERVIEW 계보 diff).

| 명령 | git | 하는 일 | 라이브러리 |
|---|---|---|---|
| `geul name <문서> <이름>` | `tag`에 가까움 | documents.geul에 이름 등록·변경 | resolve(추가 예정) |

### git 연동

- `.gitattributes`에 `*.geul diff=geul`, `git config diff.geul.textconv "geul show"`를 두면 geul-dev에서 `git diff`·`git log -p`가 `.geul`을 텍스트로 보여 준다. 줄 단위 비교이므로 보기용이고, 정확한 비교는 `geul diff`.

### git에서 가져오지 않는 것

- `add`(스테이징), `restore`·`reset`·`stash`(작업 트리), `branch`·`merge`·`rebase`·`cherry-pick`, `remote`·`fetch`·`pull`, 서명. 작업 공간 개념이거나 WMS(GEUL DBMS)·SIGN 영역에서 다룬다. `push`는 차후 과제(아래).

### 차후 과제

- **`geul push <문서>|<Issuer>`** (git `push`): 로컬 아카이브에서 공표한 버전을 geul.org(또는 Issuer의 공표 주소)에 올린다. 버전 파일, `h/` 사본, `versions.geul`, `documents.geul`을 함께 보낸다.
  - 선행: Issuer 소유 인증, 버전 목록 서명(SIGN 영역), 서버 쪽 수신·검증 API(해시·버전 목록 연속성·NEXT-LOCAL-ID 비감소 검사).
  - 이미 올라간 버전은 바꿀 수 없다. 원격 목록이 로컬보다 앞서 있으면 거부한다(fast-forward만).

- **geul.org 등록 Issuer.** `geul issuer register [--domain <도메인>]`: 16 WORD 키로 서명한 신청서를 geul.org에 보내 2 WORD(도메인 소유 증명)나 4 WORD(증명 없음)를 받는다. 1 WORD는 geul.org가 배정하므로 명령이 없다.
  - 등록할 때 16 WORD 키를 그대로 등록 Issuer의 서명 키로 쓴다(키를 따로 만들지 않는다).
  - 선행: geul.org 등록 API, 도메인 소유 증명 방식(DNS TXT·HTTP `.well-known`), 검증키 공개 문서, 키 분실 복구(docs ISSUER 미결정).

### 먼저 정할 것

1. diff 순서 제약: 새 서술이 기준 문서 중간에 들어가면 diff를 만들 수 없다(wikiport는 레이블→설명→별칭→클레임 순서라 대부분 해당). 끝에 붙이는 순서를 허용할지, diff에 위치를 적을지.
2. 스냅숏 간격과 diff 방향(최신을 스냅숏, 과거를 diff로 둘지).
3. diff 파일의 저장 위치와 `h/`(해시 경로) 사본의 관계.
4. **명세 모순: Issuer 폭.** docs ISSUER.md는 "폭은 식별자의 일부가 아니다(`0x0001` = `0x00000001`), 정규형은 가장 짧은 폭"이고, docs OVERVIEW Issuer 표 아래는 "폭은 식별자의 일부다(`0x0005` ≠ `0x0000 0x0005`)"다. 정규형 규칙(schema F5·R4·R5)으로 보아 ISSUER.md가 현행이니 OVERVIEW를 고친다.

### 완료 조건

- filefunc 위반 0건, TSMA 새 세션 TODO 0, `go build`·`go vet`·`go test` 통과.
- 1단계: `.well-known/geul/` 아카이브 전체에 대해 `geul check --archive` Fact 0건, 규격 문서 몇 개로 `log`·`show`·`diff`·`cat`·`hash`·`ls` 출력 확인.
