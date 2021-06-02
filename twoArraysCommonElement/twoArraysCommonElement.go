package main
import (
  "fmt"
)

func main(){
  fmt.Print("Hello World!")
  array1, array2 := []string{"a","b","c","x"},[]string{"cx","y","i"}
  fmt.Print(anyCommonElements(array1, array2))
}
func anyCommonElements(arr1, arr2 []string) bool{
  l1, l2 := len(arr1),len(arr2)
  
  if l1 == 0 || l2 == 0 {
  	return false
  }
  // only works if the arrays are sorted
  // for ; i < l1 && j<l2 ; {
  //   fmt.Println(i,j,l1,l2)
  // 	if arr1[i] == arr2[j]{
  // 		return true
  // 	}else if arr1[i] <= arr2[j]{
  // 		i++
  // 	}else{
  // 		j++
  // 	}
  // }
  mapX := make(map[string]int)
	for _, val := range arr1{
		mapX[val] = 1
	}
	for _,val := range arr2{
		if _, ok := mapX[val]; ok{
			return true
		}
	}
  return false 	 
}