package main

import "testing"

func BenchmarkIoTest_1024_1024_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 1024, 32)
    }
}

func BenchmarkIoTest_1024_1024_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024,1024, 64)
    }
}

func BenchmarkIoTest_1024_1024_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 1024, 128)
    }
}

func BenchmarkIoTest_1024_1024_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 1024, 256)
    }
}

func BenchmarkIoTest_1024_1024_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 1024, 512)
    }
}

func BenchmarkIoTest_1024_1024_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 1024, 1024)
    }
}

func BenchmarkIoTest_2048_2048_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 32)
    }
}

func BenchmarkIoTest_2048_2048_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 64)
    }
}

func BenchmarkIoTest_2048_2048_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 128)
    }
}

func BenchmarkIoTest_2048_2048_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 256)
    }
}

func BenchmarkIoTest_2048_2048_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 512)
    }
}

func BenchmarkIoTest_2048_2048_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 1024)
    }
}

func BenchmarkIoTest_2048_2048_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 2048, 2048)
    }
}

func BenchmarkIoTest_4096_4096_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 32)
    }
}

func BenchmarkIoTest_4096_4096_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 64)
    }
}

func BenchmarkIoTest_4096_4096_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 128)
    }
}

func BenchmarkIoTest_4096_4096_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 256)
    }
}

func BenchmarkIoTest_4096_4096_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 512)
    }
}

func BenchmarkIoTest_4096_4096_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 1024)
    }
}

func BenchmarkIoTest_4096_4096_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 2048)
    }
}

func BenchmarkIoTest_4096_4096_4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 4096, 4096)
    }
}

func BenchmarkIoTest_2048_1024_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 32)
    }
}

func BenchmarkIoTest_2048_1024_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 64)
    }
}

func BenchmarkIoTest_2048_1024_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 128)
    }
}

func BenchmarkIoTest_2048_1024_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 256)
    }
}

func BenchmarkIoTest_2048_1024_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 512)
    }
}

func BenchmarkIoTest_2048_1024_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 1024, 1024)
    }
}

func BenchmarkIoTest_4096_2048_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 32)
    }
}

func BenchmarkIoTest_4096_2048_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 64)
    }
}

func BenchmarkIoTest_4096_2048_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 128)
    }
}

func BenchmarkIoTest_4096_2048_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 256)
    }
}

func BenchmarkIoTest_4096_2048_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 512)
    }
}

func BenchmarkIoTest_4096_2048_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 1024)
    }
}

func BenchmarkIoTest_4096_2048_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 2048, 2048)
    }
}

func BenchmarkIoTest_4096_1024_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 32)
    }
}

func BenchmarkIoTest_4096_1024_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 64)
    }
}

func BenchmarkIoTest_4096_1024_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 128)
    }
}

func BenchmarkIoTest_4096_1024_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 256)
    }
}

func BenchmarkIoTest_4096_1024_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 512)
    }
}

func BenchmarkIoTest_4096_1024_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(4096, 1024, 1024)
    }
}

func BenchmarkIoTest_1024_2048_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 32)
    }
}

func BenchmarkIoTest_1024_2048_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024,2048, 64)
    }
}

func BenchmarkIoTest_1024_2048_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 128)
    }
}

func BenchmarkIoTest_1024_2048_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 256)
    }
}

func BenchmarkIoTest_1024_2048_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 512)
    }
}

func BenchmarkIoTest_1024_2048_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 1024)
    }
}

func BenchmarkIoTest_1024_2048_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 2048, 2048)
    }
}

func BenchmarkIoTest_1024_4096_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 32)
    }
}

func BenchmarkIoTest_1024_4096_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024,4096, 64)
    }
}

func BenchmarkIoTest_1024_4096_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 128)
    }
}

func BenchmarkIoTest_1024_4096_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 256)
    }
}

func BenchmarkIoTest_1024_4096_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 512)
    }
}

func BenchmarkIoTest_1024_4096_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 1024)
    }
}

func BenchmarkIoTest_1024_4096_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 2048)
    }
}

func BenchmarkIoTest_1024_4096_4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(1024, 4096, 4096)
    }
}

func BenchmarkIoTest_2048_4096_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 32)
    }
}

func BenchmarkIoTest_2048_4096_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 64)
    }
}

func BenchmarkIoTest_2048_4096_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 128)
    }
}

func BenchmarkIoTest_2048_4096_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 256)
    }
}

func BenchmarkIoTest_2048_4096_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 512)
    }
}

func BenchmarkIoTest_2048_4096_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 1024)
    }
}

func BenchmarkIoTest_2048_4096_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 2048)
    }
}

func BenchmarkIoTest_2048_4096_4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
        ioTest(2048, 4096, 4096)
    }
}

// func BenchmarkIoTest_1024(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//         ioTest2(1024)
//     }
// }

// func BenchmarkIoTest_2048(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//         ioTest2(2048)
//     }
// }

// func BenchmarkIoTest_4096(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//         ioTest2(4096)
//     }
// }

func BenchmarkIoTest_1024_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 32)
	}
}

func BenchmarkIoTest_1024_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 64)
	}
}

func BenchmarkIoTest_1024_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 128)
	}
}

func BenchmarkIoTest_1024_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 256)
	}
}

func BenchmarkIoTest_1024_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 512)
	}
}

func BenchmarkIoTest_1024_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(1024, 1024)
	}
}

func BenchmarkIoTest_2048_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 32)
	}
}

func BenchmarkIoTest_2048_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 64)
	}
}

func BenchmarkIoTest_2048_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 128)
	}
}

func BenchmarkIoTest_2048_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 256)
	}
}

func BenchmarkIoTest_2048_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 512)
	}
}

func BenchmarkIoTest_2048_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 1024)
	}
}

func BenchmarkIoTest_2048_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(2048, 2048)
	}
}

func BenchmarkIoTest_4096_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 32)
	}
}

func BenchmarkIoTest_4096_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 64)
	}
}

func BenchmarkIoTest_4096_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 128)
	}
}

func BenchmarkIoTest_4096_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 256)
	}
}

func BenchmarkIoTest_4096_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 512)
	}
}

func BenchmarkIoTest_4096_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 1024)
	}
}

func BenchmarkIoTest_4096_2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 2048)
	}
}

func BenchmarkIoTest_4096_4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioTest3(4096, 4096)
	}
}