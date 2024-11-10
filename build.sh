 docker image rm -f ghcr.io/astriaorg/go-ethereum:auctioneer-28;
docker build \
  --build-arg COMMIT=$(git rev-parse HEAD) \
  --build-arg VERSION=0.1 \
  --build-arg BUILDNUM=1 \
  --tag ghcr.io/astriaorg/go-ethereum:auctioneer-40 .;