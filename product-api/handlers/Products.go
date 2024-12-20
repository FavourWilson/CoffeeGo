package handlers

import (
	"log"
	"net/http"
	"github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "Unable to unmarshal json",http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	
}