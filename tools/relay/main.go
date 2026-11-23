// Mika's optional single-owner, single-device gateway. TLS terminates at the reverse proxy.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	listenDefault := "0.0.0.0:10000"
	if port := os.Getenv("PORT"); port != "" {
		listenDefault = net.JoinHostPort("0.0.0.0", port)
	}
	listen := flag.String("listen", listenDefault, "HTTP address behind a TLS reverse proxy")
	publicDefault := os.Getenv("RENDER_EXTERNAL_URL")
	public := flag.String("public-url", publicDefault, "public HTTPS origin, without a path")
	data := flag.String("data", "/data/relay.json", "persistent state file")
	ownerFile := flag.String("owner-key-file", "/data/owner-key", "file containing the owner secret (at least 32 characters)")
	initKey := flag.Bool("init-owner-key", false, "create an owner key file and exit; never overwrite")
	trustProxyDefault := os.Getenv("RENDER") != ""
	trustProxy := flag.Bool("trust-proxy", trustProxyDefault, "trust the reverse proxy address headers")
	local := flag.Bool("local-http", false, "development only: allow HTTP with both public URL and listen address on 127.0.0.1")
	flag.Parse()
	if *initKey {
		if err := os.MkdirAll(filepath.Dir(*ownerFile), 0700); err != nil {
			log.Fatal(err)
		}
		f, err := os.OpenFile(*ownerFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintln(f, secret())
		closeErr := f.Close()
		if err != nil {
			log.Fatal(err)
		}
		if closeErr != nil {
			log.Fatal(closeErr)
		}
		fmt.Println("Owner key created. Read the file locally; do not share it with MCP clients.")
		return
	}
	base, err := gatewayOrigin(*public, *listen, *local)
	if err != nil {
		log.Fatal("public-url must be an HTTPS origin")
	}
	var key []byte
	if envKey := strings.TrimSpace(os.Getenv("MIKA_OWNER_KEY")); envKey != "" {
		key = []byte(envKey)
	} else {
		key, err = os.ReadFile(*ownerFile)
		if err != nil {
			log.Fatal("owner key unavailable; set MIKA_OWNER_KEY or use -init-owner-key")
		}
	}
	if len(strings.TrimSpace(string(key))) < 32 {
		log.Fatal("owner key too short; it must contain at least 32 characters")
	}
	s, err := openStore(*data)
	if err != nil {
		log.Fatal(err)
	}
	g := newGateway(base, strings.TrimSpace(string(key)), s)
	g.trustProxy = *trustProxy
	srv := &http.Server{Addr: *listen, Handler: g, ReadHeaderTimeout: 10 * time.Second, IdleTimeout: 90 * time.Second, MaxHeaderBytes: 16 << 10}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		g.disconnect()
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdown)
	}()
	log.Printf("Mika relay listening on %s", *listen)
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

// Match the phone's URL canonicalization, including case and default ports.
func gatewayOrigin(raw, listen string, local bool) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	u.Scheme = strings.ToLower(u.Scheme)
	listenHost, _, _ := net.SplitHostPort(listen)
	localHTTP := local && u.Scheme == "http" && u.Hostname() == "127.0.0.1" && listenHost == "127.0.0.1"
	if u.Scheme != "https" && !localHTTP || u.Hostname() == "" || u.User != nil || u.RawQuery != "" || u.ForceQuery || strings.Contains(raw, "#") || (u.Path != "" && u.Path != "/") {
		return "", errors.New("invalid origin")
	}
	host, port := strings.ToLower(u.Hostname()), u.Port()
	if port != "" {
		n, e := strconv.Atoi(port)
		if e != nil || n < 1 || n > 65535 {
			return "", errors.New("invalid port")
		}
	}
	if u.Scheme == "https" && port == "443" || u.Scheme == "http" && port == "80" {
		port = ""
	}
	if port != "" {
		host = net.JoinHostPort(host, port)
	} else if strings.Contains(host, ":") {
		host = "[" + host + "]"
	}
	return u.Scheme + "://" + host, nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func failure(w http.ResponseWriter, status int, code string) {
	writeJSON(w, status, map[string]any{"error": code})
}

func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(dst); err != nil {
		failure(w, 400, "invalid_request")
		return false
	}
	if err := d.Decode(new(any)); err != io.EOF {
		failure(w, 400, "invalid_request")
		return false
	}
	return true
}

func bearer(r *http.Request) string {
	h := r.Header.Get("Authorization")
	if len(h) > 7 && strings.EqualFold(h[:7], "Bearer ") {
		return strings.TrimSpace(h[7:])
	}
	return ""
}
