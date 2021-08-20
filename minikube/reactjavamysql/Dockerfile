FROM openjdk:8

RUN mkdir /app

WORKDIR /app

COPY target/springboot-0.0.1-SNAPSHOT.jar /app/springboot.jar

EXPOSE 8090

CMD ["java", "-jar", "/app/springboot.jar"]

