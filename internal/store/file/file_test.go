package file

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInMemoryRepository_GetByUser(t *testing.T) {

}

func BenchmarkSet(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	store := New("fileStoragePath")

	b.ResetTimer()
	b.Run("set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			b.StopTimer() // останавливаем таймер
			line := Line{
				User:          "1" + fmt.Sprint(rand.Intn(100000)),
				URL:           "https://test" + fmt.Sprint(rand.Intn(100000)) + ".com",
				Short:         "short" + fmt.Sprint(rand.Intn(100000)),
				CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
				Status:        0,
			}
			b.StartTimer() // возобновляем таймер

			_ = store.Set(line)
		}
	})

	b.Run("get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer() // останавливаем таймер
			line := Line{
				User:          "1" + fmt.Sprint(rand.Intn(100000)),
				URL:           "https://test" + fmt.Sprint(rand.Intn(100000)) + ".com",
				Short:         "short" + fmt.Sprint(rand.Intn(100000)),
				CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
				Status:        0,
			}
			_ = store.Set(line)
			b.StartTimer() // возобновляем таймер
			_, _ = store.Get(line.Short)
		}
	})

	b.Run("delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer() // останавливаем таймер
			line := Line{
				User:          "1" + fmt.Sprint(rand.Intn(100000)),
				URL:           "https://test" + fmt.Sprint(rand.Intn(100000)) + ".com",
				Short:         "short" + fmt.Sprint(rand.Intn(100000)),
				CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
				Status:        0,
			}
			_ = store.Set(line)
			b.StartTimer() // возобновляем таймер
			_ = store.Delete([]Line{
				{
					Short: line.Short,
					User:  line.User,
				},
			})
		}
	})
	b.Run("getByUser", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer() // останавливаем таймер
			line := Line{
				User:          "1" + fmt.Sprint(rand.Intn(100000)),
				URL:           "https://test" + fmt.Sprint(rand.Intn(100000)) + ".com",
				Short:         "short" + fmt.Sprint(rand.Intn(100000)),
				CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
				Status:        0,
			}
			_ = store.Set(line)
			b.StartTimer() // возобновляем таймер
			_ = store.GetByUser(line.User)
		}
	})

}
