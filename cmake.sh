#!/bin/sh

if [ $# -ne 1 ]; then
    echo "Usuage: gnew.sh ProjectName"
    exit -1;
fi

project=$1;
package=`echo $project | tr '[:upper:]' '[:lower:]'`

mkdir $1
cd $1

# Create Project class
mkdir -p src
mkdir -p build
mkdir -p lib
mkdir -p app
mkdir -p test

touch CMakeLists.txt
echo "CMAKE_MINIMUM_REQUIRED(VERSION 2.8)"                                      > CMakeLists.txt
echo "PROJECT($1)"                                                              >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(src .src)"                                               >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(app .app)"                                               >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(lib .lib)"                                               >> CMakeLists.txt
echo "SET(CMAKE_VERBOSE_MAKEFILE on)"                                           >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(test .test)"                                             >> CMakeLists.txt
echo "SET(CFLAGS \"-Wall -Werror\")"                                            >> CMakeLists.txt

touch lib/CMakeLists.txt
echo "ADD_SUBDIRECTORY(unity)"                                                  > lib/CMakeLists.txt
# echo "ADD_SUBDIRECTORY(mysql)"                                                  >> lib/CMakeLists.txt
echo "ADD_SUBDIRECTORY(libds)"                                                  >> lib/CMakeLists.txt
echo "ADD_SUBDIRECTORY(collection)"                                             >> lib/CMakeLists.txt
# echo "ADD_SUBDIRECTORY(onion)"                                                  >> lib/CMakeLists.txt
echo "ADD_SUBDIRECTORY(llog)"                                                   >> lib/CMakeLists.txt
echo "SET(CMAKE_VERBOSE_MAKEFILE on)"                                           >> lib/CMakeLists.txt

touch src/$1.c
echo "#include <stdio.h>"                                                       > src/$1.c
echo "void $1(void)"                                                            >> src/$1.c
echo "{"                                                                        >> src/$1.c
echo "      printf(\"Hello world!\\\n\");"                                      >> src/$1.c
echo "}"                                                                        >> src/$1.c
touch src/$1.h
echo "#ifndef __$1_H__"                                                         > src/$1.h
echo "#define __$1_H__"                                                         >> src/$1.h
echo "void $1(void);"                                                           >> src/$1.h
echo "#endif /* __$1_H__ */"                                                    >> src/$1.h
touch src/CMakeLists.txt

echo "AUX_SOURCE_DIRECTORY(\${CMAKE_CURRENT_SOURCE_DIR} LIB_SRC)"               > src/CMakeLists.txt
echo "ADD_LIBRARY(kkkmmu SHARED \${LIB_SRC})"                                   >> src/CMakeLists.txt
echo "ADD_LIBRARY(kkkmmu_static STATIC \${LIB_SRC})"                            >> src/CMakeLists.txt
echo "SET(LIBRARY_OUTPUT_PATH \${PROJECT_BINARY_DIR}/lib)"                      >> src/CMakeLists.txt
echo "SET_TARGET_PROPERTIES (kkkmmu_static PROPERTIES OUTPUT_NAME kkkmmu)"      >> src/CMakeLists.txt
echo "SET_TARGET_PROPERTIES (kkkmmu PROPERTIES CLEAN_DIRECT_OUTPUT 1)"          >> src/CMakeLists.txt
echo "SET_TARGET_PROPERTIES (kkkmmu PROPERTIES CLEAN_DIRECT_OUTPUT 1)"          >> src/CMakeLists.txt
echo "SET_TARGET_PROPERTIES (kkkmmu PROPERTIES VERSION 1.1 SOVERSION 1)"        >> src/CMakeLists.txt
echo "# install (TARGETS $1 $1_static "                                         >> src/CMakeLists.txt
echo "#     LIBRARY DESTINATION \${PROJECT_BINARY_DIR}/lib"                     >> src/CMakeLists.txt
echo "#     ARCHIVE DESTINATION \${PROJECT_BINARY_DIR}/lib)"                    >> src/CMakeLists.txt
echo "INSTALL (FILES $1.h DESTINATION \${PROJECT_BINARY_DIR}/include/$1)"       >> src/CMakeLists.txt

touch app/$1Main.c
echo "#include \"$1.h\""                                                        > app/$1Main.c
echo "int main(int argc, char** argv)"                                          >> app/$1Main.c
echo "{"                                                                        >> app/$1Main.c
echo "      $1();"                                                              >> app/$1Main.c
echo "}"                                                                        >> app/$1Main.c

touch app/CMakeLists.txt
echo "AUX_SOURCE_DIRECTORY(\${CMAKE_CURRENT_SOURCE_DIR} APP_SRC)"               > app/CMakeLists.txt
echo "ADD_EXECUTABLE($1 \${APP_SRC})"                                           >> app/CMakeLists.txt
echo "INCLUDE_DIRECTORIES(\${PROJECT_SOURCE_DIR}/src \${CMAKE_CURRENT_SOURCE_DIR})"                          >> app/CMakeLists.txt
echo "SET(EXECUTABLE_OUTPUT_PATH \${PROJECT_BINARY_DIR}/bin)"                   >> app/CMakeLists.txt
echo "LINK_DIRECTORIES (\${PROJECT_BINARY_DIR}/lib)"                            >> app/CMakeLists.txt
echo "TARGET_LINK_LIBRARIES($1 kkkmmu)"                                         >> app/CMakeLists.txt

touch test/$1Test.c
echo "#include \"$1.h\""                                                        > test/$1Test.c
echo "#include <llog.h>"                                                        >> test/$1Test.c
echo "#include <unity.h>"                                                       >> test/$1Test.c
echo "#include \"unity_internals.h\""                                           >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "void test_failure_func(void)"                                             >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      TEST_ASSERT_EQUAL_STRING(\"ello\", \"Hello\");"                     >> test/$1Test.c
echo "}"                                                                        >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "void test_success_func(void)"                                             >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      TEST_ASSERT_EQUAL_STRING(\"Hello\", \"Hello\");"                    >> test/$1Test.c
echo "}"                                                                        >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "void test_ignore_func(void)"                                              >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      TEST_IGNORE_MESSAGE(\"This is a ignored test!\");"                  >> test/$1Test.c
echo "}"                                                                        >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "int main(int argc, char** argv)"                                          >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      llog(INFO, \"%s\", \"Start the test for Project tt!\");"            >> test/$1Test.c
echo "      UNITY_BEGIN();"                                                     >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "      RUN_TEST(test_failure_func);"                                       >> test/$1Test.c
echo "      RUN_TEST(test_success_func);"                                       >> test/$1Test.c
echo "      RUN_TEST(test_ignore_func);"                                        >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "      UNITY_END();"                                                       >> test/$1Test.c
echo "      llog(INFO, \"%s\", \"Test For Project tt is Finished!\");"          >> test/$1Test.c
echo " }"                                                                       >> test/$1Test.c

touch test/CMakeLists.txt
echo "AUX_SOURCE_DIRECTORY(\${CMAKE_CURRENT_SOURCE_DIR} APP_SRC)"               > test/CMakeLists.txt
echo "ADD_EXECUTABLE($1Test \${APP_SRC})"                                       >> test/CMakeLists.txt
echo "SET(EXECUTABLE_OUTPUT_PATH \${PROJECT_BINARY_DIR}/bin)"                   >> test/CMakeLists.txt
echo "LINK_DIRECTORIES (\${PROJECT_BINARY_DIR}/lib)"                            >> test/CMakeLists.txt
echo "TARGET_LINK_LIBRARIES($1Test kkkmmu unity llog)"                          >> test/CMakeLists.txt
echo "INCLUDE_DIRECTORIES(\${PROJECT_SOURCE_DIR}/src \${PROJECT_SOURCE_DIR}/lib/unity \${PROJECT_SOURCE_DIR}/lib/llog/src)"        >> test/CMakeLists.txt

mkdir -p lib/
git clone https://github.com/lamproae/unity.git                                 lib/unity
# For mysql operation
# git clone https://github.com/lamproae/libmysql.git                              lib/mysql
git clone https://github.com/lamproae/c-algorithms.git                          lib/libds
git clone https://github.com/lamproae/Collections-C.git                         lib/collection
# Onion can be enabled if necessary
# git clone https://github.com/lamproae/onion.git                                 lib/onion
# libnet
# git clone https://github.com/lamproae/libnet.git                                lib/net
git clone https://github.com/lamproae/llog.git                                  lib/llog

#Common library for practise
# git clone https://github.com/lamproae/liblw.git                                 lib/liblw

# git clone https://github.com/lamproae/gperftools.git                            lib/gperftools


