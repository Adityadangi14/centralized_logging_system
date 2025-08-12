package initializers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	db "github.com/Adityadangi14/centralized_logging_system/log_collector/db/gen"
	"github.com/Adityadangi14/centralized_logging_system/log_collector/models"
)

type TCPServer struct {
	Listener    net.Listener
	StopChan    chan struct{}
	LogChan     chan string
	WorkerPool  chan string
	Wg          sync.WaitGroup
	Q           *db.Queries
	WorkerCount int
}

func NewTCPServer(address string, workerCount int) (*TCPServer, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	dns := os.Getenv("postgresurl")

	q, err := InitPostgres(PostgresConnector{}, dns)

	if err != nil {
		fmt.Println(err)

	}

	server, err := &TCPServer{
		Listener:    listener,
		StopChan:    make(chan struct{}),
		LogChan:     make(chan string, 2),
		WorkerPool:  make(chan string, 2),
		Q:           q,
		WorkerCount: workerCount,
	}, nil

	for i := 0; i <= workerCount; i++ {
		server.Wg.Add(1)

	}

	return server, nil
}

func (s *TCPServer) Start() {

	go s.LogHandler()

	for i := 0; i < s.WorkerCount; i++ {
		s.Wg.Add(1)
		go s.Worker()
	}
	go func() {
		for {
			conn, err := s.Listener.Accept()
			if err != nil {
				select {
				case <-s.StopChan:
					return
				default:
					log.Println("Accept error:", err)
					continue
				}
			}

			go s.TcpConnHandlers(conn)

		}
	}()
}

func (s *TCPServer) TcpConnHandlers(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		logLine := scanner.Text()
		s.LogChan <- logLine

	}
	if err := scanner.Err(); err != nil {
		log.Println("Connection read error:", err)
	}
}

func (s *TCPServer) LogHandler() {
	for log := range s.LogChan {

		s.WorkerPool <- log
	}
}

func (s *TCPServer) Worker() {
	defer s.Wg.Done()

	for {
		select {
		case logMsg := <-s.WorkerPool:

			s.Parselogs(logMsg)
		case <-s.StopChan:
			return
		}
	}
}

func (s *TCPServer) Stop() {
	close(s.StopChan)
	s.Listener.Close()
}

func (s *TCPServer) Parselogs(log string) {

	var parsedLog models.ParsedLog

	err := json.Unmarshal([]byte(log), &parsedLog)

	if err != nil {
		fmt.Println(err)
		return
	}

	s.WriteLog(parsedLog)

}

func (s *TCPServer) WriteLog(log models.ParsedLog) (db.ParsedLog, error) {
	ctx := context.Background()

	res, err := s.Q.InsertParsedLog(ctx, db.InsertParsedLogParams{
		Timestamp:     log.Timestamp,
		EventCategory: log.EventCategory,
		SourceType:    log.SourceType,
		Username:      log.Username,
		Hostname:      log.Hostname,
		Severity:      log.Severity,
		RawMessage:    log.RawMessage,
	})

	if err != nil {
		fmt.Println(err)
		return db.ParsedLog{}, err
	}

	return res, nil
}
