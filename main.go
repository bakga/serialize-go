package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
)

type Person struct {
 Name string json:"name"
 Age  int    json:"age"
 City string json:"city"
}

// Сериализация объекта в файл
func SerializeObject(obj interface{}, filePath string) error {
 data, err := json.Marshal(obj)
 if err != nil {
  return err
 }

 err = ioutil.WriteFile(filePath, data, 0644)
 if err != nil {
  return err
 }

 return nil
}

// Десериализация объекта из файла
func DeserializeObject(filePath string) (*Person, error) {
 data, err := ioutil.ReadFile(filePath)
 if err != nil {
  return nil, err
 }

 var obj Person
 err = json.Unmarshal(data, &obj)
 if err != nil {
  return nil, err
 }

 return &obj, nil
}

func main() {
 data := Person{Name: "John", Age: 30, City: "New York"}
 filePath := "data.json"

 // Сериализация объекта в файл
 err := SerializeObject(data, filePath)
 if err != nil {
  fmt.Println("Ошибка при сериализации объекта:", err)
  return
 }

 // Десериализация объекта из файла
 deserializedData, err := DeserializeObject(filePath)
 if err != nil {
  fmt.Println("Ошибка при десериализации объекта:", err)
  return
 }

 fmt.Println(deserializedData) // &{John 30 New York}
}
