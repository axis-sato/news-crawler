# News Crawler

## Getting started

### dev

crawling

```bash
docker-compose up -d
docker-compose exec app go run main.go
```

migration

```bash
./bin/dev.sh migrate! up
```

seeding

```bash
./bin/dev.sh seed
```


### production

crawling

```bash
docker build -t crawler -f docker/go/Dockerfile .
docker run crawler
```