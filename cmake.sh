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
echo "ADD_SUBDIRECTORY(src)"                                                    >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(app)"                                                    >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(lib)"                                                    >> CMakeLists.txt
echo "SET(CMAKE_VERBOSE_MAKEFILE on)"                                           >> CMakeLists.txt
echo "ADD_SUBDIRECTORY(test)"                                                   >> CMakeLists.txt

touch lib/CMakeLists.txt
echo "ADD_SUBDIRECTORY(unity)"                                                   > lib/CMakeLists.txt
echo "SET(CMAKE_VERBOSE_MAKEFILE on)"                                           >> lib/CMakeLists.txt

touch src/$1.c
echo "#include <stdio.h>"                                                        > src/$1.c
echo "void $1(void)"                                                            >> src/$1.c
echo "{"                                                                        >> src/$1.c
echo "      printf(\"Hello world!\");"                                          >> src/$1.c
echo "}"                                                                        >> src/$1.c
touch src/$1.h
echo "#ifndef __$1_H__"                                                         > src/$1.h
echo "#define __$1_H__"                                                         >> src/$1.h
echo "void $1(void);"                                                           >> src/$1.h
echo "#endif /* __$1_H__ */"                                                    >> src/$1.h
touch src/CMakeLists.txt
echo "SET(LIB_SRC $1.c)"                                                        > src/CMakeLists.txt
echo "ADD_LIBRARY(lib$1 SHARED \${LIB_SRC})"                                    >> src/CMakeLists.txt
echo "SET(LIBRARY_OUTPUT_PATH \${PROJECT_BINARY_DIR}/lib)"                      >> src/CMakeLists.txt
echo "SET_TARGET_PROPERTIES(lib$1 PROPERTIES OUTPUT_NAME "$1")"                 >> src/CMakeLists.txt
echo "INSTALL(TARGETS lib$1"                                                    >> src/CMakeLists.txt
echo "      LIBRARY DESTINATION lib"                                            >> src/CMakeLists.txt      
echo "      ARCHIVE DESTINATION lib)"                                           >> src/CMakeLists.txt
echo "INSTALL(FILES $1.h DESTINATION include/$1)"                               >> src/CMakeLists.txt

touch app/$1Main.c
echo "#include \"$1.h\""                                                         > app/$1Main.c
echo "int main(int argc, char** argv)"                                          >> app/$1Main.c
echo "{"                                                                        >> app/$1Main.c
echo "      $1();"                                                              >> app/$1Main.c
echo "}"                                                                        >> app/$1Main.c

touch app/CMakeLists.txt
echo "INCLUDE_DIRECTORIES(\${PROJECT_SOURCE_DIR}/src)"                           > app/CMakeLists.txt
echo "SET(APP_SRC $1Main.c)"                                                    >> app/CMakeLists.txt
echo "SET(EXECUTABLE_OUTPUT_PATH \${PROJECT_BINARY_DIR}/bin)"                   >> app/CMakeLists.txt
echo "ADD_EXECUTABLE($1 \${APP_SRC})"                                           >> app/CMakeLists.txt
echo "TARGET_LINK_LIBRARIES($1 lib$1)"                                          >> app/CMakeLists.txt

touch test/$1Test.c
echo "#include \"$1.h\""                                                         > test/$1Test.c
echo "#include \"unity.h\""                                                     >> test/$1Test.c
echo "#include \"unity_internals.h\""                                           >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "void test_dummy_func(void)  "                                             >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      $1();"                                                              >> test/$1Test.c
echo "      TEST_ASSERT_EQUAL_STRING(\"Hello\", \"Hello\");"                    >> test/$1Test.c
echo "      TEST_ASSERT_EQUAL(16, 8);"                                          >> test/$1Test.c
echo "}"                                                                        >> test/$1Test.c
echo ""                                                                         >> test/$1Test.c
echo "int main(int argc, char** argv)"                                          >> test/$1Test.c
echo "{"                                                                        >> test/$1Test.c
echo "      UnityBegin(\"test/$1Test.c\");"                                    >> test/$1Test.c
echo "      $1();"                                                              >> test/$1Test.c
echo "      RUN_TEST(test_dummy_func);"                                         >> test/$1Test.c
echo "}"                                                                        >> test/$1Test.c

touch test/CMakeLists.txt
echo "INCLUDE_DIRECTORIES(\${PROJECT_SOURCE_DIR}/src \${PROJECT_SOURCE_DIR}/lib/unity)"                           > test/CMakeLists.txt
echo "SET(APP_SRC $1Test.c)"           >> test/CMakeLists.txt
echo "SET(EXECUTABLE_OUTPUT_PATH \${PROJECT_BINARY_DIR}/bin)"                   >> test/CMakeLists.txt
echo "ADD_EXECUTABLE($1Test \${APP_SRC})"                                       >> test/CMakeLists.txt
echo "TARGET_LINK_LIBRARIES($1Test lib$1 libunity)"                                      >> test/CMakeLists.txt

mkdir -p lib/
git clone https://github.com/lamproae/unity.git  lib/unity
