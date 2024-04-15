# Docker Images
Have a look, fill ya boots.

## Developing
Build it with docker and scan it with trivy it before you push it

```shell
docker build -t local/[IMAGE]:v0.0.0 -f [IMAGE]/Dockerfile .
trivy image local/[IMAGE]:v0.0.0
```


## Releasing
Just tag it and watch it go!
