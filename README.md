# Market Brief BE

뉴스 기사 요약 API 서버

## 요구사항

- Go 1.21+
- [Air](https://github.com/air-verse/air) (개발용 핫 리로드)

## Development (Air)

```bash
# Air 설치 (최초 1회)
go install github.com/air-verse/air@latest

# 개발 서버 실행 (핫 리로드)
air
```

## Build & Run

```bash
# 빌드
go build -o market-brief-server ./cmd/server

# 실행
./market-brief-server
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | 서버 상태 확인 |
| POST | `/articles/summarize` | 기사 요약 |

### POST /articles/summarize

```bash
curl -X POST http://localhost:8080/articles/summarize \
  -H "Content-Type: application/json" \
  -d '{"url":"https://news.example.com/article"}'
```

