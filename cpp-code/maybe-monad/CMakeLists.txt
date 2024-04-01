cmake_minimum_required(VERSION 3.10)  # Minimum version of CMake required

project(Maybe)  # Name of your project

# Set C++ standard/version, e.g., C++11, C++14, C++17, etc.
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add executable target with source files
add_executable(Maybe maybe.cc)



