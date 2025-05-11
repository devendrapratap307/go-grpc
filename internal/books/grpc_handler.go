package books

import (
	"context"
	pb "go-fiber-grpc/proto/books"
)

type GRPCHandler struct {
	pb.UnimplementedBookServiceServer
	service Service
}

func NewGRPCHandler(s Service) *GRPCHandler {
	return &GRPCHandler{service: s}
}

func (h *GRPCHandler) CreateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	book := &Book{
		Title:     req.GetTitle(),
		Author:    req.GetAuthor(),
		Year:      int(req.GetYear()),
		Publisher: req.GetPublisher(),
	}
	if err := h.service.CreateBook(book); err != nil {
		return nil, err
	}
	return toProto(book), nil
}

func (h *GRPCHandler) GetBook(ctx context.Context, req *pb.BookID) (*pb.Book, error) {
	book, err := h.service.GetBook(uint(req.GetId()))
	if err != nil {
		return nil, err
	}
	return toProto(book), nil
}

func (h *GRPCHandler) ListBooks(ctx context.Context, _ *pb.Empty) (*pb.BookList, error) {
	books, err := h.service.ListBooks()
	if err != nil {
		return nil, err
	}
	var protoBooks []*pb.Book
	for _, b := range books {
		protoBooks = append(protoBooks, toProto(&b))
	}
	return &pb.BookList{Books: protoBooks}, nil
}

func (h *GRPCHandler) UpdateBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	book := &Book{
		ID:        uint(req.GetId()),
		Title:     req.GetTitle(),
		Author:    req.GetAuthor(),
		Year:      int(req.GetYear()),
		Publisher: req.GetPublisher(),
	}
	if err := h.service.UpdateBook(book); err != nil {
		return nil, err
	}
	return toProto(book), nil
}

func (h *GRPCHandler) DeleteBook(ctx context.Context, req *pb.BookID) (*pb.Empty, error) {
	err := h.service.DeleteBook(uint(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func toProto(book *Book) *pb.Book {
	return &pb.Book{
		Id:        uint32(book.ID),
		Title:     book.Title,
		Author:    book.Author,
		Year:      int32(book.Year),
		Publisher: book.Publisher,
	}
}
