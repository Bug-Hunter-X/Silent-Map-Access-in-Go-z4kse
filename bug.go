func main() {
  var m = make(map[string]int)
  m["a"] = 1
  delete(m, "a")
  fmt.Println(m["a"]) // this will print 0, not throw an error
}