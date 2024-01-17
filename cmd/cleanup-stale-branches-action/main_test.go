package main

import "testing"

func TestParseRepoName(t *testing.T) {
	tests := []struct {
		name     string
		repoName string
		want     string
	}{
		{"ValidRepo", "owner/repo", "repo"},
		{"EmptyRepo", "owner/", ""},
		{"NoSlash", "repo", "repo"},
		{"ExtraSlash", "owner/repo/extra", "repo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRepoName(tt.repoName); got != tt.want {
				t.Errorf("parseRepoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRepoOwner(t *testing.T) {
	tests := []struct {
		name     string
		repoName string
		want     string
	}{
		{"ValidRepo", "owner/repo", "owner"},
		{"EmptyOwner", "/repo", ""},
		{"NoSlash", "owner", "owner"},
		{"ExtraSlash", "owner/repo/extra", "owner"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRepoOwner(tt.repoName); got != tt.want {
				t.Errorf("parseRepoOwner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidRepoNameFormat(t *testing.T) {
	tests := []struct {
		name     string
		repoName string
		want     bool
	}{
		{"ValidFormat", "owner/repo", true},
		{"NoOwner", "/repo", false},
		{"NoRepo", "owner/", false},
		{"NoSlash", "owner", false},
		{"ExtraSlash", "owner/repo/extra", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRepoNameFormat(tt.repoName); got != tt.want {
				t.Errorf("isValidRepoNameFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		item  string
		want  bool
	}{
		{"Present", []string{"a", "b", "c"}, "b", true},
		{"NotPresent", []string{"a", "b", "c"}, "d", false},
		{"EmptySlice", []string{}, "a", false},
		{"EmptyString", []string{"a", "b", ""}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.slice, tt.item); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		prefixes []string
		str      string
		want     bool
	}{
		{
			name:     "Single prefix, positive match",
			prefixes: []string{"test"},
			str:      "testString",
			want:     true,
		},
		{
			name:     "Single prefix, negative match",
			prefixes: []string{"no"},
			str:      "testString",
			want:     false,
		},
		{
			name:     "Multiple prefixes, positive match",
			prefixes: []string{"no", "test", "yes"},
			str:      "testString",
			want:     true,
		},
		{
			name:     "Multiple prefixes, no match",
			prefixes: []string{"no", "none", "na"},
			str:      "testString",
			want:     false,
		},
		{
			name:     "Empty prefix list",
			prefixes: []string{},
			str:      "testString",
			want:     true,
		},
		{
			name:     "Empty string",
			prefixes: []string{"test"},
			str:      "",
			want:     false,
		},
		{
			name:     "Prefix longer than string",
			prefixes: []string{"longprefix"},
			str:      "short",
			want:     false,
		},
		{
			name:     "Exact match",
			prefixes: []string{"exact"},
			str:      "exact",
			want:     true,
		},
		{
			name:     "Case sensitivity",
			prefixes: []string{"Test"},
			str:      "testString",
			want:     false,
		},
		{
			name:     "Prefix is a substring, but not at start",
			prefixes: []string{"String"},
			str:      "testString",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startsWith(tt.prefixes, tt.str); got != tt.want {
				t.Errorf("startsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}