# nfs-check
Test if RWX or NFS file system ready before start

# build

```shell
export nfscheckversion=0.0.x

docker build -t dataplane/nfscheck:$nfscheckversion -f Dockerfile.alpine .
docker tag dataplane/nfscheck:$nfscheckversion dataplane/nfscheck:$nfscheckversion
docker push dataplane/nfscheck:$nfscheckversion
docker tag dataplane/nfscheck:$nfscheckversion dataplane/nfscheck:latest
docker push dataplane/nfscheck:latest

docker run --rm dataplane/nfscheck:$nfscheckversion
```
