package grpchandler

import (
	"context"
	"strconv"

	"github.com/evgensr/practicum1/internal/app"
	"github.com/evgensr/practicum1/internal/helper"
	"github.com/evgensr/practicum1/internal/pb"
)

func NewGRPCHandler(service app.Storage) *URLServer {
	return &URLServer{
		service: service,
	}
}

type URLServer struct {
	pb.UnimplementedURLServer
	service app.Storage
}

func (us *URLServer) Retrieve(ctx context.Context, in *pb.RetrieveRequest) (*pb.RetrieveResponse, error) {

	long, err := us.service.Get(in.ShortUrlId)
	if err != nil {

		return &pb.RetrieveResponse{
			Status: "internal server error",
		}, nil

	}
	return &pb.RetrieveResponse{
		RedirectUrl: long.URL,
		Status:      "ok",
	}, nil
}

func (us *URLServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {

	hash := helper.GetShort(in.OriginalUrl)
	err := us.service.Set(
		app.Line{
			URL:   in.OriginalUrl,
			User:  in.UserId,
			Short: hash,
		})
	if err != nil {
		return &pb.CreateResponse{
			Status: "internal server error",
		}, nil
	}
	return &pb.CreateResponse{
		Status:      "ok",
		ResponseUrl: hash,
	}, nil
}

func (us *URLServer) GetUserURLs(ctx context.Context, in *pb.GetUserURLsRequest) (*pb.GetUserURLsResponse, error) {
	urls := us.service.GetByUser(in.UserId)

	var result []*pb.GetUserURLsResponse_URL
	for i := 0; i < len(urls); i++ {
		result = append(result, &pb.GetUserURLsResponse_URL{
			OriginalUrl: urls[0].URL,
			ShortUrl:    urls[0].Short,
		})
	}
	return &pb.GetUserURLsResponse{
		Status: "ok",
		Urls:   result,
	}, nil
}

func (us *URLServer) CreateBatch(ctx context.Context, in *pb.CreateBatchRequest) (*pb.CreateBatchResponse, error) {

	type LineRequest struct {
		CorrelationID string `json:"correlation_id"`
		OriginalURL   string `json:"original_url"`
	}

	type LineResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	var response []*pb.CreateBatchResponse_URL

	for _, line := range in.Urls {
		// получаем short
		hash := helper.GetShort(line.OriginalUrl)
		// записываем в хранилище ключ значение
		us.service.Set(app.Line{
			URL:           line.OriginalUrl,
			Short:         hash,
			User:          in.UserId,
			CorrelationID: strconv.Itoa(int(line.CorrelationId)),
		})

		response = append(response, &pb.CreateBatchResponse_URL{
			CorrelationId: line.CorrelationId,
			ShortUrl:      hash,
		})
	}

	return &pb.CreateBatchResponse{
		Status: "ok",
		Urls:   response,
	}, nil
}

func (us *URLServer) DeleteBatch(ctx context.Context, in *pb.DeleteBatchRequest) (*pb.DeleteBatchResponse, error) {

	for _, row := range in.Urls {
		go us.service.Delete([]app.Line{
			{
				Short: row,
				User:  in.UserId,
			},
		})
	}

	return &pb.DeleteBatchResponse{
		Status: "accepted",
	}, nil
}

func (us *URLServer) GetStats(ctx context.Context, in *pb.GetStatsRequest) (*pb.GetStatsResponse, error) {

	urls, users, err := us.service.GetStats()

	if err != nil {
		return &pb.GetStatsResponse{
			Status: "forbidden",
		}, nil
	}
	if err != nil {
		return &pb.GetStatsResponse{
			Status: "bad request",
		}, nil
	}

	return &pb.GetStatsResponse{
		Status: "ok",
		Users:  int32(users),
		Urls:   int32(urls),
	}, nil
}
