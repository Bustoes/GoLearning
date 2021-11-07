module package_test/main

go 1.17

replace package_test/calc => ./../calc

require package_test/calc v0.0.0-00010101000000-000000000000
