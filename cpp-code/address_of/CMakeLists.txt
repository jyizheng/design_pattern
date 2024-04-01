cmake_minimum_required(VERSION 3.10)  # Minimum version of CMake required

project(AddressOf)  # Name of your project

# Set C++ standard/version, e.g., C++11, C++14, C++17, etc.
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add executable target with source files
add_executable(AddressOf address_of.cc)

# If your project is split across multiple files, list them all as separate arguments to add_executable, like so:
# add_executable(NumberProject main.cpp RealNumber.cpp Complex.cpp Number.cpp)


