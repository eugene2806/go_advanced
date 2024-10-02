FROM alpine:latest AS base

RUN apk add git openjdk17

RUN git clone https://github.com/dmdev2020/spring-starter.git
WORKDIR spring-starter
RUN git checkout lesson-125 && ./gradlew bootJar

FROM alpine:latest AS result

RUN apk add openjdk17

WORKDIR /app
COPY --from=base /spring-starter/build/libs/spring-starter-*.jar ./service.jar

COPY application-dev.yaml .

EXPOSE 8080
ENTRYPOINT ["java", "-jar", "service.jar"]
CMD ["--spring.config.location=classpath:/application.yml,file:application-dev.yaml"]


