cmake_minimum_required(VERSION 3.10)
project(lab3 CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_compile_options(--coverage)
add_link_options(--coverage)

find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})
include_directories(include)

add_executable(run_google tests/tests_google.cpp)
target_link_libraries(run_google GTest::gtest_main pthread)
find_package(benchmark REQUIRED)
add_executable(run_bench bench/bench_main.cpp)
target_link_libraries(run_bench benchmark::benchmark)