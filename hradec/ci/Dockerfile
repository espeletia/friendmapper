FROM alpine:3.20
ARG app="hradec"

ADD ./${app}/configurations /app/configurations
ADD ./${app}/build /app
ADD ./${app}/migrations /app/migrations

WORKDIR /app

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
RUN chown -R nonroot:nonroot /app
USER nonroot

EXPOSE 8080

ENTRYPOINT ./start_server


