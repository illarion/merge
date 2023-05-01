package merge

import "testing"

func TestMergeMaps(t *testing.T) {
	cases := []struct {
		name     string
		dst      map[string]interface{}
		src      map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "simple",
			dst: map[string]interface{}{
				"foo": "bar",
			},
			src: map[string]interface{}{
				"foo": "baz",
			},
			expected: map[string]interface{}{
				"foo": "baz",
			},
		},
		{
			name: "nested",
			dst: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			src: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "qux",
				},
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "qux",
				},
			},
		},
		{
			name: "nested with different keys",
			dst: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			src: map[string]interface{}{
				"foo": map[string]interface{}{},
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
		},
		{
			name: "different types",
			dst: map[string]interface{}{
				"foo": "bar",
			},
			src: map[string]interface{}{
				"foo": 13,
			},
			expected: map[string]interface{}{
				"foo": 13,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := Maps(c.dst, c.src)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !mapsEqual(actual, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}

func TestMergeMapsDstInSrc(t *testing.T) {

	dst := map[string]interface{}{
		"foo": "bar",
	}

	src := map[string]interface{}{
		"foo":       "baz",
		"recursive": dst,
	}

	expected := map[string]interface{}{
		"foo": "baz",
		"recursive": map[string]interface{}{
			"foo": "bar",
		},
	}

	actual, err := Maps(dst, src)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !mapsEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}

}

func mapsEqual(actual map[string]interface{}, expected map[string]interface{}) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, v := range actual {
		am, aok := v.(map[string]interface{})
		em, eok := expected[k].(map[string]interface{})
		if aok && eok {
			return mapsEqual(am, em)
		}
		if aok != eok {
			return false
		}
		if v != expected[k] {
			return false
		}
	}

	return true
}
