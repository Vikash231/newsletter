package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Article represents a news article in InShorts style
type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"` // max 60 words
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Source      string    `json:"source"`
	SourceURL   string    `json:"source_url"`
	ImageURL    string    `json:"image_url"`
	PublishedAt time.Time `json:"published_at"`
	ReadTime    int       `json:"read_time"` // seconds
}

// Response wraps API responses
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Total   int         `json:"total,omitempty"`
}

var articles = []Article{
	{
		ID:          1,
		Title:       "SpaceX Starship Completes First Fully Successful Round-Trip Test Flight",
		Content:     "SpaceX's Starship rocket completed a landmark test flight, successfully launching from Texas and landing both the Super Heavy booster and upper stage spacecraft. This marks a major milestone for Elon Musk's vision of making humanity multi-planetary and reducing the cost of space travel significantly.",
		Category:    "technology",
		Author:      "Tech Desk",
		Source:      "Space News",
		SourceURL:   "https://example.com/spacex-starship",
		ImageURL:    "https://picsum.photos/seed/spacex/800/450",
		PublishedAt: time.Now().Add(-2 * time.Hour),
		ReadTime:    45,
	},
	{
		ID:          2,
		Title:       "India Wins T20 World Cup 2026 in a Thrilling Final Against Australia",
		Content:     "India clinched the T20 World Cup title defeating Australia by 6 wickets in a nail-biting final. Virat Kohli's unbeaten 82 off 54 balls guided India to victory, chasing 187 runs with 2 balls to spare. The stadium erupted in celebration as India lifted the trophy.",
		Category:    "sports",
		Author:      "Sports Desk",
		Source:      "Cricket Today",
		SourceURL:   "https://example.com/t20-worldcup",
		ImageURL:    "https://picsum.photos/seed/cricket/800/450",
		PublishedAt: time.Now().Add(-3 * time.Hour),
		ReadTime:    40,
	},
	{
		ID:          3,
		Title:       "Global Markets Rally as Inflation Cools to 3-Year Low",
		Content:     "Stock markets worldwide surged after inflation data showed consumer prices rising at the slowest pace in three years. The S&P 500 gained 2.1% while European markets rose over 1.5%. Analysts expect central banks to begin cutting interest rates sooner than previously anticipated.",
		Category:    "business",
		Author:      "Finance Desk",
		Source:      "Market Watch",
		SourceURL:   "https://example.com/markets-rally",
		ImageURL:    "https://picsum.photos/seed/markets/800/450",
		PublishedAt: time.Now().Add(-4 * time.Hour),
		ReadTime:    38,
	},
	{
		ID:          4,
		Title:       "Scientists Discover Potential Cure for Alzheimer's Disease in Early Trials",
		Content:     "Researchers at MIT have identified a protein that could halt the progression of Alzheimer's disease. Early clinical trials involving 200 patients showed 70% reduction in memory decline. Scientists caution that larger trials are needed, but call it the most promising breakthrough in decades.",
		Category:    "science",
		Author:      "Health Desk",
		Source:      "Medical Journal",
		SourceURL:   "https://example.com/alzheimers-cure",
		ImageURL:    "https://picsum.photos/seed/science1/800/450",
		PublishedAt: time.Now().Add(-5 * time.Hour),
		ReadTime:    42,
	},
	{
		ID:          5,
		Title:       "Apple Unveils iPhone 17 with Revolutionary AI Camera System",
		Content:     "Apple announced iPhone 17 featuring an AI-powered camera that can generate photorealistic images in real time. The device includes a 48MP periscope lens with 10x optical zoom and can shoot 8K video. Pre-orders begin Friday with availability in 40 countries from next week.",
		Category:    "technology",
		Author:      "Tech Desk",
		Source:      "Tech Crunch",
		SourceURL:   "https://example.com/iphone-17",
		ImageURL:    "https://picsum.photos/seed/apple/800/450",
		PublishedAt: time.Now().Add(-6 * time.Hour),
		ReadTime:    35,
	},
	{
		ID:          6,
		Title:       "Bollywood Blockbuster 'Mahakaal' Crosses ₹500 Crore in Opening Week",
		Content:     "Hrithik Roshan's mythological action film Mahakaal shattered box office records earning ₹500 crore in its first week, making it the fastest Bollywood film to reach this milestone. The film received 4.5-star reviews praising its visual effects, performances, and screenplay.",
		Category:    "entertainment",
		Author:      "Entertainment Desk",
		Source:      "Bollywood News",
		SourceURL:   "https://example.com/mahakaal",
		ImageURL:    "https://picsum.photos/seed/bollywood/800/450",
		PublishedAt: time.Now().Add(-7 * time.Hour),
		ReadTime:    32,
	},
	{
		ID:          7,
		Title:       "New Study Links Ultra-Processed Foods to Increased Risk of Depression",
		Content:     "A study of over 100,000 adults found that consuming ultra-processed foods daily increases depression risk by 23%. Researchers from Harvard University tracked participants for 10 years. They recommend replacing processed foods with whole grains, fruits, and vegetables to protect mental health.",
		Category:    "health",
		Author:      "Health Desk",
		Source:      "Health Weekly",
		SourceURL:   "https://example.com/food-depression",
		ImageURL:    "https://picsum.photos/seed/health1/800/450",
		PublishedAt: time.Now().Add(-8 * time.Hour),
		ReadTime:    40,
	},
	{
		ID:          8,
		Title:       "India Launches World's Largest Solar Park in Rajasthan",
		Content:     "Prime Minister inaugurated a 30 GW solar park in Rajasthan, making it the world's largest renewable energy installation. The ₹1.5 lakh crore project will power 20 million homes and reduce carbon emissions by 60 million tonnes annually, significantly boosting India's clean energy goals.",
		Category:    "science",
		Author:      "Environment Desk",
		Source:      "Energy Today",
		SourceURL:   "https://example.com/solar-park",
		ImageURL:    "https://picsum.photos/seed/solar/800/450",
		PublishedAt: time.Now().Add(-9 * time.Hour),
		ReadTime:    38,
	},
	{
		ID:          9,
		Title:       "Neymar Returns to Barcelona in Shock Transfer Deal",
		Content:     "Brazilian superstar Neymar has rejoined FC Barcelona in a stunning transfer deal worth €45 million. The 34-year-old left Al-Hilal after two seasons and signed a two-year contract. Barcelona fans celebrated outside the Nou Camp as Neymar promised to help the club win La Liga.",
		Category:    "sports",
		Author:      "Sports Desk",
		Source:      "Football News",
		SourceURL:   "https://example.com/neymar-barca",
		ImageURL:    "https://picsum.photos/seed/football/800/450",
		PublishedAt: time.Now().Add(-10 * time.Hour),
		ReadTime:    35,
	},
	{
		ID:          10,
		Title:       "Tesla Announces $25,000 Electric Car for Indian Market",
		Content:     "Tesla revealed plans to launch an affordable electric car priced at ₹20 lakh for the Indian market by 2027. The compact model will offer 400 km range on a single charge. Elon Musk said India represents Tesla's biggest growth opportunity with its 1.4 billion population.",
		Category:    "business",
		Author:      "Auto Desk",
		Source:      "Auto World",
		SourceURL:   "https://example.com/tesla-india",
		ImageURL:    "https://picsum.photos/seed/tesla/800/450",
		PublishedAt: time.Now().Add(-11 * time.Hour),
		ReadTime:    37,
	},
	{
		ID:          11,
		Title:       "Netflix Drops Teaser for Season 2 of 'Sacred Games' After 6 Years",
		Content:     "Netflix India surprised fans by releasing a teaser for Sacred Games Season 2, six years after the first season. Saif Ali Khan and Nawazuddin Siddiqui reprise their roles. The 8-episode season premieres on April 15th and is already trending on social media with excitement.",
		Category:    "entertainment",
		Author:      "Entertainment Desk",
		Source:      "OTT Today",
		SourceURL:   "https://example.com/sacred-games-2",
		ImageURL:    "https://picsum.photos/seed/netflix/800/450",
		PublishedAt: time.Now().Add(-12 * time.Hour),
		ReadTime:    33,
	},
	{
		ID:          12,
		Title:       "WHO Declares Monkeypox Outbreak Contained; Cases Drop 90% Globally",
		Content:     "The World Health Organization announced that the monkeypox outbreak is now effectively contained after vaccination campaigns reached 50 million people across 40 countries. Global cases have fallen 90% from peak levels. WHO credited rapid international cooperation and vaccine distribution for the successful response.",
		Category:    "health",
		Author:      "Health Desk",
		Source:      "Health Wire",
		SourceURL:   "https://example.com/monkeypox-contained",
		ImageURL:    "https://picsum.photos/seed/health2/800/450",
		PublishedAt: time.Now().Add(-13 * time.Hour),
		ReadTime:    40,
	},
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /api/news - get all news with optional category filter
func handleGetNews(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 && v <= 50 {
			limit = v
		}
	}

	filtered := []Article{}
	for _, a := range articles {
		if category == "" || strings.EqualFold(a.Category, category) {
			filtered = append(filtered, a)
		}
	}

	total := len(filtered)
	start := (page - 1) * limit
	end := start + limit
	if start >= total {
		filtered = []Article{}
	} else {
		if end > total {
			end = total
		}
		filtered = filtered[start:end]
	}

	jsonResponse(w, http.StatusOK, Response{
		Success: true,
		Data:    filtered,
		Total:   total,
	})
}

