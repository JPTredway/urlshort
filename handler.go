package urlshort

import "net/http"

// MapHandler will return an http.HandlerFunc (which also implements
// http.Handler) that will attempt to map any paths (keys in their map)
// to their corresponding URL (values that each key in the map points to,
// in string format). If the path is not provided in the map, then the
// fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return an
// http.HandlerFunc (which also implements http.Handler) that will
// attempt to map any paths to their corresponding URL. If the path
// is not provided in the YAML, then the fallback http.Handler will
// be called instead.
//
// YAML is expected to be in the format:
//
//		- path: /somepath
//			url: https://www.some-url.com/demo
//
// The only errors that can be returned are all related to having
// invalid YAML data.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO
	return nil, nil
}
