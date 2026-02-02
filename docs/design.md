# Market Brief 설계 문서

뉴스 기반 주가 영향(호재/악재) 분석 백엔드

## LLM 연동 방식 옵션

| 방식 | 장점 | 단점 |
|------|------|------|
| **OpenAI API** | 안정적, 문서 풍부 | 유료, 레이턴시 |
| **Anthropic API** | 긴 컨텍스트, 분석 강점 | 유료 |
| **Ollama (로컬)** | 무료, 프라이버시 | 성능 제한, 서버 리소스 필요 |
| **Google Gemini** | 무료 티어 있음 | API 제한 |

## 구현 우선순위

```
1. 뉴스 수집 (입력 없으면 분석 불가)
   └─ 크롤링 or 뉴스 API 연동

2. LLM 연동 (핵심 기능)
   └─ 호재/악재 판단 프롬프트 설계

3. DB 연동 (결과 저장)
   └─ PostgreSQL or MongoDB

4. 종목 관리 (부가 기능)
   └─ 티커 매핑, 관심 종목 등
```

## 뉴스 수집 (Finnhub 연동)

### 개요
Finnhub API를 통해 실시간 시장 뉴스 및 종목별 뉴스를 수집합니다.

### API 엔드포인트

| 엔드포인트 | 설명 | 파라미터 |
|-----------|------|----------|
| `GET /news/market` | 시장 전체 뉴스 조회 | `category` (기본값: general) |
| `GET /news/company/:symbol` | 특정 종목 뉴스 조회 | 티커 심볼 (예: AAPL, TSLA) |

### 환경 변수

```env
FINNHUB_API_KEY=<your-finnhub-api-key>
FINNHUB_BASE_URL=https://finnhub.io/api/v1  # 선택사항
```

### 아키텍처

```
cmd/server/main.go
    └── internal/
        ├── config/config.go      # 환경 변수 로드 (.env 지원)
        ├── finnhub/
        │   ├── client.go         # HTTP 클라이언트 (10초 타임아웃)
        │   └── types.go          # MarketNews, CompanyNews 타입
        ├── news/
        │   ├── handler.go        # Gin 핸들러
        │   ├── service.go        # 비즈니스 로직
        │   └── dto.go            # 응답 DTO
        └── router/router.go      # 라우트 등록
```

### 응답 예시

```json
{
  "news": [
    {
      "id": 123456,
      "headline": "Apple Reports Record Revenue",
      "summary": "Apple Inc. announced...",
      "source": "Reuters",
      "url": "https://...",
      "image": "https://...",
      "datetime": 1706889600,
      "category": "technology"
    }
  ]
}
```

---

## 결정 필요 사항

- [x] 뉴스 소스: ~~직접 크롤링 vs 뉴스 API~~ → **Finnhub API 사용**
- [ ] LLM 선택: OpenAI / Anthropic / Ollama / Gemini
- [ ] 타겟 시장: 한국 주식 / 미국 주식 / 둘 다
