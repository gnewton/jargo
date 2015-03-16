jargo
======
jargo is a simple library allowing

It has two functions:

    func GetManifest(filename string) (error, *Manifest)

It returns a pointer to a Manifest (map[string]string) which is the key:values pairs from the META-INF/MANIFEST.MF file

    func GetJarInfo(filename string) (error, *JarInfo)

* GetJarInfo extracts various info from a Java JAR file
* It extracts the Manifest (like GetManifest)
* It extracts an array of the filenames in the JAR file
* It returns a pointer to a JarInfo struct

###Example Usage###
Using Lucene analyzer jar file from: http://central.maven.org/maven2/org/apache/lucene/lucene-analyzers-stempel/5.0.0/lucene-analyzers-stempel-5.0.0.jar

```
package main

import (
	"fmt"
	"github.com/gnewton/jargo"
)

func main() {
	err, manifest := jargo.GetManifest("lucene-analyzers-stempel-5.0.0.jar")
	fmt.Println(err)
	fmt.Println(manifest)

	err, jar := jargo.GetJarInfo("lucene-analyzers-stempel-5.0.0.jar")
	fmt.Println(err)
	fmt.Printf("%+v\n", jar)
}
```

Output:
```
<nil>
&map[X-Compile-Source-JDK:1.7 X-Compile-Target-JDK:1.7 Ant-Version:Apache Ant 1.8.3 Created-By:1.7.0_55-b13 (Oracle Corporation) Extension-Name:org.apache.lucene Specification-Title:Lucene Search Engine: analyzers-stempel Implementation-Vendor:The Apache Software Foundation Manifest-Version:1.0 Specification-Version:5.0.0 Specification-Vendor:The Apache Software Foundation Implementation-Title:org.apache.lucene Implementation-Version:5.0.0 1659987 - anshumgupta - 2015-02-15 12:20 :34]
<nil>
&{Manifest:0xc208034038 Files:[META-INF/ META-INF/MANIFEST.MF META-INF/services/ org/ org/apache/ org/apache/lucene/ org/apache/lucene/analysis/ org/apache/lucene/analysis/pl/ org/apache/lucene/analysis/stempel/ org/egothor/ org/egothor/stemmer/ META-INF/services/org.apache.lucene.analysis.util.TokenFilterFactory org/apache/lucene/analysis/pl/PolishAnalyzer$DefaultsHolder.class org/apache/lucene/analysis/pl/PolishAnalyzer.class org/apache/lucene/analysis/pl/stemmer_20000.tbl org/apache/lucene/analysis/pl/stopwords.txt org/apache/lucene/analysis/stempel/StempelFilter.class org/apache/lucene/analysis/stempel/StempelPolishStemFilterFactory.class org/apache/lucene/analysis/stempel/StempelStemmer.class org/egothor/stemmer/Cell.class org/egothor/stemmer/Compile.class org/egothor/stemmer/Diff.class org/egothor/stemmer/DiffIt.class org/egothor/stemmer/Gener.class org/egothor/stemmer/Lift.class org/egothor/stemmer/MultiTrie.class org/egothor/stemmer/MultiTrie2.class org/egothor/stemmer/Optimizer.class org/egothor/stemmer/Optimizer2.class org/egothor/stemmer/Reduce$Remap.class org/egothor/stemmer/Reduce.class org/egothor/stemmer/Row.class org/egothor/stemmer/Trie$StrEnum.class org/egothor/stemmer/Trie.class META-INF/LICENSE.txt META-INF/NOTICE.txt]}
```
