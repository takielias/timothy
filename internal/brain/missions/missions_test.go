package missions

import "testing"

func TestRepoSourceReadsEitherKind(t *testing.T) {
	t.Parallel()
	gh := Mission{Sources: []SourceEntry{{Source: SourceKindGitHub, RepoURL: "https://github.com/o/r", ConnectorID: "c1"}}}
	bb := Mission{Sources: []SourceEntry{{Source: SourceKindBitbucket, RepoURL: "https://bitbucket.org/w/s", ConnectorID: "c2"}}}
	none := Mission{Sources: []SourceEntry{{Source: SourceKindPDF, ID: "a"}}}

	if bb.RepoURL() != "https://bitbucket.org/w/s" || bb.ConnectorID() != "c2" {
		t.Fatalf("bitbucket source not read: %q %q", bb.RepoURL(), bb.ConnectorID())
	}
	if _, ok := bb.GitHubSource(); ok {
		t.Fatal("GitHubSource must stay github-only")
	}
	if e, ok := bb.RepoSource(); !ok || e.Source != SourceKindBitbucket {
		t.Fatalf("RepoSource = %+v, %v", e, ok)
	}
	if e, ok := gh.GitHubSource(); !ok || e.ConnectorID != "c1" {
		t.Fatalf("GitHubSource = %+v, %v", e, ok)
	}
	if none.RepoURL() != "" || none.ConnectorID() != "" {
		t.Fatal("mission without a repo source must read empty")
	}
}
