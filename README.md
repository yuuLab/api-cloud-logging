# api-cloud-logging

## Local Build

```
git clone https://github.com/yuuLab/api-cloud-logging.git
```

```
cd api-cloud-logging
```

```
docker build -t cloudrunapp .
```

```
docker run -p 8080:8080 -e "PORT=8080" cloudrunapp
```
