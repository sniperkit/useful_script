# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.8

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/leeweop/kkkmmu/pppp

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/leeweop/kkkmmu/pppp

# Include any dependencies generated for this target.
include .app/CMakeFiles/pppp.dir/depend.make

# Include the progress variables for this target.
include .app/CMakeFiles/pppp.dir/progress.make

# Include the compile flags for this target's objects.
include .app/CMakeFiles/pppp.dir/flags.make

.app/CMakeFiles/pppp.dir/instrument.c.o: .app/CMakeFiles/pppp.dir/flags.make
.app/CMakeFiles/pppp.dir/instrument.c.o: app/instrument.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/leeweop/kkkmmu/pppp/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object .app/CMakeFiles/pppp.dir/instrument.c.o"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/pppp.dir/instrument.c.o   -c /home/leeweop/kkkmmu/pppp/app/instrument.c

.app/CMakeFiles/pppp.dir/instrument.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/pppp.dir/instrument.c.i"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/leeweop/kkkmmu/pppp/app/instrument.c > CMakeFiles/pppp.dir/instrument.c.i

.app/CMakeFiles/pppp.dir/instrument.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/pppp.dir/instrument.c.s"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/leeweop/kkkmmu/pppp/app/instrument.c -o CMakeFiles/pppp.dir/instrument.c.s

.app/CMakeFiles/pppp.dir/instrument.c.o.requires:

.PHONY : .app/CMakeFiles/pppp.dir/instrument.c.o.requires

.app/CMakeFiles/pppp.dir/instrument.c.o.provides: .app/CMakeFiles/pppp.dir/instrument.c.o.requires
	$(MAKE) -f .app/CMakeFiles/pppp.dir/build.make .app/CMakeFiles/pppp.dir/instrument.c.o.provides.build
.PHONY : .app/CMakeFiles/pppp.dir/instrument.c.o.provides

.app/CMakeFiles/pppp.dir/instrument.c.o.provides.build: .app/CMakeFiles/pppp.dir/instrument.c.o


.app/CMakeFiles/pppp.dir/ppppMain.c.o: .app/CMakeFiles/pppp.dir/flags.make
.app/CMakeFiles/pppp.dir/ppppMain.c.o: app/ppppMain.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/leeweop/kkkmmu/pppp/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building C object .app/CMakeFiles/pppp.dir/ppppMain.c.o"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/pppp.dir/ppppMain.c.o   -c /home/leeweop/kkkmmu/pppp/app/ppppMain.c

.app/CMakeFiles/pppp.dir/ppppMain.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/pppp.dir/ppppMain.c.i"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/leeweop/kkkmmu/pppp/app/ppppMain.c > CMakeFiles/pppp.dir/ppppMain.c.i

.app/CMakeFiles/pppp.dir/ppppMain.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/pppp.dir/ppppMain.c.s"
	cd /home/leeweop/kkkmmu/pppp/.app && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/leeweop/kkkmmu/pppp/app/ppppMain.c -o CMakeFiles/pppp.dir/ppppMain.c.s

.app/CMakeFiles/pppp.dir/ppppMain.c.o.requires:

.PHONY : .app/CMakeFiles/pppp.dir/ppppMain.c.o.requires

.app/CMakeFiles/pppp.dir/ppppMain.c.o.provides: .app/CMakeFiles/pppp.dir/ppppMain.c.o.requires
	$(MAKE) -f .app/CMakeFiles/pppp.dir/build.make .app/CMakeFiles/pppp.dir/ppppMain.c.o.provides.build
.PHONY : .app/CMakeFiles/pppp.dir/ppppMain.c.o.provides

.app/CMakeFiles/pppp.dir/ppppMain.c.o.provides.build: .app/CMakeFiles/pppp.dir/ppppMain.c.o


# Object files for target pppp
pppp_OBJECTS = \
"CMakeFiles/pppp.dir/instrument.c.o" \
"CMakeFiles/pppp.dir/ppppMain.c.o"

# External object files for target pppp
pppp_EXTERNAL_OBJECTS =

bin/pppp: .app/CMakeFiles/pppp.dir/instrument.c.o
bin/pppp: .app/CMakeFiles/pppp.dir/ppppMain.c.o
bin/pppp: .app/CMakeFiles/pppp.dir/build.make
bin/pppp: lib/libkkkmmu.so.1.1
bin/pppp: lib/libfault.so.1.1
bin/pppp: .lib/llog/lib/libllog.so.1.1
bin/pppp: lib/libran.so
bin/pppp: .app/CMakeFiles/pppp.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/leeweop/kkkmmu/pppp/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking C executable ../bin/pppp"
	cd /home/leeweop/kkkmmu/pppp/.app && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/pppp.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
.app/CMakeFiles/pppp.dir/build: bin/pppp

.PHONY : .app/CMakeFiles/pppp.dir/build

.app/CMakeFiles/pppp.dir/requires: .app/CMakeFiles/pppp.dir/instrument.c.o.requires
.app/CMakeFiles/pppp.dir/requires: .app/CMakeFiles/pppp.dir/ppppMain.c.o.requires

.PHONY : .app/CMakeFiles/pppp.dir/requires

.app/CMakeFiles/pppp.dir/clean:
	cd /home/leeweop/kkkmmu/pppp/.app && $(CMAKE_COMMAND) -P CMakeFiles/pppp.dir/cmake_clean.cmake
.PHONY : .app/CMakeFiles/pppp.dir/clean

.app/CMakeFiles/pppp.dir/depend:
	cd /home/leeweop/kkkmmu/pppp && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/leeweop/kkkmmu/pppp /home/leeweop/kkkmmu/pppp/app /home/leeweop/kkkmmu/pppp /home/leeweop/kkkmmu/pppp/.app /home/leeweop/kkkmmu/pppp/.app/CMakeFiles/pppp.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : .app/CMakeFiles/pppp.dir/depend
