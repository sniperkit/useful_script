#!/bin/sh

mkdir -p src/main/java
mkdir -p src/test/java

touch build.gradle
echo "apply plugin : 'java'          "> build.gradle
echo "apply plugin : 'maven'         ">> build.gradle
echo "apply plugin : 'application'   ">> build.gradle
echo "                               ">> build.gradle
echo "repositories {                 ">> build.gradle
echo "    'mavenCentral()'           ">> build.gradle
echo "}                              ">> build.gradle
echo "                               ">> build.gradle
echo "dependencies {                 ">> build.gradle
echo "    compile (                  ">> build.gradle
echo "    )                          ">> build.gradle
echo "    testCompile (              ">> build.gradle
echo "    'junit:junit:4.+'          ">> build.gradle
echo "    )                          ">> build.gradle
echo "}                              ">> build.gradle
echo "                               ">> build.gradle
echo "jar {                          ">> build.gradle
echo "    baseName = 'kkkmmu'        ">> build.gradle
echo "    version = '0.1.0'          ">> build.gradle
echo "}                              ">> build.gradle
echo "                               ">> build.gradle
echo "task wrapper(type: Wrapper) {  ">> build.gradle
echo "    gradleVersion = '3.0'      ">> build.gradle
echo "}                              ">> build.gradle
echo "                               ">> build.gradle
echo "mainClassName = 'hello.HelloWorld' ">> build.gradle

gradle wrapper


