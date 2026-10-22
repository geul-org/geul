# gsongeul

원형 GSON을 GEUL 문서로 되돌린다. `encode/geulgson` 원형 보기의 역함수다.

- 받는 것은 원형 보기뿐이다. 해석 보기와 임의의 JSON은 받지 않는다.
- 결과는 core 빌더로 다시 쓴다. 손으로 쓴 GSON도 정규형 바이트가 된다.
- 왕복 `gsongeul(geulgson(x)) == x`로 검증한다.
