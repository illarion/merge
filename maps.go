package merge

import "fmt"

// Maps merges multiple maps into one.
// The first map is the destination map, the following maps are the source maps.
// The source maps are merged into the destination map, overwriting existing keys.
// If the destination map is nil, a new map is created and returned. This is useful if you want to merge multiple maps into a new map,
// without modifying any of the source maps.
// Maps follows these rules:
// * If the maps contain nested maps, they are merged recursively.
// * If the maps contain other types, the values of the source map are used.
// * If the source maps contain keys that are not present in the destination map, they are added to the destination map.
// * If the source maps contain keys that are present in the destination map, the values of the source map are used, overwriting the values of the destination map.
// * If the destination map contains keys that are not present in the source maps, they are not modified.
func Maps(dst map[string]interface{}, srcs ...map[string]interface{}) (map[string]interface{}, error) {
	for _, src := range srcs {
		var err error
		dst, err = maps(dst, src, true)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func maps(dst, src map[string]interface{}, clone bool) (map[string]interface{}, error) {

	if dst == nil {
		dst = make(map[string]interface{})
	}

	if src == nil {
		return dst, nil
	}

	// clone the source map if requested. Cloning is necessary to avoid
	// overwriting the source map when merging nested maps.
	if clone {
		src = deepCloneMap(src)
	}

	for k, sv := range src {
		dv, ok := dst[k]

		if !ok {
			dst[k] = sv
			continue
		}

		dstMap, isDstMap := dv.(map[string]interface{})
		if !isDstMap {
			dst[k] = sv
			continue
		}

		// if both values in both maps are maps, we have to merge them recursively
		srcMap, isSrcMap := sv.(map[string]interface{})

		if !isSrcMap {
			dst[k] = sv
			continue
		}

		if _, err := maps(dstMap, srcMap, false); err != nil {
			return nil, fmt.Errorf("error while merging nested map at key %q: %w", k, err)
		}
	}

	return dst, nil
}

func deepCloneMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}

	result := make(map[string]interface{}, len(m))

	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			result[k] = deepCloneMap(val)
		default:
			result[k] = v
		}
	}

	return result
}
