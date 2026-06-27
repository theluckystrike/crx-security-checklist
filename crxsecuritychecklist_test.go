package crxsecuritychecklist
import "testing"
func TestIDs(t *testing.T) { if len(AllIDs()) < 8 { t.Fail() } }
