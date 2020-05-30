# News Crawler

## Getting started

### dev

crawling

```bash
docker-compose up -d
docker-compose exec app go run main.go
```

### production

crawling

```bash
docker build -t crawler -f docker/go/Dockerfile-crawler .
docker run crawler
```