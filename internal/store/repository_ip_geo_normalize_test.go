package store

import "testing"

func TestNormalizeIPGeoLocation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		domestic     string
		global       string
		wantDomestic string
		wantGlobal   string
	}{
		{
			name:         "non-china unchanged",
			domestic:     "Tokyo",
			global:       "Japan",
			wantDomestic: "Tokyo",
			wantGlobal:   "Japan",
		},
		{
			name:         "china remove carrier keyword",
			domestic:     "广东·深圳·电信",
			global:       "中国",
			wantDomestic: "广东·深圳",
			wantGlobal:   "中国",
		},
		{
			name:         "china only carrier fallback china",
			domestic:     "电信",
			global:       "china",
			wantDomestic: "中国",
			wantGlobal:   "china",
		},
		{
			name:         "china idc with province use province",
			domestic:     "腾讯云·广东深圳",
			global:       "中国",
			wantDomestic: "广东",
			wantGlobal:   "中国",
		},
		{
			name:         "china idc without province use datacenter",
			domestic:     "阿里云",
			global:       "中国",
			wantDomestic: "机房",
			wantGlobal:   "中国",
		},
		{
			name:         "avoid false province extraction from cloud-city token",
			domestic:     "腾讯云南京",
			global:       "中国",
			wantDomestic: "机房",
			wantGlobal:   "中国",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotDomestic, gotGlobal := normalizeIPGeoLocation(tt.domestic, tt.global)
			if gotDomestic != tt.wantDomestic || gotGlobal != tt.wantGlobal {
				t.Fatalf("normalizeIPGeoLocation(%q,%q) = (%q,%q), want (%q,%q)",
					tt.domestic, tt.global, gotDomestic, gotGlobal, tt.wantDomestic, tt.wantGlobal)
			}
		})
	}
}
