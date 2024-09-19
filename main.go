package main
import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`


}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content--Type","application/json")
    json.NewEncoder(w).Encode(movies)//json.NewEncoder(w) hume http ko data bejta h aur encode data ko json mai convert krdeta h 
}
func delMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content--Type","application/json")
	params:=mux.Vars(r)
	for k,v := range movies{
			if v.ID==params["id"]{
				movies = append(movies[:k], movies[k+1:]...)
                break
			}
	}
    json.NewEncoder(w).Encode(movies)


}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content--Type","application/json")
    params:=mux.Vars(r)
    for _, item := range movies{
        if item.ID==params["id"]{
            json.NewEncoder(w).Encode(item)
            break
        }
    }
}
func createMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content--Type","application/json")
    var m Movie
    _ = json.NewDecoder(r.Body).Decode(&m)
	if m.ID == ""{
		m.ID = strconv.Itoa(rand.Intn(1000000))

	}
    movies = append(movies, m)
    json.NewEncoder(w).Encode(m)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item := range movies{
		if params["id"]==item.ID{
			movies = append(movies[:index], movies[index+1:]...)
            var m Movie
            _ = json.NewDecoder(r.Body).Decode(&m)
            m.ID = params["id"]
            movies = append(movies, m)
            json.NewEncoder(w).Encode(m)
            return
		}
	}
}
var movies []Movie

func main(){
	movies =append(movies,Movie{ID: "1", Isbn: "234543", Title: "movie_1",Director: &Director{Firstname: "fname1", Lastname:"lname1"}})
	movies =append(movies,Movie{ID: "5", Isbn: "876567", Title: "movie 5",Director: &Director{Firstname: "fname5", Lastname:"lname5"}})
	movies =append(movies,Movie{ID: "4", Isbn: "987678", Title: "movie 4",Director: &Director{Firstname: "fname4", Lastname:"lname4"}})
	movies =append(movies,Movie{ID: "3", Isbn: "09878", Title: "movie 3",Director: &Director{Firstname: "fname3", Lastname:"lname3"}})
	movies =append(movies,Movie{ID: "2", Isbn: "987656", Title: "movie 2",Director: &Director{Firstname: "fname2", Lastname:"lname2"}})
	r:=mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",delMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}