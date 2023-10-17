package main
import "fmt"
func Merge_sort(A *[]int,start int,last int){
    if start<last{
       mid :=(start+last)/2
       fmt.Printf("mid is:%d",mid)
       Merge_sort(A, start, mid)
       Merge_sort(A,mid+1,last)
       Merge(A,start,mid,last)
       PrintArr(A)
    }
}
func Merge(A *[]int,start,mid,last int){
   n1:=mid-start+1
   n2:=last-mid
   l:=make([]int,n1 )
   r:= make([]int, n2)
   for i:=0;i<n1;i++{
     l[i] = (*A)[start+i]
   }
   for j:=0;j<n2;j++{
     r[j] = (*A)[mid+1+j]
   }
   i:=0
   j:=0
   k:=start
   for i<n1 && j<n2{
     if l[i]<r[j]{
        (*A)[k] =l[i]
        i++
        k++
     }else{
        (*A)[k]=r[j]
         j++
         k++
     }
  }
  for i==n1 &&j<n2{
     (*A)[k] = r[j]
      j++
      k++
  }
  for j==n2 &&i<n1{
     (*A)[k] = l[i]
      i++
      k++
  }
}
func PrintArr(A *[] int){
    for i:=0; i<len(*A);i++{
       fmt.Printf("A[%d]:%d\n", i ,(*A)[i])
    }
}
func main(){
    A:=[]int{38,27,43,3,9,82,10}
    Merge_sort(&A, 0,(len(A)-1))
    fmt.Printf("Array after merge sort:\n")
    PrintArr(&A)
}