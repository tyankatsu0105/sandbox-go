FROM golang:1.15 as builder
# =======================
# ARG
ARG WORKDIR
ARG APP_NAME
# =======================

ADD . ${WORKDIR}
WORKDIR ${WORKDIR}
RUN go build -o ${APP_NAME}


FROM alpine:latest
# =======================
# ARG
ARG WORKDIR
ARG APP_NAME
# =======================

RUN apk --no-cache add ca-certificates
COPY --from=builder ${WORKDIR}/${APP_NAME} .
CMD ["./${APP_NAME}"]