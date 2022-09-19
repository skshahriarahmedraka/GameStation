package errors

import "log"
func ERROR(s string,e error){
	if e!=nil {
		log.Println("❌ ",s,e)
	}
}