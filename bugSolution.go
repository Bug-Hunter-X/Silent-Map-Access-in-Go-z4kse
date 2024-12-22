func main() {
  var m = make(map[string]int)
  m["a"] = 1
  delete(m, "a")
  value, ok := m["a"]
  if ok {
    fmt.Println(value) // Access value if key exists
  } else {
    fmt.Println("Key does not exist") // Handle case of non-existent key
  }
}