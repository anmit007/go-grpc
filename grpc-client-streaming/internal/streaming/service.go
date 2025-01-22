package streaming

import (
	"io"
	"log"
	"time"

	"github.com/anmit007/go-grpc/grpc-streaming/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

func (s Service) StreamServerTime(request *proto.StreamServerTimeRequest, stream proto.StreamingService_StreamServerTimeServer) error {

	if request.GetIntervalSeconds() == 0 {
		return status.Error(codes.InvalidArgument, "interval cannot be 0")
	}
	// iniit ticker for our interval
	interval := time.Duration(request.GetIntervalSeconds()) * time.Second
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			currentTime := time.Now()

			resp := &proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}

}
func (s *Service) LogStream(stream proto.StreamingService_LogStreamServer) error {
	// init a count
	count := 0

	// loop through all received messages
	for {
		logEntry, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.LogStreamResponse{
				EntriesLogged: int32(count),
			})
		}
		if err != nil {
			return err // Don't log.Fatal here as it will kill the server
		}
		//log message
		log.Printf("Recieved Log [%s]: %s %s", logEntry.GetTimestamp().AsTime(), logEntry.GetLevel().Type().Descriptor().Syntax().String(), logEntry.GetMessage())
		//increament count
		count++
	}

}
