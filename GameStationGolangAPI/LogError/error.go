package LogError

import "log"
func LogError(s string,e error){
	if e!=nil {
		log.Println("❌ ",s,e)
	}
}