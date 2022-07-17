package memory

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"

	"github.com/evgensr/practicum1/internal/helper"
)

type MemoryRepositoryTestSuite struct {
	suite.Suite
	store *Box
}

func (s *MemoryRepositoryTestSuite) SetupSuite() {

	store := New("")
	s.store = store
}

func (s *MemoryRepositoryTestSuite) TestSet() {
	var fetched Line
	rand.Seed(time.Now().UnixNano())
	URL := "https://test" + helper.GeneratorUUID() + ".com"
	short := helper.GetHash(URL)
	line := Line{
		User:          helper.GeneratorUUID(),
		URL:           URL,
		Short:         short,
		CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
		Status:        0,
	}
	err := s.store.Set(line)
	require.NoError(s.T(), err)

	s.store.RLock()
	defer s.store.RUnlock()

	for _, u := range s.store.Items {
		if u.Short == line.Short {
			fetched = u
		}
	}

	assert.Equal(s.T(), line, fetched)
}

func (s *MemoryRepositoryTestSuite) TestGet() {
	rand.Seed(time.Now().UnixNano())
	URL := "https://test" + helper.GeneratorUUID() + ".com"
	short := helper.GetHash(URL)
	line := Line{
		User:          helper.GeneratorUUID(),
		URL:           URL,
		Short:         short,
		CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
		Status:        0,
	}
	err := s.store.Set(line)
	require.NoError(s.T(), err)

	fetched, err := s.store.Get(line.Short)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), line, fetched)

}

func (s *MemoryRepositoryTestSuite) TestGetByUser() {
	rand.Seed(time.Now().UnixNano())
	URL := "https://test" + helper.GeneratorUUID() + ".com"
	short := helper.GetHash(URL)
	line := Line{
		User:          helper.GeneratorUUID(),
		URL:           URL,
		Short:         short,
		CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
		Status:        0,
	}
	err := s.store.Set(line)
	require.NoError(s.T(), err)

	fetched := s.store.GetByUser(line.User)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), []Line{{
		User:          line.User,
		URL:           line.URL,
		Short:         line.Short,
		CorrelationID: line.CorrelationID,
		Status:        line.Status,
	}}, fetched)

}

func (s *MemoryRepositoryTestSuite) TestDelete() {
	rand.Seed(time.Now().UnixNano())
	URL := "https://test" + helper.GeneratorUUID() + ".com"
	short := helper.GetHash(URL)
	line := Line{
		User:          helper.GeneratorUUID(),
		URL:           URL,
		Short:         short,
		CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
		Status:        0,
	}
	err := s.store.Set(line)
	require.NoError(s.T(), err)

	err = s.store.Delete([]Line{{
		User:          line.User,
		URL:           line.URL,
		Short:         line.Short,
		CorrelationID: line.CorrelationID,
		Status:        line.Status,
	}})
	require.NoError(s.T(), err)

	fetched, err := s.store.Get(line.Short)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), Line{
		User:          line.User,
		URL:           line.URL,
		Short:         line.Short,
		CorrelationID: line.CorrelationID,
		Status:        1,
	}, fetched)

}

func TestMemoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryRepositoryTestSuite))
}

func BenchmarkSet(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	store := New("fileStoragePath")

	b.ResetTimer()
	b.Run("set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			b.StopTimer() // останавливаем таймер
			URL := "https://test" + helper.GeneratorUUID() + ".com"
			short := helper.GetHash(URL)
			line := Line{
				User:          helper.GeneratorUUID(),
				URL:           URL,
				Short:         short,
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
			URL := "https://test" + helper.GeneratorUUID() + ".com"
			short := helper.GetHash(URL)
			line := Line{
				User:          helper.GeneratorUUID(),
				URL:           URL,
				Short:         short,
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
			URL := "https://test" + helper.GeneratorUUID() + ".com"
			short := helper.GetHash(URL)
			line := Line{
				User:          helper.GeneratorUUID(),
				URL:           URL,
				Short:         short,
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
			URL := "https://test" + helper.GeneratorUUID() + ".com"
			short := helper.GetHash(URL)
			line := Line{
				User:          helper.GeneratorUUID(),
				URL:           URL,
				Short:         short,
				CorrelationID: "1" + fmt.Sprint(rand.Intn(100000)),
				Status:        0,
			}
			_ = store.Set(line)
			b.StartTimer() // возобновляем таймер
			_ = store.GetByUser(line.User)
		}
	})

}
