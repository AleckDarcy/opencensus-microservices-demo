# Prepare

## Versions
Java: 1.8

## Initialize Gradle cache for Java
If you already have gradle caches like *~/.gradle/caches/modules-2/files-2.1*, skip this step.

Inside folder *opencensus-microservices-demo/src/adservice*, execute:
```shell
./gradlew installDist
```

## Generate GRPC for Java
Clone code from *https://github.com/AleckDarcy/reload.git*, branch: feature/multi_threads.
Replace **.jar** files under **~/.gradle/caches** (~/.gradle/caches/modules-2/files-2.1/io.grpc/grpc-core/1.12.0/541a5c68ce85c03190e29bc9e0ec611d2b75ff24/grpc-core-1.12.0.jar & ~/.gradle/caches/modules-2/files-2.1/io.grpc/grpc-stub/1.12.0/fbd2bafe09a89442ab3d7a8d8b3e8bafbd59b4e0/grpc-stub-1.12.0.jar) by .jar files under *reload/java/grpc*.

# Run
Execute:
```shell
skaffold dev
```