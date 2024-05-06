FROM scratch

COPY ./build/counter-player ./

# Run
CMD [ "/counter-player" ]