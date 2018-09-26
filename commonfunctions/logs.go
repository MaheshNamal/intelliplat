package commonfunc

import(
	"log"
	"os"
	"time"
	// "fmt"
)

func Loginit() {
    // fmt.Println("init in sandbox.go")
	f, err := os.OpenFile("logs/logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// f.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [AAA] ")
	log.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " >>> ")
	log.SetOutput(f)
	log.Println(err)
}

