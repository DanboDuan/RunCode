#!/bin/bash
set -o pipefail
set -eux
git clone git@github.com:google/googletest.git
cd googletest && mkdir build && cd build
cmake ..
make -j8
sudo make install