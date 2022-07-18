package player
import(
	"testing"

)

func TestNew(t *testing.T){  //Test___()
	newPlayer := New("Aishwarya Anand")
	actualName := newPlayer.GetName()
	actualHit := newPlayer.GetHit()
	actualMiss := newPlayer.GetMiss()
	expectedName := "Aishwarya Anand"
	var (expectedHit,expectedMiss uint = 0,0)

	if actualName != expectedName{
		t.Error("Actual is ",actualName,"but expected is ",expectedName)
	}
	if actualHit != expectedHit{
		t.Error("Actual is ",actualName,"but expected is ",expectedName)
	}
	if actualMiss != expectedMiss{
		t.Error("Actual is ",actualName,"but expected is ",expectedName)
	}
}

func TestIncrementHit(t *testing.T){
	newPlayer := New("Aishwarya Anand")
	for i:=0;i<5;i++{
		newPlayer.IncrementHit()
	}
	actualHit := newPlayer.GetHit()
	var expectedHit uint=5
	if actualHit != expectedHit{
		t.Error("Actual is ",actualHit,"but expected is ",expectedHit)
	}

}

func TestIncrementMiss(t *testing.T){
	newPlayer := New("Aishwarya Anand")
	for i:=0;i<5;i++{
		newPlayer.IncrementMiss()
	}
	actualMiss := newPlayer.GetMiss()
	var expectedMiss uint=5
	if actualMiss != expectedMiss{
		t.Error("Actual is ",actualMiss,"but expected is ",expectedMiss)
	}

}