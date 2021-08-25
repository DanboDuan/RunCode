aux_source_directory(. EXECUTABLE)
foreach(T_FILE_NAME ${EXECUTABLE})
    string(REGEX REPLACE ".cpp" "" EXECUTABLE_NAME ${T_FILE_NAME})
    string(REGEX REPLACE "\.\/" "" EXECUTABLE_NAME ${EXECUTABLE_NAME})
    message(${EXECUTABLE_NAME})
    add_executable(${EXECUTABLE_NAME} ${T_FILE_NAME})
    target_link_libraries(
        ${EXECUTABLE_NAME}
        PRIVATE
        ${CMAKE_THREAD_LIBS_INIT}
        GTest::GTest
        GTest::Main
    )
    gtest_discover_tests(${EXECUTABLE_NAME})
endforeach()