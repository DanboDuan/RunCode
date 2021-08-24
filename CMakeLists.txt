cmake_minimum_required(VERSION 3.14)
project(leetcode)

if (NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Debug)
endif()

# GoogleTest requires at least C++11
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

## this will download every time
# include(FetchContent)
# FetchContent_Declare(
#   googletest
#   URL https://github.com/google/googletest/archive/609281088cfefc76f9d0ce82e1ff6c30cc3591e5.zip
# )
# FetchContent_MakeAvailable(googletest)

# enable_testing()
# add_executable(
#   hello_test
#   hello_test.cc
# )
# target_link_libraries(
#   hello_test
#   gtest_main
# )

# include(GoogleTest)
# gtest_discover_tests(hello_test)

## use installed GoogleTest
enable_testing()
find_package(Threads REQUIRED)
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIR})

include_directories(${CMAKE_SOURCE_DIR}/include)
# add_subdirectory(include)

add_subdirectory(src)
add_subdirectory(test)
