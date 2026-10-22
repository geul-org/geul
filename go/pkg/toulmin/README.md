# toulmin

TOULMIN 코덱(명세: geul-dev `draft/t002/docs/TOULMIN.md`)을 구현한다. toulmin 엔진의 판정·실행 그래프를 GEUL 레코드로 쓰고 읽는다.

| 묶음 | 내용 |
|---|---|
| 코드표 | NodeType, RoleCode, LevelCode, PropCode, ControlCode, SectionCode, 플래그, 비교·논리·양화 연산자, 필드 타입 |
| 코덱 대응 | 직접 TOULMIN 코덱(`0x60`~`0x76`)과 GEUL ISSUER 형식 번호(`0xC000`~) ↔ 코드표 |
| 값 | VFix(판정값·Qualifier ↔ SIGNED-DECIMAL), Gate(연산자 + UNSIGNED-DECIMAL), Span(선언 폭) |
| 레코드 | NODE, LINK, 제어(SECTION, CASE-BEGIN/END, PROSE), TEXT, SOURCE |
| 필드 도우미 | `PutID`/`TakeID`, `PutValue`/`TakeValue`, `PutSpan`/`TakeSpan`, `TakeWord`, `ExpectEnd`, `EdgeRecord`, `RecordRef` |
| 버전 | `SpecVersion`("geul-t002"), `ProfileString`/`CheckProfile`("tangeul-1.0") |

- core 위에서 동작하며 tangl AST에 의존하지 않는다.
- EXPR(PREDICATE-DECLARE 조건식)과 AST ↔ 코드 변환은 tangl AST에 묶여 있어 toulmin(`pkg/tangeul`)에 남긴다.
