# core

엔진 내장 규칙을 구현한다. 다른 패키지는 모두 core 위에 올라간다.

## 범위

- 프레이밍: TAG WORD와 Length 확장만으로 레코드 경계를 자른다. 스키마가 필요 없다.
- 문서 구조: `DOC-*`, HEAD(Meta), BODY, `END`. 필드 폭은 Codec과 Length가 정한다(문서 단위 폭 선언 없음).
- 모델: Reference, Value, Edge, Meta. 모르는 레코드는 원본 WORD를 그대로 보존한다.
- 읽기: `Decode(bytes) → Doc, warns, err`
- 쓰기: 빌더가 ID와 Length를 정하고 정규형(가장 짧은 폭) 바이트로 쓴다.
- 버전 ID: 검증 클래스 레코드를 뺀 SHA-256(OVERVIEW 계보).

## 구현된 것

| 파일 | 내용 |
|---|---|
| `record.go`, `encode_record.go`, `decode_record.go` | 레코드 하나: TAG WORD, Length(확장 포함, 정규형 검사), ISSUER 확장, Payload |
| `wrap_document.go`, `split_document.go` | DOC 컨테이너: 가장 좁은 DOC-* 로 감싸기, HEAD/BODY/END 나누기와 배치 검사 |
| `*_value.go` | Value: UTF8-TEXT·BCP47(홀짝 쌍), UNSIGNED-INTEGER, SIGNED-/UNSIGNED-DECIMAL |
| `encode_id.go`, `decode_id.go`, `width*.go`, `min_width.go` | 고정 폭(1/2/4/8 WORD) 필드와 폭 코드 (DOC-* 길이, TOULMIN SPAN) |
| `ref_codec.go`, `encode_local_ref.go`, `decode_local_ref.go` | Reference Codec, LOCAL Reference (Edge Id·문서 안 언급) |
| `meta_codec.go`, `value_codec.go`, `codec_issuer.go`, `import_issuer_width.go` | 사용하는 Meta·Value·ISSUER 코드, IMPORT-DOC-* Issuer 폭 |

- toulmin(`draft/t002/toulmin`, 브랜치 `geul-t002`)이 이 패키지를 import해 `.geul`을 읽고 쓴다.
- 아직 없는 것: 지시·지시선언 Reference 인코딩, 버전 ID, 정규형 전체 규칙, 서명.

## 원칙

- 외부 의존과 네트워크 I/O가 없다.
- 디코딩은 IMPORT를 조회하지 않는다. 풀리지 않은 외부 참조는 IMPORT ID와 대상 ID만 가진 채 둔다.
- 쓰기와 읽기를 함께 만들고 `Decode(Encode(x)) == x` 왕복으로 검증한다.

## 보류

- 서명: 검증 클래스 레코드(`SIGNATURE`)는 해석하지 않고 보존만 한다. 버전 ID 계산에서 빼는 규칙만 지킨다.
- 정규형 규칙(폭, Length 확장 조건, 레코드 순서)이 정해지기 전에는 원본 바이트 보존 왕복으로 검증한다.
