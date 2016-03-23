package main

import (
    "fmt"
    //"io/ioutil"
    "log"
    "net/http"
    "bufio"
)


func main() {
    
    /*
    // making buckets for some ASCII runes
    for i := 65; i < 122; i++ {
        fmt.Println(i, " - ", string(i), " - ", i % 11)
    }
    */
    
    
    
    
    
    /*
    // printing all moby dick in console
    res, err := http.Get("http://gutenberg.org/files/2701/old/moby10b.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    
    bs, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%s", bs)
    */
    
    
    
    
    /*
    printing all moby dick in console - one word in one line
    res, err := http.Get("http://gutenberg.org/files/2701/old/moby10b.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    
    scanner := bufio.NewScanner(res.Body)
    defer res.Body.Close()
    
    scanner.Split(bufio.ScanWords)
    
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    */
    
    
    
    quantity := 11
    
    res, err := http.Get("http://gutenberg.org/files/2701/old/moby10b.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    
    scanner := bufio.NewScanner(res.Body)
    defer res.Body.Close()
    
    scanner.Split(bufio.ScanWords)
    
    buckets := make([][]string, quantity)
    
    // create slices to hold words
    for i := 0; i < quantity; i++ {
        buckets = append(buckets, []string{})
    }
    
    for scanner.Scan() {
        word := scanner.Text()
        n := HashBucket(word, quantity)
        buckets[n] = append(buckets[n], word)
    }
    
    for i := 0; i < quantity; i++ {
        fmt.Println(i, " - ", len(buckets[i]))
    }
    
    
    // Print the words in one of the buckets
    //fmt.Println(buckets[6])
}


func HashBucket (word string, buckets int) int {
    var sum int
    
    for _, v := range word {
        sum += int(v)
    }
    return sum % buckets
}