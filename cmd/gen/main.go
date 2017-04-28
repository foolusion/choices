package main

import (
	"fmt"
	"html/template"
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
	resp, err := s.sc.List(context.TODO(), &storage.ListRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type data struct {
		DetailName, ExperimentName, ParamName, ParamValue, Team, Platform string
	}
	var out []data
	for _, e := range resp.Experiments {
		for _, p := range e.Params {
			for _, c := range p.Value.Choices {
				d := data{
					DetailName:     e.DetailName,
					ExperimentName: e.Name,
					ParamName:      p.Name,
					ParamValue:     c,
					Team:           e.Labels["team"],
					Platform:       e.Labels["platform"],
				}
				out = append(out, d)
			}
		}
	}
	if err := rootTmpl.Execute(w, out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var rootTmpl = template.Must(template.ParseFiles("root.html"))

var epvRe = regexp.MustCompile("/experiment/(.*)/param/(.*)/value/(.*)")

func (s *server) value(w http.ResponseWriter, r *http.Request) {
	matches := epvRe.FindStringSubmatch(r.URL.Path)
	log.Println(matches)
	for i := 0; i < 1000000; i++ {
		resp, err := s.ec.Get(context.TODO(), &elwin.GetRequest{
			UserID: strconv.Itoa(i),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, e := range resp.Experiments {
			for _, p := range e.Params {
				if e.Name == matches[1] && p.Name == matches[2] && p.Value == matches[3] {
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
