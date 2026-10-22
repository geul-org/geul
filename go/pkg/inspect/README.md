# inspect

GEUL 문서(`.geul`)를 해석된 문서 모델로 읽는다. 문서를 보여 주는 도구(뷰어 ve, 텍스트 표현 geulmd)가 같은 해석을 쓰도록 한 곳에 둔다.

```go
d, err := inspect.Decode(b) // Document{Head, Body, Widths, Profile, Toulmin}
```

| Entry 필드 | 내용 |
|---|---|
| `Record`, `Words` | 원본 레코드와 그 워드(TAG WORD 포함) |
| `Name` | 스펙 이름(`core.CodecName`, `toulmin.RecordName`). 모르면 `""` |
| `Data` | 해석 내용. 아래 표 |
| `ID`, `Refs` | Edge Id와 이 레코드가 참조하는 Edge Id(0 제외, 필드 순서) |

| Data | 레코드 |
|---|---|
| `Value` | HEAD의 SPEC-VERSION·PROFILE·DEFAULT-LANGUAGE, BODY의 독립 Value |
| `Widths`, `DocName`, `ImportDoc` | ID-WIDTH, DOC-NAME, IMPORT-DOC |
| `Source`, `toulmin.TextPacket`, `ControlPacket`, `NodePacket`, `EdgePacket` | PROFILE이 `tangeul-1.x`일 때의 TOULMIN 레코드 |

- **무손실 원칙.** TOULMIN 레코드는 다시 인코딩한 결과가 원래 워드와 같을 때만 `Data`를 채운다. 그 밖에는 `Data`가 nil이고 원본 레코드만 남는다.
- ID-WIDTH는 HEAD의 어느 위치에 있든 먼저 읽는다. SOURCE가 선언한 SPAN 폭은 그 뒤 레코드에 적용한다.
- PREDICATE-DECLARE(조건식, tangl AST)는 해석하지 않는다.
