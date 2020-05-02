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
./bin/dev.sh migrate! down
```

seeding

```bash
./bin/dev.sh seed
```

reset

```bash
./bin/dev.sh reset
```


### production

crawling

```bash
docker build -t crawler -f docker/go/Dockerfile-crawler .
docker run crawler
```