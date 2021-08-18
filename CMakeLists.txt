cmake_minimum_required(VERSION 3.15)
project(leetcode)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fsanitize=address")
# 遍历solutions目录下所有的 .cpp 文件
file (GLOB files solutions/*.cpp)
foreach (file ${files})
    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" exe ${file})
    add_executable (${exe} ${file})
    message (${exe}.cpp)
endforeach ()