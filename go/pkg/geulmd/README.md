# geulmd

GEUL 문서(`.geul`)를 GEULMD로 바꾼다. GEULMD는 GEUL의 텍스트 표현이며 표준 Markdown이다. 미디어 타입은 `text/vnd.geulmd`, 확장자는 `.md`다.

## 형식

````markdown
```geul
SPEC-VERSION "geul-t002"
PROFILE "tangeul-1.0"
ID-WIDTH edge=2 var=1 import=1
DEFAULT-LANGUAGE lang"ko"
DOC-NAME 0x0001 0x0000000000000005
IMPORT-DOC #1 0x0001 0x0000000000000002
```

```geul
SOURCE #40 span=1 "# 문서 …"
TEXT #44 "nlgeul"
NAMESPACE #1 role=0 level=0 @1668+32
LABEL #45 #1 -> #44 @1668+32
ATTACKS #80 #9 -> #7 @2100+40
DO-ONCE #81 #7 -> #12/0x0000 gate=at-least/75 @2140+50
VALUE uint(5)
RECORD EDGE 0x61 : 0x0067 0x0001 …
```
````

- **첫 펜스 블록(info string `geul`)이 HEAD다.** 이 블록이 있으면 GEULMD이고, 없으면 자연어 Markdown이다.
- **둘째 펜스 블록이 BODY다.** 레코드 하나가 한 줄이며 줄 순서가 레코드 순서다.
- DOC-*, BODY, END 구조 레코드는 쓰지 않는다. 길이와 위치에서 다시 계산한다.
- **무손실 원칙.** 구조 표기는 다시 인코딩한 결과가 원래 워드와 같을 때만 쓴다. 그 밖에는 `RECORD <kind> 0x<codec> [<issuer> 0x<format>] : <워드…>`로 쓴다.

| 표기 | 뜻 |
|---|---|
| `#n` | Edge Id, IMPORT ID |
| `@o+l` | SOURCE 안의 바이트 구간 |
| `0x…` | ID. 자릿수 = 폭 × 4 |
| `"…"` | UTF8-TEXT (Go 따옴표 규칙) |
| `lang"ko"` | LANGUAGE-BCP47-CODE |
| `uint(n)`, `dec(CeE)`, `udec(CeE)` | UNSIGNED-INTEGER, SIGNED-/UNSIGNED-DECIMAL (계수·지수 보존) |

- HEAD: SPEC-VERSION, PROFILE, DEFAULT-LANGUAGE, ID-WIDTH, DOC-NAME, IMPORT-DOC은 타입 표기, 나머지 Meta는 raw다.
- BODY: PROFILE이 `tangeul-*`이면 TOULMIN 레코드(SOURCE, TEXT, SECTION, CASE-BEGIN/END, PROSE, NODE, LINK)를 구조로 쓴다. PREDICATE-DECLARE의 조건식은 tangl 식 AST에 속하므로 raw로 쓴다.

## 미결정

- GEULMD → GEUL(역방향) 파서.
- TANGEUL 표면(프로필 렌더러)을 BODY 블록 대신 쓰는 방식.
- SOURCE 원문을 따옴표 문자열 대신 읽기 쉬운 블록으로 쓰는 방식.
