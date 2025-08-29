# CMake generated Testfile for 
# Source directory: /home/joe/projects/ai-daily/cpp/example/test
# Build directory: /home/joe/projects/ai-daily/cpp/example/test/build
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
add_test(YourTest "/home/joe/projects/ai-daily/cpp/example/test/build/main")
set_tests_properties(YourTest PROPERTIES  _BACKTRACE_TRIPLES "/home/joe/projects/ai-daily/cpp/example/test/CMakeLists.txt;14;add_test;/home/joe/projects/ai-daily/cpp/example/test/CMakeLists.txt;0;")
subdirs("_deps/googletest-build")
