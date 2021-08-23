cmake_minimum_required(VERSION 3.15)
project(leetcode)

set(CMAKE_CXX_STANDARD 11)
# 显示全部警告/错误 并且开启runtime检测器
# -fno-sanitize-recover=all
# -gdwarf-3 用于更友好的显示STL容器
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Werror -fsanitize=address,undefined -gdwarf-3")

# 遍历solutions目录下所有的 .cpp 文件
file (GLOB files solutions/*.cpp)
foreach (file ${files})
    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" exe ${file})
    add_executable (${exe} ${file})
    message (${exe}.cpp)
endforeach ()