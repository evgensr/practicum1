package pg

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/evgensr/practicum1/internal/helper"
)

type PgRepositoryTestSuite struct {
	suite.Suite
	store *Box
}

func (s *PgRepositoryTestSuite) SetupSuite() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/restapi?sslmode=disable"
	}
	store := New(dsn)
	s.store = store
}

func (s *PgRepositoryTestSuite) TearDownTest() {
	_, err := s.store.db.Exec("truncate table short")
	require.NoError(s.T(), err)
}

func (s *PgRepositoryTestSuite) TearDownSuite() {
	_, err := s.store.db.Exec("truncate table short")
	require.NoError(s.T(), err)
}

func (s *PgRepositoryTestSuite) TestSave() {
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

func TestPgRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PgRepositoryTestSuite))
}

func TestInMemoryRepository_GetByUser(t *testing.T) {

}

func BenchmarkPg(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/restapi?sslmode=disable"
	}
	store := New(dsn)

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
				CorrelationID: "1" + fmt.Sprint(rand.Intn(1000)),
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
				CorrelationID: "1" + fmt.Sprint(rand.Intn(1000)),
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
				CorrelationID: "1" + fmt.Sprint(rand.Intn(1000)),
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
				CorrelationID: "1" + fmt.Sprint(rand.Intn(1000)),
				Status:        0,
			}
			_ = store.Set(line)
			b.StartTimer() // возобновляем таймер
			_ = store.GetByUser(line.User)
		}
	})

	_, _ = store.db.Exec("truncate table short")

}
