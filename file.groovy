def firstTest(String str) {
    println(str);
}

def testTwo() {
    print "Hello World"
}

return [
    firstTest: this.&firstTest,
    testTwo: this.&testTwo
]