// GET /api/news/{id} - get a single article by ID
func handleGetArticle(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		jsonResponse(w, http.StatusBadRequest, Response{Success: false, Error: "invalid article ID"})
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{Success: false, Error: "invalid article ID"})
		return
	}

	for _, a := range articles {
		if a.ID == id {
			jsonResponse(w, http.StatusOK, Response{Success: true, Data: a})
			return
		}
	}
	jsonResponse(w, http.StatusNotFound, Response{Success: false, Error: "article not found"})
}

// GET /api/categories - list all categories
func handleGetCategories(w http.ResponseWriter, r *http.Request) {
	seen := map[string]bool{}
	categories := []string{}
	for _, a := range articles {
		if !seen[a.Category] {
			seen[a.Category] = true
			categories = append(categories, a.Category)
		}
	}
	jsonResponse(w, http.StatusOK, Response{Success: true, Data: categories})
}

// GET /api/health
func handleHealth(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func router(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	switch {
	case path == "/api/health":
		handleHealth(w, r)
	case path == "/api/categories":
		handleGetCategories(w, r)
	case path == "/api/news" && r.Method == http.MethodGet:
		handleGetNews(w, r)
	case strings.HasPrefix(path, "/api/news/") && r.Method == http.MethodGet:
		handleGetArticle(w, r)
	default:
		jsonResponse(w, http.StatusNotFound, Response{Success: false, Error: "endpoint not found"})
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", corsMiddleware(router))

	port := ":8080"
	log.Printf("InShorts Clone API running on http://localhost%s", port)
	log.Printf("Endpoints:")
	log.Printf("  GET /api/health")
	log.Printf("  GET /api/categories")
	log.Printf("  GET /api/news?category=<cat>&page=<n>&limit=<n>")
	log.Printf("  GET /api/news/:id")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
