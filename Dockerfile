FROM ubuntu
RUN mkdir -p /app
COPY blog /app/blog
CMD /app/blog

EXPOSE 8090