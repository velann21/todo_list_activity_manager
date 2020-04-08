def firstTest() {
    print "Hello World"
}

def testTwo() {
    print "Hello World"
}

return [
    firstTest: this.&firstTest,
    testTwo: this.&testTwo
]