cmake_minimum_required(VERSION 3.14)
project(YourProject)

include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/refs/tags/v1.14.0.zip
)
FetchContent_MakeAvailable(googletest)

add_executable(main main.cpp)
target_link_libraries(main gtest_main)
enable_testing()
add_test(NAME YourTest COMMAND main)
