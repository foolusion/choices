package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"flag"

	"github.com/foolusion/elwinprotos/elwin"
	"github.com/foolusion/elwinprotos/storage"
	"google.golang.org/grpc"
)

type server struct {
	sc storage.ElwinStorageClient
	ec elwin.ElwinClient
}

func (s *server) root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp, err := s.sc.List(context.TODO(), &storage.ListRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var epvRe = regexp.MustCompile("/experiment/(.*)/param/(.*)/value/(.*)")

func (s *server) value(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	matches := epvRe.FindStringSubmatch(r.URL.Path)
	log.Println(matches)
	for i := 0; i < 1000000; i++ {
		getResp, err := s.sc.Get(context.TODO(), &storage.GetRequest{
			Id: matches[1],
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		evalResp, err := s.ec.Eval(context.TODO(), &elwin.EvalRequest{
			UserID:      strconv.Itoa(i),
			Experiments: []*storage.Experiment{getResp.Experiment},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, e := range evalResp.Experiments {
			for _, p := range e.Params {
				if p.Name == matches[2] && p.Value == matches[3] {
					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					fmt.Fprintf(w, "%d", i)
					return
				}
			}
		}
	}
	http.Error(w, "could not find a valid userid", http.StatusInternalServerError)
}

var storageAddress = flag.String("storage_address", "localhost:8080", "address of storage server")
var elwinAddress = flag.String("elwin_address", "localhost:8083", "address of elwin")
var listenAddress = flag.String("listen_address", "localhost:8086", "address to listen on")

func main() {
	flag.Parse()
	ctx, sCancel := context.WithTimeout(context.Background(), time.Minute)
	defer sCancel()
	sc, err := grpc.DialContext(ctx, *storageAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	ctx, eCancel := context.WithTimeout(context.Background(), time.Minute)
	defer eCancel()
	ec, err := grpc.DialContext(ctx, *elwinAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	s := &server{
		sc: storage.NewElwinStorageClient(sc),
		ec: elwin.NewElwinClient(ec),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.root)
	mux.HandleFunc("/experiment/", s.value)

	server := http.Server{
		Handler:      mux,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Minute,
	}
	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
