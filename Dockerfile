FROM debian

WORKDIR /app

COPY ./build/counter-player .

# Run
ENTRYPOINT [ "/app/counter-player" ]