# api-cloud-logging

## Build

```
git clone https://github.com/yuuLab/api-cloud-logging.git
```

```
cd api-cloud-logging
```

```
docker build -t asia-northeast1-docker.pkg.dev/{PROJECTID}/cloudrun-go-app-test/app .
```

```
docker run -it -p 8080:8080 -e "PORT=8080" asia-northeast1-docker.pkg.dev/{PROJECTID}/cloudrun-go-app-test/app
```
