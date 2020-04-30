module main

go 1.13

require test/hello v0.0.0 // 虛擬版號

replace test/hello => ../test
