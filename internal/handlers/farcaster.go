package handlers

import (
	"fmt"
	"net/http"

	"farcaster-api-proxy-go/internal/upstream"
	"github.com/gorilla/mux"
)

const farcasterBaseURL = "https://client.farcaster.xyz/v2"

func UserThreadCasts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	hash := vars["hash"]
	
	if username == "" || hash == "" {
		http.Error(w, `{"error": "Missing hash or username"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/user-thread-casts?castHashPrefix=%s&username=%s&limit=5", farcasterBaseURL, hash, username)
	upstream.ProxyRequest(w, url)
}

func UserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	
	if username == "" {
		http.Error(w, `{"error": "Missing username"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/user-by-username?username=%s", farcasterBaseURL, username)
	upstream.ProxyRequest(w, url)
}

func UserByFid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fid := vars["fid"]
	
	if fid == "" {
		http.Error(w, `{"error": "Missing fid"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/user?fid=%s", farcasterBaseURL, fid)
	upstream.ProxyRequest(w, url)
}

func SearchCasts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "20"
	}
	
	if q == "" {
		http.Error(w, `{"error": "Missing q"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/search-casts?q=%s&limit=%s", farcasterBaseURL, q, limit)
	upstream.ProxyRequest(w, url)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	excludeSelf := r.URL.Query().Get("excludeSelf")
	if excludeSelf == "" {
		excludeSelf = "false"
	}
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "20"
	}
	includeDirectCastAbility := r.URL.Query().Get("includeDirectCastAbility")
	if includeDirectCastAbility == "" {
		includeDirectCastAbility = "false"
	}
	
	if q == "" {
		http.Error(w, `{"error": "Missing q"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/search-users?q=%s&excludeSelf=%s&limit=%s&includeDirectCastAbility=%s", farcasterBaseURL, q, excludeSelf, limit, includeDirectCastAbility)
	upstream.ProxyRequest(w, url)
}

func SearchSummary(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	maxChannels := r.URL.Query().Get("maxChannels")
	if maxChannels == "" {
		maxChannels = "2"
	}
	maxUsers := r.URL.Query().Get("maxUsers")
	if maxUsers == "" {
		maxUsers = "4"
	}
	maxMiniApps := r.URL.Query().Get("maxMiniApps")
	if maxMiniApps == "" {
		maxMiniApps = "2"
	}
	maxTokens := r.URL.Query().Get("maxTokens")
	if maxTokens == "" {
		maxTokens = "3"
	}
	addFollowersYouKnowContext := r.URL.Query().Get("addFollowersYouKnowContext")
	if addFollowersYouKnowContext == "" {
		addFollowersYouKnowContext = "false"
	}
	
	if q == "" {
		http.Error(w, `{"error": "Missing q"}`, http.StatusBadRequest)
		return
	}
	
	url := fmt.Sprintf("%s/search-summary?q=%s&maxChannels=%s&maxUsers=%s&maxMiniApps=%s&maxTokens=%s&addFollowersYouKnowContext=%s", farcasterBaseURL, q, maxChannels, maxUsers, maxMiniApps, maxTokens, addFollowersYouKnowContext)
	upstream.ProxyRequest(w, url)
}
