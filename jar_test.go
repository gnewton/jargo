package jargo

import (
	//"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) { os.Exit(m.Run()) }

func TestTimeConsuming(t *testing.T) {
	//err, jar := MakeJar("lucene-analyzers-stempel-5.0.0.jar")
	//err, jar := MakeJar("lucene-1.4.3.jar")
	//err, jar := MakeJar("/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.31-3.b13.fc21.x86_64/lib/tools.jar")
	err, _ := MakeJar("/usr/lib/jvm-exports/java-1.8.0-openjdk-1.8.0.31-3.b13.fc21.x86_64/jaas-1.8.0.31.jar")
	if err != nil {
		t.FailNow()
	}
	//fmt.Println(jar)
}
