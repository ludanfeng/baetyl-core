FROM --platform=$TARGETPLATFORM golang:1.13.5-stretch as devel
COPY / /go/src/
RUN cd /go/src/ && make all

FROM --platform=$TARGETPLATFORM busybox
COPY --from=devel /go/src/baetyl-core /bin/
COPY /res/*.template /var/lib/baetyl/page/
ENTRYPOINT ["baetyl-core"]