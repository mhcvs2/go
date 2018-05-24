FROM registry.bst-1.cns.bstjpc.com:5000/golang:1.10.2

RUN mkdir /go/src/update-deploy-image

WORKDIR /go/src/update-deploy-image

VOLUME /out /usr/local/bin

COPY . .

RUN go install -v update-deploy-image && cp /go/bin/update-deploy-image /out/update-deploy-image