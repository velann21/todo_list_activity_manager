package main

//import (
//	"fmt"
//	"strconv"
//)
//
//var str = make([]string,0)
//var st string
//func main() {
//    strin := "velan"
//	for i, r := range strin {
//		if (i+1)%2 == 0 {
//			st += string(r)
//			str = append(str, st)
//			st = ""
//		} else {
//			st = string(r)
//			if len(strin) == i+1{
//				if len(st) == 1{
//					str = append(str, st+"_")
//				}
//			}
//			continue
//		}
//
//	}
//
//	fmt.Println(str)
//	fmt.Println(CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0}))
//
////	sli := make([]string, 0)
////	sli = append(sli,"velann22")
////	other(sli)
////}
////
////func other(str []string) {
////    str = append(str, "velann23")
//}
//
//
//func CreatePhoneNumber(numbers [10]uint) string {
//	numLen := len(numbers)
//	if numLen < 10 || numLen > 10{
//		return "Invalid"
//	}
//	formStr := "("
//	for i,v := range numbers{
//
//		if i <= 2{
//			formStr += strconv.Itoa(int(v))
//			if i == 2{
//				formStr += ") "
//			}
//			continue
//		}
//
//		if i <= 5{
//			formStr += strconv.Itoa(int(v))
//			if i == 5{
//				formStr += "-"
//			}
//			continue
//		}else{
//			formStr += strconv.Itoa(int(v))
//		}
//
//
//	}
//
//	return formStr
//}