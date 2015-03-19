package jargo

import (
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const jar1BaseUrl = "http://search.maven.org/remotecontent?filepath=org/apache/lucene/lucene-analyzers-common/5.0.0/"
const jar1Name = "lucene-analyzers-stempel-5.0.0.jar"

const badJarName = "zzz-foobar------_______"

//resp, err := http.Get("http://example.com/")"

func TestMain(m *testing.M) {
	//err := initTestJarFile()
	os.Exit(m.Run())
}

func TestValidJarFile_JarInfo(t *testing.T) {
	_, err := GetJarInfo(jar1Name)
	if err != nil {
		t.FailNow()
	}
}

func TestMissingJarFile_JarInfo(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	_, err := GetJarInfo(badJarName)
	if err == nil {
		t.FailNow()
	}
}

func TestValidJarFile_JarManifest(t *testing.T) {
	manifest, err := GetManifest(jar1Name)
	//err, jar := MakeJar("lucene-1.4.3.jar")
	//err, jar := MakeJar("/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.31-3.b13.fc21.x86_64/lib/tools.jar")
	//err, _ := MakeJar("/usr/lib/jvm-exports/java-1.8.0-openjdk-1.8.0.31-3.b13.fc21.x86_64/jaas-1.8.0.31.jar")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if manifest == nil {
		t.FailNow()
	}
	//fmt.Println(jar)
}
