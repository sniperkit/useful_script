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
include .lib/memwatch/CMakeFiles/memwatch_static.dir/depend.make

# Include the progress variables for this target.
include .lib/memwatch/CMakeFiles/memwatch_static.dir/progress.make

# Include the compile flags for this target's objects.
include .lib/memwatch/CMakeFiles/memwatch_static.dir/flags.make

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o: .lib/memwatch/CMakeFiles/memwatch_static.dir/flags.make
.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o: lib/memwatch/memwatch.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/leeweop/kkkmmu/pppp/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o"
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/memwatch_static.dir/memwatch.c.o   -c /home/leeweop/kkkmmu/pppp/lib/memwatch/memwatch.c

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/memwatch_static.dir/memwatch.c.i"
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/leeweop/kkkmmu/pppp/lib/memwatch/memwatch.c > CMakeFiles/memwatch_static.dir/memwatch.c.i

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/memwatch_static.dir/memwatch.c.s"
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/leeweop/kkkmmu/pppp/lib/memwatch/memwatch.c -o CMakeFiles/memwatch_static.dir/memwatch.c.s

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.requires:

.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.requires

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.provides: .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.requires
	$(MAKE) -f .lib/memwatch/CMakeFiles/memwatch_static.dir/build.make .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.provides.build
.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.provides

.lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.provides.build: .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o


# Object files for target memwatch_static
memwatch_static_OBJECTS = \
"CMakeFiles/memwatch_static.dir/memwatch.c.o"

# External object files for target memwatch_static
memwatch_static_EXTERNAL_OBJECTS =

lib/libmemwatch.a: .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o
lib/libmemwatch.a: .lib/memwatch/CMakeFiles/memwatch_static.dir/build.make
lib/libmemwatch.a: .lib/memwatch/CMakeFiles/memwatch_static.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/leeweop/kkkmmu/pppp/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C static library ../../lib/libmemwatch.a"
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && $(CMAKE_COMMAND) -P CMakeFiles/memwatch_static.dir/cmake_clean_target.cmake
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/memwatch_static.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
.lib/memwatch/CMakeFiles/memwatch_static.dir/build: lib/libmemwatch.a

.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/build

.lib/memwatch/CMakeFiles/memwatch_static.dir/requires: .lib/memwatch/CMakeFiles/memwatch_static.dir/memwatch.c.o.requires

.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/requires

.lib/memwatch/CMakeFiles/memwatch_static.dir/clean:
	cd /home/leeweop/kkkmmu/pppp/.lib/memwatch && $(CMAKE_COMMAND) -P CMakeFiles/memwatch_static.dir/cmake_clean.cmake
.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/clean

.lib/memwatch/CMakeFiles/memwatch_static.dir/depend:
	cd /home/leeweop/kkkmmu/pppp && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/leeweop/kkkmmu/pppp /home/leeweop/kkkmmu/pppp/lib/memwatch /home/leeweop/kkkmmu/pppp /home/leeweop/kkkmmu/pppp/.lib/memwatch /home/leeweop/kkkmmu/pppp/.lib/memwatch/CMakeFiles/memwatch_static.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : .lib/memwatch/CMakeFiles/memwatch_static.dir/depend
