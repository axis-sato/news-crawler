# News Crawler

![Deploy Crawler](https://github.com/c8112002/news-crawler/workflows/Deploy%20Crawler/badge.svg)

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
./bin/deploy.sh
```