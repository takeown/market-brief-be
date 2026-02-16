# Repository Guidelines

## 프로젝트 구조 및 모듈 구성
이 저장소는 시장/뉴스 브리핑 API를 제공하는 Go 백엔드입니다.
- `cmd/server/main.go`: 애플리케이션 진입점, 의존성 연결
- `internal/config`: 환경 변수 로딩(`.env` 지원)
- `internal/router`: Gin 라우터 구성 및 라우트 등록
- `internal/news`: 시장/종목 뉴스 핸들러, DTO, 서비스 로직
- `internal/finnhub`: Finnhub 외부 API 클라이언트 및 타입 정의
- `internal/article`: 기사 요약 엔드포인트(현재 더미 서비스)
- `internal/pkg/response`: 공통 API 응답 헬퍼
- `docs/design.md`: 아키텍처 메모 및 로드맵
- `tmp/`: Air 빌드 산출물(생성 파일, 소스 아님)

## 빌드, 테스트, 개발 명령어
- `air`: 핫 리로드 개발 서버 실행(`.air.toml` 사용)
- `go run ./cmd/server`: Air 없이 서버 실행
- `go build -o market-brief-server ./cmd/server`: 배포용 바이너리 빌드
- `go test ./...`: 전체 테스트 실행(기능 추가 시 테스트 함께 작성)
- `go fmt ./...`: 전체 패키지 포맷팅

## 코딩 스타일 및 네이밍 규칙
- Go 표준 포맷(`gofmt`)과 관용적 패키지 구조를 따릅니다.
- 패키지는 도메인 중심으로 분리합니다(`news`, `article`, `finnhub` 등).
- 공개 식별자는 `PascalCase`, 비공개 식별자는 `camelCase`를 사용합니다.
- 핸들러/서비스 이름은 역할이 드러나게 작성합니다. 예: `GetMarketNews`, `NewArticleHandler`
- DTO는 `dto.go`에 작고 명시적인 구조체로 정의합니다.

## 테스트 가이드라인
- 테스트 파일은 코드 인접 위치에 `*_test.go`로 작성합니다. 예: `internal/news/service_test.go`
- 서비스 로직/검증은 테이블 기반 테스트를 우선합니다.
- 외부 API 호출과 HTTP 핸들러의 정상/오류 경로를 모두 검증합니다.
- PR 생성 전 `go test ./...`를 반드시 실행합니다.

## 커밋 및 PR 가이드라인
Git 히스토리 기준 `type: summary` 형식(`feat:`, `docs:`, `refactor:`)을 권장합니다.
- `feat: add company news date range validation`
- `fix: handle empty FINNHUB_API_KEY`

PR에는 아래 내용을 포함합니다.
- 변경 동작 요약
- 영향받는 엔드포인트/파일
- 테스트 근거(`go test ./...` 결과)
- 관련 이슈/티켓 링크(있다면)

## 보안 및 설정 팁
- 실제 시크릿은 `.env`에 커밋하지 말고 로컬 전용 값만 사용하세요.
- 필수 런타임 변수: `PORT`, `FINNHUB_API_KEY` (선택: `FINNHUB_BASE_URL`)
- 핸들러 경계에서 사용자 입력을 검증하고 정제하세요.
